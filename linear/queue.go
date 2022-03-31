package linear

// package linear

// CirQueue
type CirQueue struct {
	Data [10]int // queue
	Head int     // point to first element
	Tail int     // point to end element + 1
}

// IsFull have one element free spare
func (q *CirQueue) IsFull() bool {
	size := len(q.Data)
	return (q.Tail+1)%size == q.Head
}

func (q *CirQueue) IsNull() bool {
	return q.Tail == q.Head
}

func (q *CirQueue) Add(e int) {
	if q.IsFull() {
		return
	}

	size := len(q.Data)
	q.Data[q.Tail] = e
	q.Tail = (q.Tail + 1) % size
}

func (q *CirQueue) Get() {
	if q.IsNull() {
		return
	}

	size := len(q.Data)
	q.Data[q.Head] = 0
	q.Head = (q.Head + 1) % size
}

func (q *CirQueue) Length() int {
	size := len(q.Data)
	return (q.Tail + size - q.Head) % size
}

// LinkQueue
type LinkQueue struct {
	Data int
	Next *LinkQueue
}
