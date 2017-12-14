package main

import "fmt"

type Queue struct {
	slice []int
}

// Enqueue adds the integer provided to the back
func (q *Queue)Enqueue(i int)  {
	q.slice = append(q.slice, i)
}

// Dequeue returns the first item in the Queue
// and removes that item from the Queue
// or panic if there isn't one
func (q *Queue)Dequeue() int {
	ret := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return ret
}

func (q *Queue)String() string {
	return fmt.Sprint(q.slice)
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(123)
	q.Enqueue(13)
	q.Enqueue(3)
	q.Enqueue(0)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
