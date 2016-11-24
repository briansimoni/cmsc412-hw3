package main

import (
	"sort"
	"strconv"
	"os"
)


// the main function will create a graph object
// read the graph.txt file
func main() {

	g := createGraphFromFile("graph4.txt")
	m, k := g.maxKcore()

	// convert to ints rite quick
	kcore := make ([]int, 0)
	for i := range m {
		n, err := strconv.Atoi(m[i])
		check(err)
		kcore = append(kcore, n)
	}
	sort.Ints(kcore)
	maxK := strconv.Itoa(k)

	s := maxK
	s += "\n"

	for i := range kcore{
		id := strconv.Itoa(kcore[i])
		s += id + " "
	}

	f, err := os.Create("kcore.txt")
	check(err)
	f.WriteString(s)


}

func check(e error) {
	if e != nil {
		panic(e)
	}
}