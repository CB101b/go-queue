package queue

// This is a queue implemented with a dynamically adjusted array.
// The time complexity is quite high I think, due to copying of everything.
// Perhaps a linked list is better, but not for now.
type Queue[E any] struct {
	slice []E
}

func (q *Queue[E]) Push(elem E) {
	q.slice = append([]E{elem}, q.slice...)
}

func (q *Queue[E]) Take() (elem E) {
	elem = q.slice[len(q.slice)-1]
	newSlice := make([]E, len(q.slice)-1)
	copy(newSlice, q.slice[:len(q.slice)-1])
	q.slice = newSlice
	return
}
