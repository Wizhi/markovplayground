package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const (
		Bear = iota
		Bull
		Stagnant
	)

	labels := map[int]string{
		Bear:     "Bear",
		Bull:     "Bull",
		Stagnant: "Stagnant",
	}

	var transition [3][3]float32

	transition[Bear] = [3]float32{0.9, 0.075, 0.025}
	transition[Bull] = [3]float32{0.15, 0.8, 0.05}
	transition[Stagnant] = [3]float32{0.25, 0.25, 0.5}

	rng := rand.New(rand.NewSource(time.Now().Unix()))

	state := Bear

	for {
		fmt.Println(labels[state])

		p := transition[state]

		for i := range p {
			if p[i] < rng.Float32() {
				state = i
				break
			}
		}
	}
}
