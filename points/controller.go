package points

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

func FindDistance(x, y float64) ([]Points, error) {
	var (
		err                error
		pointsFile, points []Points
	)

	file, _ := ioutil.ReadFile("data/points.json")

	err = json.Unmarshal([]byte(file), &pointsFile)
	if err != nil {
		return points, err
	}

	var x1 = []float64{x}
	var y1 = []float64{y}

	distanceXY := manhattanDistance(x1, y1)

	points = []Points{}

	for _, coord := range pointsFile {
		distance := manhattanDistance([]float64{coord.X}, []float64{coord.Y})

		if distance <= distanceXY {
			coord.Distance = distance
			points = append(points, coord)
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].Distance < points[j].Distance
	})

	return points, err
}

func manhattanDistance(x []float64, y []float64) float64 {

	var s float64

	if len(y) != len(x) {
		fmt.Println("The x and y coordinate vector must be the same size!")
		return 0
	}

	s = 0

	for i := 0; i < len(x); i += 1 {

		s += math.Abs(y[i] - x[i])
	}

	return float64(s)
}
