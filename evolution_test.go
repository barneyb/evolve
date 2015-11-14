package evolve

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	g := Genome([]int{1, 2, 3})
	dev := func(*Genome) *Individual {
		return nil
	}
	e := New(&g, dev)
	if e.Latest == nil {
		fmt.Println("latest wasn't set")
		t.FailNow()
	}
	if !reflect.DeepEqual(*e.Latest, g) {
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
}

func TestEvolve(t *testing.T) {
	rand.Seed(1980)
	g := Genome([]int{0, 0, 0, 0})
	dev := func(g *Genome) *Individual {
		return &Individual{
			Genotype: g,
		}
	}
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
	dev := func(g *Genome) *Individual {
		return &Individual{
			Genotype: g,
		}
	}
	e := New(&g, dev)
	ng := e.Evolve(1)[0]
	e.Select(&ng)
	if !reflect.DeepEqual(ng.Genotype, e.Latest) {
		fmt.Printf("latest wasn't updated: %v vs %v\n", ng.Genotype, e.Latest)
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
