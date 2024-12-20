package main

import (
	"flag"
	"fmt"
)

var node_id_flag = flag.Uint("node_id", 0, "node id")

var http_port_flag = flag.Int("http_port", 0, "http server port")

func MainDaemon() {
	flag.Parse()

	addr0 := Addr{domain: "localhost", port: 10010}
	addr1 := Addr{domain: "localhost", port: 10011}
	addr2 := Addr{domain: "localhost", port: 10012}

	addrs := []Addr{addr0, addr1, addr2}

	fmt.Printf("Node id: %v\n", *node_id_flag)
	grpc_port := addrs[*node_id_flag].port

	state := NewState(addrs, *node_id_flag)

	go func() {
		HttpDaemon(state, *http_port_flag)
	}()

	GrpcServiceDaemon(int(grpc_port), state)

	EnableFullReplication(state)
}
