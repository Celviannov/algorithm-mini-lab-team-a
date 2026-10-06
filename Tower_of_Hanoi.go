package main

import "fmt"

var moveCount int

func hanoi(n int, source, target, helper string) {
	if n == 0 {
		return
	}
	hanoi(n-1, source, helper, target) // move n-1 disks to helper
	moveCount++
	fmt.Printf("Move disk %d: %s -> %s\n", n, source, target)
	hanoi(n-1, helper, target, source) // move n-1 disks to target
}

func main() {
	disks := 3
	hanoi(disks, "A", "C", "B")
	fmt.Printf("Done! Total moves: %d (expected: %d)\n", moveCount, (1<<disks)-1)
}