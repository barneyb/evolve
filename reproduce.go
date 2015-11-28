package evolve

import (
	"math/rand"
)

var reproduceRand *rand.Rand

/*
ReproduceSeed seeds the shared randomness source for Reproduce invocations.  If
not invoked, the shared source will be initialized with 1, just like math/rand's
shared source.
*/
func ReproduceSeed(seed int64) {
	reproduceRand = rand.New(rand.NewSource(seed))
}

/*
Reproduce takes a *Genome and returns a new *Genome with a single point
mutation, using a shared source for randomness.
*/
func Reproduce(parent *Genome) *Genome {
	if reproduceRand == nil {
		ReproduceSeed(1)
	}
	return ReproduceRand(parent, reproduceRand)
}

/*
ReproduceRand takes a *Genome and returns a new *Genome with a single point
mutation, using the passed *Rand for randomness.
*/
func ReproduceRand(parent *Genome, rand *rand.Rand) *Genome {
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
