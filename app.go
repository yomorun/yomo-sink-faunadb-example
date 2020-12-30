package main

import (
	"context"
	"fmt"
	"os"
	"time"

	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/yomorun/yomo/pkg/rx"
)

var client *f.FaunaClient

// define a struct to encode data to db
type Noise struct {
	Value float32 `fauna:"value"`
}

func init() {
	// create a new FaunaClient
	client = f.NewFaunaClient(os.Getenv("FAUNA_SECRET"))
}

// store: save value to FaunaDB
var store = func(_ context.Context, i interface{}) (interface{}, error) {
	value := i.(float32)

	noise := Noise{Value: value}
	_, err := client.Query(f.Create(f.Collection("noise"), f.Obj{"data": noise}))
	if err != nil {
		panic(err)
	}

	fmt.Printf("save `%v` to FaunaDB\n", value)
	return value, nil
}

// Handler will handle data in Rx way
func Handler(rxstream rx.RxStream) rx.RxStream {
	stream := rxstream.
		Y3Decoder("0x10", float32(0)).
		AuditTime(100 * time.Millisecond).
		Map(store)
	return stream
}
