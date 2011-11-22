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

func (a *triggers) Len() int {
	return len(a.v)
}
