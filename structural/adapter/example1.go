package adapter

/*
This example of the Adapter pattern is based on
classic conflict between square pegs and round holes
Pseudocode based on https://refactoring.guru/design-patterns/adapter
*/

import "math"

// RoundHole is
type RoundHole struct {
	radius float32
}

// NewRoundHole is
func NewRoundHole(radius float32) *RoundHole {
	return &RoundHole{
		radius,
	}
}

func (rh *RoundHole) getRadius() float32 {
	return rh.radius
}

// Fits is
func (rh RoundHole) Fits(roundPeg RoundPegItf) bool {
	return rh.getRadius() >= roundPeg.getRadius()
}

// ==================

// RoundPeg is
type RoundPeg struct {
	radius float32
}

// NewRoundPeg is
func NewRoundPeg(radius float32) *RoundPeg {
	return &RoundPeg{
		radius,
	}
}

func (rp *RoundPeg) getRadius() float32 {
	return rp.radius
}

// ==================

// SquarePeg is
type SquarePeg struct {
	width float32
}

// NewSquarePeg is
func NewSquarePeg(width float32) *SquarePeg {
	return &SquarePeg{
		width,
	}
}

func (sp *SquarePeg) getWidth() float32 {
	return sp.width
}

// ==================

// RoundPegItf is
type RoundPegItf interface {
	getRadius() float32
}

// ==================

// SquarePegAdapter is
type SquarePegAdapter struct {
	*SquarePeg
}

// NewSquarePegAdapter is
func NewSquarePegAdapter(squarePeg *SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		squarePeg,
	}
}

func (spa *SquarePegAdapter) getRadius() float32 {
	return (spa.getWidth() * float32(math.Sqrt(2))) / 2
}
