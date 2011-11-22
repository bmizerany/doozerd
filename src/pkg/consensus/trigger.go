package consensus

type trigger struct {
	t int64 // trigger time
	n int64 // seqn
}

func (t trigger) Less(y interface{}) bool {
	return t.t < y.(trigger).t
}

type triggers struct {
	v []trigger
}

func (a triggers) Len() int {
	return len(a.v)
}

func (a triggers) Less(i, j int) bool {
	return a.v[i].t < a.v[j].t
}

func (a triggers) Pop() (x interface{}) {
	n := len(a.v)
	x, a.v = a.v[n-1], a.v[:n-1]
	return
}

func (a triggers) Push(x interface{}) {
	a.v = append(a.v, x.(trigger))
}

func (a triggers) Swap(i, j int) {
	a.v[i], a.v[j] = a.v[j], a.v[i]
}
