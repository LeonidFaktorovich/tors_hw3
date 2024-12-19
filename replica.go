package main

import (
	"context"
	"crdt/internal_service"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Addr struct {
	domain string
	port   uint16
}

type Replica struct {
	client         internal_service.InternalServiceClient
	queue          chan Entry
	settings_queue chan struct{}
}

func NewReplica(addr Addr) (*Replica, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", addr.domain, addr.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Replica{client: internal_service.NewInternalServiceClient(conn), queue: make(chan Entry, 1000000), settings_queue: make(chan struct{}, 100000)}, nil
}

func (r *Replica) SendAsync(e Entry) {
	r.queue <- e
}

func (r *Replica) RunDaemonAsync() {
	go func() {
		for {
			var entry Entry
			select {
			case entry = <-r.queue:
			case <-r.settings_queue:
				fmt.Printf("Block send entries\n")
				<-r.settings_queue
				fmt.Printf("Unblock send entries\n")
				continue
			}
			proto_tp := internal_service.VectorTime{HostTp: entry.value.tp.Tp}

			value := internal_service.Value{Data: entry.value.data, Tp: &proto_tp}
			proto_entry := internal_service.Entry{Key: entry.key, Value: &value}
			proto_enties := []*internal_service.Entry{&proto_entry}
			request := internal_service.EntriesRequest{Entries: proto_enties}
			for {
				ctx := context.Background()
				ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
				fmt.Printf("Send entries: %v\n", request.Entries)
				_, err := r.client.SendEntries(ctx, &request)
				cancel()
				if err != nil {
					fmt.Printf("Error with SendEntries: %v\n", err.Error())
					time.Sleep(5 * time.Second)
					continue
				}
				break
			}
		}
	}()
}
