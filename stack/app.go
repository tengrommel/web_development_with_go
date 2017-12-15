package main

import "fmt"

type Stack struct {
	slice []int
}

// Push adds the integer provided to the top of the stack
func (s *Stack)Push(i int)  {
	s.slice = append(s.slice, i)
}

// Peek returns the top item from the stack, but DOES NOT
// remove it from the stack.
func (s *Stack)Peek() int {
	return s.slice[len(s.slice)-1]
}

// Pop removes and returns the top item from the stack
func (s *Stack)Pop() int {
	var ret int = s.Peek()
	s.slice = s.slice[0:len(s.slice)-2]
	return ret
}

func (s *Stack)String() string {
	return fmt.Sprint(s.slice)
}

func main() {
	var s *Stack = new(Stack)
	s.Push(200)
	s.Push(123)
	s.Push(23)
	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
}
