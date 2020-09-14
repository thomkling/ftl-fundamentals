package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	name string
	a    float64
	b    float64
	want float64
}

type testCaseWithErr struct {
	name        string
	a           float64
	b           float64
	errExpected bool
	want        float64
}

func TestAdd(t *testing.T) {
	cases := []testCase{
		{"two positive numbers", 2, 2, 4},
		{"two negative numbers", -2, -2, -4},
		{"one negative and one positive number equaling zero", -1, 1, 0},
		{"one fractional and one whole number", 5.4, 2, 7.4},
		{"two fractional numbers", 2.3, 4.3, 6.6},
		{"two fractional numbers equaling a whole number", 2.3, 3.7, 6},
		{"two negative fractional numbers", -1.5, -2.5, -4},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Add(c.a, c.b)
			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	cases := []testCase{
		{"two positive numbers", 5, 1, 4},
		{"two negative numbers", -5, -1, -4},
		{"two negative numbers equaling zero", -1, -1, 0},
		{"a fractional number from a whole number", 5.0, 1.5, 3.5},
		{"two fractional numbers", 2.3, 4.3, -2},
		{"two fractional numbers equaling a whole number", 2.5, 1.5, 1},
		{"two negative fractional numbers", -2.5, -1.5, -1},
		{"a negative number from a positive number", 5, -3, 8},
		{"a postitive number from a negative number", -4, 2, -6},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Subtract(c.a, c.b)
			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	cases := []testCase{
		{"two positive numbers", 2, 2, 4},
		{"two negative numbers", -2, -2, 4},
		{"a postive number by zero", 1, 0, 0},
		{"a fractional by a whole number", 5.4, 2, 10.8},
		{"two fractional numbers", 2.6, 5.3, 13.78},
		{"two fractional numbers equaling a whole number", 1.5, 4, 6},
		{"two negative fractional numbers", -1.5, -2.5, 3.75},
		{"a postive number by a negative number", -5, 3, -15},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Multiply(c.a, c.b)
			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	cases := []testCaseWithErr{
		{"two positive numbers", 4, 2, false, 2},
		{"two negative numbers", -4, -2, false, 2},
		{"a positive number by a negative number", 4, -2, false, -2},
		{"a negative number by a postive number", -4, 2, false, -2},
		{"two fractional positive numbers", 4.2, 2.1, false, 2},
		{"two fractional negative numbers equaling a whole number", -4.2, -2.1, false, 2},
		{"a positive number by zero", 4, 0, true, 0},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got, err := calculator.Divide(c.a, c.b)
			if err != nil && !c.errExpected {
				t.Errorf(c.name+": wanted %f got an error: %v ", want, err)
			}

			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Parallel()
	for i := 0; i < 100; i++ {
		a := random.Float64() * 100
		b := random.Float64() * 100
		name := fmt.Sprintf("Adding %f and %f", a, b)
		result := a + b
		c := testCase{name, a, b, result}

		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Add(c.a, c.b)
			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}

}

func TestSqrt(t *testing.T) {
	cases := []testCaseWithErr{
		{"Square root of zero should be zero", 0, 0, false, 0},
		{"Square root of a negative number should return an error", -4, 0, true, 0},
		{"Square root of a positive whole number", 4, 0, false, 2},
	}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got, err := calculator.Sqrt(c.a)
			if err != nil && !c.errExpected {
				t.Errorf(c.name+": wanted %f got an error: %v ", want, err)
			}

			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		a := random.Float64() * 100
		name := fmt.Sprintf("Square root of %f", a)
		result := math.Sqrt(a)
		c := testCaseWithErr{name, a, 0, false, result}

		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got, err := calculator.Sqrt(c.a)
			if err != nil && !c.errExpected {
				t.Errorf(c.name+": wanted %f got an error: %v ", want, err)
			}

			if want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}
