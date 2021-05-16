package main

import "fmt"

func traverse(depGraph map[string][]string) {
	visited := map[string]struct{}{}
	for n := range depGraph {
		visit(n, depGraph, visited)
	}
	fmt.Printf("\n")
}

func visit(node string, depGraph map[string][]string, visited map[string]struct{}) {
	if _, ok := visited[node]; ok {
		return
	}
	nodes := depGraph[node]
	for _, n := range nodes {
		visit(n, depGraph, visited)
	}
	visited[node] = struct{}{}
	fmt.Printf("%s ", node)
}

func main() {
	depGraph := map[string][]string{
		"A": {"B", "C"},
		"B": {"C"},
		"C": {"D"},
		"D": {"E"},
		"E": {},
		"F": {"A"},
	}
	traverse(depGraph)
}
