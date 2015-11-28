package evolve

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func dev(g *Genome) *Individual {
	return &Individual{
		Genotype: g,
	}
}

func TestNew(t *testing.T) {
	g := Genome([]int{1, 2, 3})
	e := New(&g, dev)
	if e.Latest == nil {
		fmt.Println("latest wasn't set")
		t.FailNow()
	}
	if !reflect.DeepEqual(*e.Latest.Genotype, g) {
		fmt.Println("latest was set incorrectly")
		t.FailNow()
	}
	if len(e.Ancestry) != 0 {
		fmt.Println("ancestry isn't empty")
		t.FailNow()
	}
	if e.Development == nil {
		fmt.Println("didn't save Development")
		t.FailNow()
	}
	if e.Rand == nil {
		fmt.Println("didn't initialize Rand")
		t.FailNow()
	}
}

func TestNewRand(t *testing.T) {
	g := Genome([]int{1, 2, 3})
	r := rand.New(rand.NewSource(12345))
	e := NewRand(&g, dev, r)
	if e.Rand != r {
		fmt.Println("didn't save Rand")
		t.FailNow()
	}
}

func TestEvolve(t *testing.T) {
	rand.Seed(1980)
	g := Genome([]int{0, 0, 0, 0})
	e := New(&g, dev)
	ng := e.Evolve(3)
	if ng == nil {
		fmt.Println("next generation was nil")
		t.FailNow()
	}
	if len(ng) != 3 {
		fmt.Printf("wrong generation size: %v\n", len(ng))
		t.FailNow()
	}
	for i, ind := range ng {
		if reflect.DeepEqual(g, ind.Genotype) {
			fmt.Printf("child %v didn't mutate\n", i)
			t.FailNow()
		}
	}
}

func TestSelect(t *testing.T) {
	rand.Seed(1980)
	g := Genome([]int{0, 0, 0, 0})
	e := New(&g, dev)
	ng := e.Evolve(1)[0]
	e.Select(&ng)
	if ng != *e.Latest {
		fmt.Printf("latest wasn't updated: %v vs %v\n", ng, *e.Latest)
		t.FailNow()
	}
	if len(e.Ancestry) != 1 {
		fmt.Printf("ancestry wasn't appended to\n")
		t.FailNow()
	}
	if !reflect.DeepEqual(e.Ancestry[0], g) {
		fmt.Printf("ancestry had the wrong thing added\n")
		t.FailNow()
	}
}
