package main

import "fmt"

type VectorTimeManager struct {
	curr_time       VectorTime
	curr_host_index uint
}

func NewVectorTimeManager(hosts_count uint, curr_host_index uint) *VectorTimeManager {
	curr_time := VectorTime{Tp: make([]uint64, hosts_count)}
	return &VectorTimeManager{curr_time: curr_time, curr_host_index: curr_host_index}
}

func (v *VectorTimeManager) NextTime() VectorTime {
	new_time := make([]uint64, len(v.curr_time.Tp))
	copy(new_time, v.curr_time.Tp)
	new_time[v.curr_host_index] += 1
	v.curr_time.Tp[v.curr_host_index] += 1
	return VectorTime{Tp: new_time}
}

func (v *VectorTimeManager) UpdateTime(msg_time VectorTime) {
	for i, tp := range msg_time.Tp {
		v.curr_time.Tp[i] = max(v.curr_time.Tp[i], tp)
	}
	fmt.Printf("New time: %v\n", v.curr_time.Tp)
}
