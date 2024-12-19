package main

import (
	"fmt"
)

func WriteWithoutLock(state *State, update map[string]string) error {
	next_time := state.vector_time_manager.NextTime()
	fmt.Printf("New time: %v\n", next_time)
	for update_key, update_value := range update {
		if !state.kv_storage.Update(update_key, Value{data: update_value, tp: next_time}) {
			panic("Local update failed")
		}
		state.replication_service.SendAsync(Entry{key: update_key, value: Value{data: update_value, tp: next_time}})
	}
	return nil
}
