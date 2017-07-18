package pred

import (
	"testing"

	floc "github.com/workanator/go-floc"
)

func TestOrTrue(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		predicates := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			predicates[n] = alwaysTrue
		}

		p := Or(predicates...)

		if p(nil) == false {
			t.Fatalf("%s expects true with %d predicates", t.Name(), i)
		}
	}
}

func TestOrFalse(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		predicates := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			predicates[n] = alwaysFalse
		}

		p := Or(predicates...)

		if p(nil) == true {
			t.Fatalf("%s expects false with %d predicates", t.Name(), i)
		}
	}
}

func TestOrMixed(t *testing.T) {
	const max = 100

	for i := 2; i < max; i++ {
		predicates := make([]floc.Predicate, i)
		for n := 0; n < i; n++ {
			if n%2 == 0 {
				predicates[n] = alwaysTrue
			} else {
				predicates[n] = alwaysFalse
			}
		}

		p := Or(predicates...)

		if p(nil) == false {
			t.Fatalf("%s expects true with %d predicates", t.Name(), i)
		}
	}
}

func TestOrPanic(t *testing.T) {
	panicFunc := func(n int) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s must panic with %d predicates", t.Name(), n)
			}
		}()

		predicates := make([]floc.Predicate, n)
		for i := 0; i < n; i++ {
			predicates[n] = alwaysFalse
		}

		Or(predicates...)
	}

	panicFunc(0)
	panicFunc(1)
	panicFunc(2)
}
