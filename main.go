package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/reactivex/rxgo/v2"
	y3 "github.com/yomorun/y3-codec-golang"

	"github.com/yomorun/yomo/pkg/quic"
	"github.com/yomorun/yomo/pkg/rx"
)

const batchSize = 1000

var bufferTime = rxgo.WithDuration(3 * time.Second)

var (
	client         *f.FaunaClient
	secret         = os.Getenv("FAUNA_SECRET")
	sinkServerAddr = "0.0.0.0:4141"
)

func init() {
	if secret == "" {
		panic("please set the secret in env FAUNA_SECRET")
	}
	// create a new FaunaClient
	client = f.NewFaunaClient(secret)
}

func main() {
	log.Print("Starting sink server...")
	quicServer := quic.NewServer(&quicServerHandler{
		readers: make(chan io.Reader),
	})
	err := quicServer.ListenAndServe(context.Background(), sinkServerAddr)
	if err != nil {
		log.Printf("❌ Serve the sink server on %s failure with err: %v", sinkServerAddr, err)
	}
}

type quicServerHandler struct {
	readers chan io.Reader
}

func (s *quicServerHandler) Listen() error {
	rxstream := rx.FromReaderWithY3(s.readers)
	observer := rxstream.Subscribe(0x10).
		OnObserve(decode).
		BufferWithTimeOrCount(bufferTime, batchSize)

	rxstream.Connect(context.Background())

	go bulkInsert(observer)
	return nil
}

func (s *quicServerHandler) Read(qs quic.Stream) error {
	s.readers <- qs
	return nil
}

// decode the noise value via Y3
func decode(v []byte) (interface{}, error) {
	data, err := y3.ToFloat32(v)
	if err != nil {
		log.Printf("err: %s\n", err.Error())
	}
	return data, err
}

// bulk insert data into FaunaDB
func bulkInsert(observer rx.RxStream) error {
	for ch := range observer.Observe() {
		if ch.Error() {
			log.Println(ch.E.Error())
		} else if ch.V != nil {
			items, ok := ch.V.([]interface{})
			if !ok {
				log.Println(ok)
				continue
			}

			// bulk insert
			_, err := client.Query(
				f.Map(
					items,
					f.Lambda(
						"noise",
						f.Create(
							f.Collection("noise"),
							f.Obj{"data": f.Obj{"noise": f.Var("noise")}},
						),
					),
				),
			)
			if err != nil {
				return err
			}

			log.Printf("Insert %d noise data into InfluxDB...", len(items))
		}
	}

	return nil
}
