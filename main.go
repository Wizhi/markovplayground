package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	options := [...]string{"Bear", "Bull", "Stagnant"}
	transition := [...][3]float32{
		{0.9, 0.075, 0.025},
		{0.15, 0.8, 0.05},
		{0.25, 0.25, 0.5},
	}

	rng := rand.New(rand.NewSource(time.Now().Unix()))

	state := 0

	for {
		fmt.Println(options[state])

		p := transition[state]

		for i := range p {
			if p[i] < rng.Float32() {
				state = i
				break
			}
		}
	}
}
