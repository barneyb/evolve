package goevolve

import (
	"math/rand"
)

/*
Reproduce takes a *Genome and returns a new *Genome with a single point
mutation.
*/
func Reproduce(parent *Genome) *Genome {
	if len(*parent) == 0 {
		panic(InvalidGenomeError(*parent))
	}
	child := make(Genome, len(*parent))
	copy(child, *parent)
	i := rand.Intn(len(child))
	pos := rand.Float32() < 0.5
	if pos {
		child[i] += 1
	} else {
		child[i] -= 1
	}
	return &child
}
