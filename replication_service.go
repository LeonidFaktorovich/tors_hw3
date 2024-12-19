package main

type ReplicationService struct {
	replicas []*Replica
}

func NewReplicationService(addrs []Addr) *ReplicationService {
	replicas := make([]*Replica, len(addrs))
	for i := 0; i < len(addrs); i++ {
		replica, err := NewReplica(addrs[i])
		if err != nil {
			panic(err)
		}
		replicas[i] = replica
		replicas[i].RunDaemonAsync()
	}
	return &ReplicationService{replicas: replicas}
}

func (r *ReplicationService) SendAsync(e Entry) {
	for _, replica := range r.replicas {
		replica.SendAsync(e)
	}
}

func (r *ReplicationService) ChangeMode() {
	for _, replica := range r.replicas {
		replica.settings_queue <- struct{}{}
	}
}
