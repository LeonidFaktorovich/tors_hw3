package main

type VectorTime struct {
	Tp []uint64 `json:"tp"`
}

func Less(lhs VectorTime, rhs VectorTime) bool {
	for i, tp := range rhs.Tp {
		if lhs.Tp[i] == tp {
			continue
		}
		return lhs.Tp[i] < tp
	}
	return false
}
