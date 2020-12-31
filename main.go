package main

import (
	"context"
	"fmt"
	"log"
	"os"

	f "github.com/fauna/faunadb-go/v3/faunadb"

	"github.com/yomorun/yomo/pkg/quic"
	"github.com/yomorun/yomo/pkg/rx"
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
	rxStream := rx.FromReader(st).
		Y3Decoder("0x10", float32(0)).
		StdOut()

	go func() {
		for customer := range rxStream.Observe() {
			if customer.Error() {
				fmt.Println(customer.E.Error())
			} else if customer.V != nil {
				err := store(customer.V)
				if err != nil {
					log.Printf("save data `%v` error: %s", customer.V, err.Error())
				} else {
					log.Printf("save `%v` to FaunaDB\n", customer.V)
				}
			}
		}
	}()

	return nil
}

type Noise struct {
	Value float32 `fauna:"value"`
}

// store save data to the FaunaDB
func store(i interface{}) error {
	value := i.(float32)

	noise := Noise{Value: value}
	_, err := client.Query(f.Create(f.Collection("noise"), f.Obj{"data": noise}))
	if err != nil {
		return err
	}

	return nil
}
