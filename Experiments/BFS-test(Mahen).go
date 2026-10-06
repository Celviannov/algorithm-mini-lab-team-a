package main

import (
	"reflect"
	"testing"
)

func TestBFSVisitsValuesInBreadthFirstOrder(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 4, 5},
		3: {1, 5},
		4: {2},
		5: {2, 3},
	}

	want := []int{1, 2, 3, 4, 5}

	if got := bfs(graph, 1); !reflect.DeepEqual(got, want) {
		t.Fatalf("bfs(graph, 1) = %v; want %v", got, want)
	}
}

func TestBFSHandlesCyclesAndUnreachableValues(t *testing.T) {
	graph := map[int][]int{
		1: {2, 3},
		2: {1, 3},
		3: {1, 2},
		4: {4},
	}

	want := []int{1, 2, 3}

	if got := bfs(graph, 1); !reflect.DeepEqual(got, want) {
		t.Fatalf("bfs(graph, 1) = %v; want %v", got, want)
	}
}

func TestBFSIncludesStartWhenItIsNotInGraph(t *testing.T) {
	graph := map[int][]int{}

	want := []int{10}

	if got := bfs(graph, 10); !reflect.DeepEqual(got, want) {
		t.Fatalf("bfs(empty graph, 10) = %v; want %v", got, want)
	}
}

func TestBFSAnotherGraph(t *testing.T) {
	graph := map[int][]int{
		10: {20, 30},
		20: {10, 40},
		30: {10, 50},
		40: {20},
		50: {30},
	}

	want := []int{10, 20, 30, 40, 50}

	if got := bfs(graph, 10); !reflect.DeepEqual(got, want) {
		t.Fatalf("bfs(graph, 10) = %v; want %v", got, want)
	}
}
