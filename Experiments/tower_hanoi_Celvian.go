package main

import "fmt"

var moveCount int

func hanoi(n int, source, target, helper string) {
	if n == 0 {
		return
	}
	hanoi(n-1, source, helper, target)
	moveCount++
	fmt.Printf("Move disk %d: %s -> %s\n", n, source, target)
	hanoi(n-1, helper, target, source)
}

func main() {
	var disks int
	fmt.Print("Enter number of disks: ")
	fmt.Scanln(&disks)

	if disks < 1 {
		fmt.Println("Invalid input: number of disks must be at least 1")
		return
	}

	hanoi(disks, "A", "C", "B")
	fmt.Printf("Done! Total moves: %d (expected: %d)\n", moveCount, (1<<disks)-1)
}
