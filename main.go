package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	options := [...]string{"Bear", "Bull", "Stagnant"}
	transition := [...][3]float64{
		{0.9, 0.075, 0.025},
		{0.15, 0.8, 0.05},
		{0.25, 0.25, 0.5},
	}

	spectrum := make([][]float64, len(transition))

	for i, p := range transition {
		a := make([]float64, len(p))

		for j, acc := 0, float64(0); j < len(a); j++ {
			acc += p[j]
			a[j] = acc
		}

		spectrum[i] = a
	}

	rng := rand.New(rand.NewSource(time.Now().Unix()))

	state := 0

	for i := 0; i < 10; i++ {
		fmt.Println(options[state])

		r := rng.Float64()
		s := spectrum[state]

		for j := range s {
			if r < s[j] {
				state = j
				break
			}
		}
	}
}
