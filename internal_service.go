package main

import (
	"context"
	"crdt/internal_service"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InternalService struct {
	internal_service.UnimplementedInternalServiceServer
	state *State
}

func (i *InternalService) SendEntries(ctx context.Context, request *internal_service.EntriesRequest) (*internal_service.EntriesResponse, error) {
	i.state.global_mutex.RLock()
	sleep_duration := i.state.sleep
	block_receive := i.state.block_receive
	i.state.global_mutex.RUnlock()

	if block_receive {
		return nil, status.Error(codes.Unavailable, "fault injection")
	}
	if sleep_duration > time.Second*0 {
		time.Sleep(sleep_duration)
	}

	i.state.global_mutex.Lock()
	defer i.state.global_mutex.Unlock()

	fmt.Printf("Receive entries %v\n", request.Entries)

	for _, entry := range request.Entries {
		tp := VectorTime{Tp: entry.Value.Tp.HostTp}

		i.state.kv_storage.Update(entry.Key, Value{data: entry.Value.Data, tp: tp})
		i.state.vector_time_manager.UpdateTime(tp)
	}

	return &internal_service.EntriesResponse{}, nil
}
