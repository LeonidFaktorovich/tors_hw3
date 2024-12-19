package main

import (
	"sync"
	"time"
)

type State struct {
	kv_storage          *KVStorage
	vector_time_manager *VectorTimeManager

	replication_service *ReplicationService

	sleep         time.Duration
	block_receive bool
	block_send    bool

	global_mutex sync.RWMutex
}

func NewState(addrs []Addr, curr_host_index uint) *State {
	kv_storage := NewKVStorage()
	vector_time_manager := NewVectorTimeManager(uint(len(addrs)), curr_host_index)

	result_addrs := append(addrs[:curr_host_index], addrs[curr_host_index+1:]...)
	replication_service := NewReplicationService(result_addrs)
	return &State{kv_storage: kv_storage, vector_time_manager: vector_time_manager, replication_service: replication_service, sleep: 0, block_receive: false, block_send: false}
}
