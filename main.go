package main

import (
	"fmt"
	"time"
)

func mainPi() {
	fmt.Println(pi(5000))
}

func mainTest() {
	test()
}

// The prime sieve: Daisy-chain Filter processes.
func mainSieve() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

func mainSolitaire() {
	if !solve() {
		fmt.Println("no solution found")
	}
	fmt.Println(moves, "moves tried")
}

func mainTree() {
	t1 := New(100, 1)
	fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
	fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
	fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
	fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
}

func mainLife() {
	l := NewLife(40, 15)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
}

func mainFib() {
	f := fib()
	// Function calls are evaluated left-to-right.
	for i := 0; i < 22; i++ {
		fmt.Println(f(), " ")
	}
}

// -------------------------------------
// Print i! for i in [0,9]

func mainPeano() {
	for i := 0; i <= 9; i++ {
		f := count(fact(gen(i)))
		fmt.Println(i, "! =", f)
	}
}

func main() {
	fmt.Println("pi")
	mainPi()
	// fmt.Println("test")
	// main_test()
	fmt.Println("sieve")
	mainSieve()
	fmt.Println("solitaire")
	mainSolitaire()
	fmt.Println("tree")
	mainTree()
	// fmt.Println("life")
	// main_life()
	fmt.Println("fib")
	mainFib()
	fmt.Println("peano")
	mainPeano()
}
