package main

import (
	"math"
)

type circle struct {
	centre vector
	radius float64
}

func getDist(p1, p2 vector) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) +
		math.Pow(p2.y-p1.y, 2))
}

func collides(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c2.centre.x-c1.centre.x, 2) +
		math.Pow(c2.centre.y-c1.centre.y, 2))
	return dist <= c1.radius+c2.radius
}

func checkCollisions() error {
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if collides(c1, c2) && elements[i].active && elements[j].active {
						err := elements[i].collision(elements[j])
						if err != nil {
							return err
						}
						err = elements[j].collision(elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
