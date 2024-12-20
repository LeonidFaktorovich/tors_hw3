package main

import (
	"context"
	"crdt/internal_service"
	"fmt"
	"time"
)

func EnableFullReplication(state *State) {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			state.global_mutex.RLock()
			if state.block_send {
				state.global_mutex.RUnlock()
				continue
			}

			proto_enties := []*internal_service.Entry{}
			for key, value := range state.kv_storage.storage {
				proto_tp := internal_service.VectorTime{HostTp: value.tp.Tp}
				value := internal_service.Value{Data: value.data, Tp: &proto_tp}
				proto_entry := internal_service.Entry{Key: key, Value: &value}
				proto_enties = append(proto_enties, &proto_entry)
			}

			request := internal_service.EntriesRequest{Entries: proto_enties}

			state.global_mutex.RUnlock()

			for _, replica := range state.replication_service.replicas {
				ctx := context.Background()
				ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
				fmt.Printf("Send entries: %v\n", request.Entries)
				_, err := replica.client.SendEntries(ctx, &request)
				cancel()
				if err != nil {
					fmt.Printf("Error with SendEntries in full replication: %v\n", err.Error())
					continue
				}
			}
		}
	}()
}
