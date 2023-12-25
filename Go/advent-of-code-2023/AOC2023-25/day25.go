package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name string
}

type Graph struct {
	nodes map[string]Node
	edges map[string][]string
}

func main() {
	var input = "AOC2023-25/input1.txt"
	var ex = "AOC2023-25/ex.txt"
	fmt.Println("Result ex:", P1(ex))
	fmt.Println("Result:", P1(input))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var graph = Graph{}
	var nodes = make(map[string]Node)
	var edges = make(map[string][]string)
	var adjacent = make(map[string][]string)
	var lastNode = ""
	for scanner.Scan() {
		line := scanner.Text()
		var elems = strings.Split(line, ": ")
		var key = elems[0]
		var values = strings.Split(elems[1], " ")
		var node = Node{key}
		for _, value := range values {
			if adjacent[key] == nil {

			}
			if !contains(adjacent[key], value) {
				adjacent[key] = append(adjacent[key], value)
			}
			if !contains(adjacent[value], key) {
				adjacent[value] = append(adjacent[value], key)
			}

			var keyEdge = encodeEdge(key, value)
			edges[keyEdge] = append(edges[keyEdge], key)
			edges[keyEdge] = append(edges[keyEdge], value)
			nodes[value] = Node{value}
		}
		lastNode = node.name
		nodes[key] = node
	}
	graph.nodes = nodes
	graph.edges = edges
	return solve(graph, adjacent, lastNode)
}

func solve(graph Graph, adjacent map[string][]string, sourceNode string) int {
	//We start with the source node
	var group1 = make([]Node, 0)
	var group2 = make([]Node, 0)
	var startNode = graph.nodes[sourceNode]
	group1 = append(group1, startNode)
	for _, node := range graph.nodes {
		if node.name != startNode.name {
			if isReachableAfterCut(startNode.name, node.name, adjacent, 3) {
				group1 = append(group1, node)
			} else {
				group2 = append(group2, node)
			}
		}
	}
	fmt.Println(len(group1), len(group2))
	return len(group1) * len(group2)
}

func isReachableAfterCut(start string, end string, adjacent map[string][]string, numCut int) bool {
	var isReachable = true
	var visitedEdges = make(map[string]bool)
	var currentAdjacent = adjacent
	for i := 0; i <= numCut; i++ {
		var ok, newAdjacent, newEdges = findAndRemoveBFS(start, end, currentAdjacent, visitedEdges)
		for k := range newEdges {
			visitedEdges[k] = true
		}
		if ok {
			currentAdjacent = newAdjacent
		} else {
			isReachable = false
			break
		}
	}
	return isReachable
}

func findAndRemoveBFS(start string, end string, adjacent map[string][]string, visitedEdges map[string]bool) (bool, map[string][]string, map[string]bool) {
	var queue = make([]string, 0)
	var visited = make(map[string]bool)
	var parent = make(map[string]string)
	var found = false
	queue = append(queue, start)
	visited[start] = true
	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]
		if current == end {
			var ok = false
			for !ok {
				if current == "" {
					ok = true
					break
				}
				var currentParent = parent[current]
				if currentParent != "" {
					var keyEdge = encodeEdge(current, currentParent)
					if !visitedEdges[keyEdge] {
						visitedEdges[keyEdge] = true
					}
				}

				current = currentParent
			}
			found = true
			break
		}
		for _, neighbor := range adjacent[current] {
			if !visited[neighbor] && !visitedEdges[encodeEdge(current, neighbor)] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				parent[neighbor] = current
			}
		}
	}
	return found, adjacent, visitedEdges
}

func encodeEdge(a string, b string) string {
	if a < b {
		return a + b
	}
	return b + a
}

func contains(list []string, word string) bool {
	for _, value := range list {
		if value == word {
			return true
		}
	}
	return false
}
