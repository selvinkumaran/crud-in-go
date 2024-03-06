package main

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Topic 1: Concurrency in Go
func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

// Topic 2: Interfaces and Type Assertion
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func getArea(s Shape) float64 {
	return s.Area()
}

// Topic 3: Error Handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Topic 6: Concurrency Patterns
func squareWorker(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()
	for num := range in {
		out <- num * num
	}
}

// Topic 7: Package Management with Go Modules (No code needed here)

// Topic 8: Testing in Go
func TestArea(t *testing.T) {
	rect := Rectangle{Width: 3, Height: 4}
	expectedArea := 12.0
	actualArea := rect.Area()
	if actualArea != expectedArea {
		t.Errorf("Expected area: %f, but got: %f", expectedArea, actualArea)
	}
}

func main() {
	// Topic 1: Concurrency in Go
	fmt.Println("=== Concurrency in Go ===")
	go printNumbers()
	time.Sleep(time.Second * 3)

	// Topic 2: Interfaces and Type Assertion
	fmt.Println("\n=== Interfaces and Type Assertion ===")
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Area of rectangle:", getArea(rect))

	// Topic 3: Error Handling
	fmt.Println("\n=== Error Handling ===")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Topic 6: Concurrency Patterns
	fmt.Println("\n=== Concurrency Patterns ===")
	nums := []int{1, 2, 3, 4, 5}
	in := make(chan int)
	out := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go squareWorker(&wg, in, out)
	}
	go func() {
		for _, num := range nums {
			in <- num
		}
		close(in)
	}()
	go func() {
		wg.Wait()
		close(out)
	}()
	for squaredNum := range out {
		fmt.Println("Squared number:", squaredNum)
	}

	// Topic 8: Testing in Go (No code needed here)
}
