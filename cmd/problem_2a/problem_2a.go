package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Graph struct will represent all edges.
// Key in the map       -> vertex value.
// Val for each vertex  -> list of edges.
type Graph struct {
	edges map[int][]*Edge
}

// to       -> destination vertex, where this edge goes to.
// capacity -> weight of the edge, capacity that can flow through it.
// flow     -> what is currently flowing through the edge.
type Edge struct {
	to       int
	capacity int
	flow     int
}

// NewGraph() just creates a new graph instance.
// The map will hold vertex values as the keys and their outgoing edges in a slice as the val.
// NOTE: the edges slice for each vertex will also contain any backwards edges.
func NewGraph(vertexCount int) *Graph {
	graph := &Graph{
		edges: make(map[int][]*Edge),
	}

	for i := range vertexCount {
		graph.edges[i] = []*Edge{}
	}

	return graph
}

func (graph *Graph) AddEdge(v1, v2, weight int) {
	// Forward edge (v1 -> v2).
	forward := &Edge{
		to:       v2,
		capacity: weight,
		flow:     0,
	}

	// Backward edge (v2 -> v1).
	backward := &Edge{
		to:       v1,
		capacity: 0,
		flow:     0,
	}

	// Append to respective vertex Edge slice.
	graph.edges[v1] = append(graph.edges[v1], forward)
	graph.edges[v2] = append(graph.edges[v2], backward)
}

func bronKerbosch() {}

func readFile(fileName string) (*Graph, int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening the file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("readFile(): Issue scanning the file -- may be empty..")
	}

	// First row (2 values) -> vertex count and then the arc count.
	firstLine := strings.Fields(scanner.Text())
	if len(firstLine) < 2 || len(firstLine) > 2 {
		log.Fatal(
			"readFile(): Error parsing the first line -- either too few arguments or too many. Should only include vertex count and arc count",
		)
	}

	vertexCount, err := strconv.Atoi(firstLine[0])
	if err != nil {
		log.Fatal("readFile(): Error reading in the vertexCount")
	}

	arcCount, err := strconv.Atoi(firstLine[1])
	if err != nil {
		log.Fatal("readFile(): Error reading in the arcCount")
	}

	graph := NewGraph(vertexCount)
	source := 0
	sink := 0
	arcReadCount := 1

	// Read in the remaining rows containing the edges (v1, v2, arc weight).
	for scanner.Scan() {

		// If there's any extra rows that are not included in the arcCount defined, break.
		if arcReadCount > arcCount {
			break
		}

		line := strings.Fields(scanner.Text())
		if len(line) < 3 || len(line) > 3 {
			log.Fatal("readFile(): vertex and weight row either has too few or too many arguments")
		}

		v1, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal("readFile(): Error reading in v1")
		}

		v2, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal("readFile(): Error reading in the v2")
		}

		weight, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatal("readFile(): Error reading in the weight")
		}

		// Track the sink (max v2 found).
		if v2 > sink {
			sink = v2
		}

		// Create and add edge to the graph (both forward and backward)
		graph.AddEdge(v1, v2, weight)
		arcReadCount++
	}

	return graph, source, sink
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: ./bin/problem2a filename")
	}

  // build the graph
  graph, source, sink := readFile(os.Args[1])

  log.Println(graph)
  log.Println(source)
  log.Println(sink)
}
