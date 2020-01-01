package main

import (
	"fmt"

	"github.com/fsetiawan29/design-pattern/structural/adapter"
)

func main() {
	roundHole := adapter.NewRoundHole(5)
	roundPeg := adapter.NewRoundPeg(5)
	fmt.Printf("%+v\n", roundHole.Fits(roundPeg))

	smallSquarePeg := adapter.NewSquarePeg(5)
	largeSquarePeg := adapter.NewSquarePeg(10)
	// roundHole.Fits(smallSquarePeg) // doesn't work

	smallSquarePegAdapter := adapter.NewSquarePegAdapter(smallSquarePeg)
	fmt.Printf("%+v\n", roundHole.Fits(smallSquarePegAdapter))
	largeSquarePegAdapter := adapter.NewSquarePegAdapter(largeSquarePeg)
	fmt.Printf("%+v\n", roundHole.Fits(largeSquarePegAdapter))
}
