package evolve

import (
	"fmt"
)

/*
Genome is an []int that represents the "genes" of the "organism" being
evolved.
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
