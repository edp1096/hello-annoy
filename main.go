package main // import "hello-annoy"

import (
	"annoyindex"
	"fmt"
	"math/rand"
)

func angular() {
	dimension := 40
	indexing := annoyindex.NewAnnoyIndexAngular(dimension)
	for i := 0; i < 1000; i++ {
		item := make([]float32, 0, dimension)
		for x := 0; x < dimension; x++ {
			item = append(item, rand.Float32())
		}
		indexing.AddItem(i, item)
	}
	indexing.Build(10)
	indexing.Save("test_ang.ann")

	annoyindex.DeleteAnnoyIndexAngular(indexing)

	indexing = annoyindex.NewAnnoyIndexAngular(dimension)
	indexing.Load("test_ang.ann")

	var result []int
	indexing.GetNnsByItem(0, 1000, -1, &result)
	fmt.Printf("%v\n", result)
}

func euclidean() {
	dimension := 3
	indexing := annoyindex.NewAnnoyIndexEuclidean(dimension)

	indexing.AddItem(0, []float32{1.2, 3.4, 5.6})
	indexing.AddItem(1, []float32{2.1, 4.3, 6.5})
	indexing.AddItem(2, []float32{0.8, 1.5, 2.7})
	indexing.AddItem(3, []float32{3.2, 2.8, 1.4})

	indexing.Build(10)
	indexing.Save("test_euc.ann")

	indexing.Unload()

	var nearest []int

	indexing = annoyindex.NewAnnoyIndexEuclidean(dimension)
	indexing.Load("test_euc.ann")

	query := []float32{0.9, 2.3, 4.5}
	nNeighbors := 2

	indexing.GetNnsByVector(query, nNeighbors, -1, &nearest)

	fmt.Println("Query Vector:", query)
	fmt.Println("Nearest Neighbors:", nearest)

	var valueFirst, valueSecond []float32
	indexing.GetItem(nearest[0], &valueFirst)
	indexing.GetItem(nearest[1], &valueSecond)
	fmt.Println(nearest[0], valueFirst)
	fmt.Println(nearest[1], valueSecond)
	dist := indexing.GetDistance(nearest[0], nearest[1])
	fmt.Println("Distance between nearest 1st and 2nd:", dist)
}

func main() {
	angular()
	euclidean()
}
