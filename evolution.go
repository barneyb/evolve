package evolve

import (
	"fmt"
)

/*
Genome stores an isolated set of genes. Think a set of chromosomes.
*/
type Genome []int

/*
InvalidGenomeError is used to indicate a Genome that is invalid (e.g., has
length zero).
*/
type InvalidGenomeError Genome

func (g InvalidGenomeError) Error() string {
	return fmt.Sprintf("Invalid genome: %v", g)
}

/*
Individual is a genotype:phenotype tuple. Think a body with DNA in it's cells.
*/
type Individual struct {
	Genotype  *Genome
	Phenotype *interface{}
}

/*
Evolution is the process by which Individuals reproduce and develop different
Individuals with slightly different characteristics based on random mutation
and natural selection.  This is a highly simplified version of Evolution where
reproduction happens from a single individual (asexual) and each generation
has only a single survivor which passes it's genes on.
*/
type Evolution struct {
	Latest      *Individual
	Ancestry    []Genome
	Development func(*Genome) *Individual
}

/*
New starts a new Evolution with the given starting Genome and a development
function that can take a Genome and create an Individual from it.
*/
func New(start *Genome, development func(*Genome) *Individual) *Evolution {
	return &Evolution{
		development(start),
		[]Genome{},
		development,
	}
}

/*
Evolve creates a new generation of Individuals of the specified size from the
most recently selected Individual.  This is the "random mutation" part of the
evolutionary process.  Once Individual from the returned generation should be
selected as the survivor for the following generation to evolve from.
*/
func (e *Evolution) Evolve(size int) []Individual {
	nextGen := make([]Individual, size)
	for i := 0; i < size; i++ {
		nextGen[i] = *e.Development(Reproduce(e.Latest.Genotype))
	}
	return nextGen
}

/*
Select indicates the Individual from a given generation that survivied.  This
is the "natural selection" part of the evolutionary process.  Only one
Individual should be selected per generation.
*/
func (e *Evolution) Select(survivor *Individual) {
	e.Ancestry = append(e.Ancestry, *e.Latest.Genotype)
	e.Latest = survivor
}
