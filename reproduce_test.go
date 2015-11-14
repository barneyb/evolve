package goevolve

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestStableCase(t *testing.T) {
	rand.Seed(1980)
	parent := Genome([]int{5, 5, 5})
	child := Reproduce(&parent)
	if !reflect.DeepEqual(*child, Genome([]int{5, 6, 5})) {
		fmt.Printf("didn't mutate correctly: %v\n", *child)
		t.FailNow()
	}
}

func TestEmptyGenome(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			fmt.Println("didn't fail on zero-length genome")
			t.FailNow()
		} else {
			switch r.(type) {
			case InvalidGenomeError:
			default:
				fmt.Printf("incorrect panic: %v\n", r)
				t.FailNow()
			}
		}
	}()
	parent := Genome([]int{})
	Reproduce(&parent)
}

func TestLotsOfRounds(t *testing.T) {
	rand.Seed(1980)
	parent := Genome([]int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	child := Reproduce(&parent)
	for i := 0; i < 10000; i++ {
		child = Reproduce(child)
	}
	fmt.Printf("10000 rounds: %v\n", *child)
}
