package queue

// This is a queue implemented with a dynamically adjusted array.
// The time complexity is quite high I think, due to copying of everything.
// Perhaps a linked list is better, but not for now.
type Queue[E any] struct {
	slice []E
}

func (q *Queue[E]) Push(elem E) {
	q.slice = append(q.slice, elem)
}

func (q *Queue[E]) Take() (elem E) {
	elem = q.slice[0]
	q.slice = q.slice[1:]
	return
}
