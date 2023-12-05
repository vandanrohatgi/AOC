package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	var solution int
	var isTest bool
	var f []byte
	var err error

	flag.IntVar(&solution, "s", 1, "Solution to execute (1 or 2)")
	flag.BoolVar(&isTest, "t", false, "run solution in testing mode")
	flag.Parse()

	if isTest {
		f, err = os.ReadFile("test.txt")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err = os.ReadFile("input.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	f = bytes.TrimSpace(f)
	if solution == 1 {
		sol1(f)
	} else {
		sol2(f)
	}
}

type Graph struct {
	Vertices map[string]Vertex
}

type Vertex struct {
	Edges map[string]int
}

func (g *Graph) InitializeGraph() {
	g.Vertices = make(map[string]Vertex)
}

func (g *Graph) AddVertex(src string) {
	if _, ok := g.Vertices[src]; !ok {
		g.Vertices[src] = Vertex{Edges: make(map[string]int)}
	}
}

func (g *Graph) AddEdge(src, dst string, weight int) {
	g.Vertices[src].Edges[dst] = weight
}

// findClosest find the closest node which hasn't been traversed
func (v Vertex) findClosest(travelled map[string]int) string {
	minDist := math.MaxInt32
	var minKey string
	for k, v := range v.Edges {
		if _, found := travelled[k]; !found && v < minDist {
			minDist = v
			minKey = k
		}
	}
	return minKey
}

func sol1(f []byte) {
	r := bytes.NewReader(f)
	buf := bufio.NewScanner(r)

	var src, dst string
	var weight int
	var inputGraph Graph
	inputGraph.InitializeGraph()

	for buf.Scan() {
		fmt.Fscanf(strings.NewReader(buf.Text()), "%s to %s = %d", &src, &dst, &weight)

		inputGraph.AddVertex(src)
		inputGraph.AddVertex(dst)

		inputGraph.AddEdge(src, dst, weight)
		inputGraph.AddEdge(dst, src, weight)
	}

	fmt.Println(inputGraph.Vertices)
	var travelled = make(map[string]int)

	var startingPoint = "AlphaCentauri"
	travelled[startingPoint] = 0

	for len(travelled) < len(inputGraph.Vertices) {
		time.Sleep(1 * time.Second)
		fmt.Println(travelled)
		closest := inputGraph.Vertices[startingPoint].findClosest(travelled)
		travelled[closest] = inputGraph.Vertices[startingPoint].Edges[closest]
		startingPoint = closest
	}

	fmt.Println(travelled)
	var plswork int
	for _, v := range travelled {
		plswork += v
	}

	fmt.Println(plswork)
}

func sol2(f []byte) {

}
