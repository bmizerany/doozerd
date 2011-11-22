package consensus

import (
	"net"
)

type packet struct {
	Addr *net.UDPAddr
	msg
}

type packets struct {
	v []packet
}

func (a packets) Len() int {
	return len(a.v)
}

func (a packets) Less(i, j int) bool {
	return *a.v[i].Seqn < *a.v[j].Seqn
}

func (a packets) Pop() (x interface{}) {
	n := len(a.v)
	x, a.v = a.v[n-1], a.v[:n-1]
	return
}

func (a packets) Push(x interface{}) {
	a.v = append(a.v, x.(packet))
}

func (a packets) Swap(i, j int) {
	a.v[i], a.v[j] = a.v[j], a.v[i]
}
