package main

import (
	"context"
	"log"
	"os"

	f "github.com/fauna/faunadb-go/v3/faunadb"
	y3 "github.com/yomorun/y3-codec-golang"

	"github.com/yomorun/yomo/pkg/quic"
)

var (
	client         *f.FaunaClient
	secret         = os.Getenv("FAUNA_SECRET")
	sinkServerAddr = "0.0.0.0:4141"
)

func init() {
	// create a new FaunaClient
	client = f.NewFaunaClient(secret)
}

func main() {
	go serveSinkServer(sinkServerAddr)
	select {}
}

// serveSinkServer serves the Sink server over QUIC.
func serveSinkServer(addr string) {
	log.Print("Starting sink server...")
	quicServer := quic.NewServer(&quicServerHandler{})
	err := quicServer.ListenAndServe(context.Background(), addr)
	if err != nil {
		log.Printf("❌ Serve the sink server on %s failure with err: %v", addr, err)
	}
}

type quicServerHandler struct {
}

func (s *quicServerHandler) Listen() error {
	return nil
}

func (s *quicServerHandler) Read(st quic.Stream) error {
	// decode the data via Y3 Codec.
	ch := y3.
		FromStream(st).
		Subscribe(0x10).
		OnObserve(onObserve)

	go func() {
		for {
			item, ok := <-ch
			if ok {
				// store data to FaunaDB
				err := store(item)
				if err != nil {
					log.Printf("save data `%v` error: %s", item, err.Error())
				} else {
					log.Printf("save `%v` to FaunaDB\n", item)
				}
			}
		}
	}()

	return nil
}

type noiseData struct {
	Noise float32 `yomo:"0x11" fauna:"noise"` // Noise value
	Time  int64   `yomo:"0x12" fauna:"time"`  // Timestamp (ms)
	From  string  `yomo:"0x13" fauna:"from"`  // Source IP
}

func onObserve(v []byte) (interface{}, error) {
	// decode the data via Y3 Codec.
	data := noiseData{}
	err := y3.ToObject(v, &data)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return data, nil
}

// store save data to the FaunaDB
func store(v interface{}) error {
	_, err := client.Query(f.Create(f.Collection("noise"), f.Obj{"data": v}))
	if err != nil {
		return err
	}

	return nil
}
