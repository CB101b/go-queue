package queue

type Queue[E any] struct {
	slice []E
}

func (q *Queue[E]) Push(elem E) {
	q.slice = append([]E{elem}, q.slice...)
}
