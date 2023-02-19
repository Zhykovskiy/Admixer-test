package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Point struct {
	X, Y int
}

type Rectangle struct {
	A, B, C, D Point
}

func dotProduct(a, b, c Point) int {
	return (b.X-a.X)*(c.X-a.X) + (b.Y-a.Y)*(c.Y-a.Y)
}

func sidesAreOrthogonal(a, b, c, d Point) bool {
	return dotProduct(a, b, c) == 0 &&
		dotProduct(b, a, d) == 0 &&
		dotProduct(d, b, c) == 0 &&
		dotProduct(c, a, d) == 0
}

func findRectangles(points []Point) (rectangles []Rectangle) {
	length := len(points)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				for l := k + 1; l < length; l++ {
					if sidesAreOrthogonal(points[i], points[j], points[k], points[l]) {
						rectangles = append(rectangles, Rectangle{points[i], points[j], points[l], points[k]})
					}
				}
			}
		}
	}

	return rectangles
}

func sortPoints(points []Point) {
	sort.Slice(points, func(i, j int) bool {
		if points[i].X != points[j].X {
			return points[i].X < points[j].Y
		}

		return points[i].Y < points[j].Y
	})
}

func countRectangles(points []Point) (answer int) {
	sortPoints(points)

	rectangles := findRectangles(points)

	return len(rectangles)
}

func contains(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}

	return false
}

func main() {
	n := rand.Intn(100)
	points := []Point{}

	for i := 0; i < n; i++ {
		x := rand.Intn(20)
		y := rand.Intn(20)

		if !contains(points, Point{x, y}) {
			points = append(points, Point{x, y})
		}
	}

	fmt.Println("Answer: ", countRectangles(points))
}
