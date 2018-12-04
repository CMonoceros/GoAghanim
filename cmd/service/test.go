package main

import (
	"cmonoceros.com/GoAghanim/gen-go/test"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"strings"
)

type FormatDataImpl struct{}

func (fdi *FormatDataImpl) DoFormat(data *test.Data) (r *test.Data, err error) {
	var rData test.Data
	rData.Text = strings.ToUpper(data.Text)

	return &rData, nil
}

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	handler := &FormatDataImpl{}
	processor := test.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST+":"+PORT)
	if err := server.Serve(); err != nil {
		fmt.Println("Error Running at:", err)
	}
}
