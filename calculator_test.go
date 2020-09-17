package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"two positive numbers", 2, 2, 4},
		{"two negative numbers", -2, -2, -4},
		{"one negative and one positive number equaling zero", -1, 1, 0},
		{"one fractional and one whole number", 5.4, 2, 7.4},
		{"two fractional numbers", 2.3, 4.3, 6.6},
		{"two fractional numbers equaling a whole number", 2.3, 3.7, 6},
		{"two negative fractional numbers", -1.5, -2.5, -4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Add(c.a, c.b)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Subtract(c.a, c.b)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{"two positive numbers", 2, 2, 4},
		{"two negative numbers", -2, -2, 4},
		{"a postive number by zero", 1, 0, 0},
		{"a fractional by a whole number", 5.4, 2, 10.8},
		{"two fractional numbers", 2.6, 5.3, 13.78},
		{"two fractional numbers equaling a whole number", 1.5, 4, 6},
		{"two negative fractional numbers", -1.5, -2.5, 3.75},
		{"a postive number by a negative number", -5, 3, -15},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got := calculator.Multiply(c.a, c.b)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		name        string
		a           float64
		b           float64
		want        float64
		errExpected bool
	}{
		{"two positive numbers", 4, 2, 2, false},
		{"two negative numbers", -4, -2, 2, false},
		{"a positive number by a negative number", 4, -2, -2, false},
		{"a negative number by a postive number", -4, 2, -2, false},
		{"two fractional positive numbers", 4.2, 2.1, 2, false},
		{"two fractional negative numbers equaling a whole number", -4.2, -2.1, 2, false},
		{"a positive number by zero", 4, 0, 0, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got, err := calculator.Divide(c.a, c.b)
			errReceived := (err != nil)
			if c.errExpected != errReceived {
				if c.errExpected {
					t.Fatalf(c.name+": wanted an error and instead got: %f", got)
				}
				t.Fatalf(c.name+": wanted %f got an error: %v ", want, err)
			}

			if !c.errExpected && want != got {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		a := random.Float64() * 100
		b := random.Float64() * 100
		name := fmt.Sprintf("Adding %f and %f", a, b)
		result := a + b
		c := struct {
			name string
			a    float64
			b    float64
			want float64
		}{name, a, b, result}

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
	cases := []struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}{
		{"Square root of zero should be zero", 0, 0, false},
		{"Square root of a negative number should return an error", -4, 0, true},
		{"Square root of a positive whole number", 4, 2, false},
		{"Square root of NaN should return an error", math.NaN(), 0, true},
		{"Square root of negative infinity should return an error", math.Inf(-1), 0, true},
		{"Square root of positive infinity should be positive infinity", math.Inf(1), math.Inf(1), false},
		{"Square root of 5.3", 5.3, 2.302173, false},
		{"Square root of 12.2", 12.2, 3.492850, false},
		{"Square root of 987.345", 987.345, 31.422046, false},
		{"Square root of 123.098", 123.098, 11.094954, false},
		{"Square root of 0.456", 0.456, 0.675278, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := c.want
			got, err := calculator.Sqrt(c.a)
			errReceived := (err != nil)
			if c.errExpected != errReceived {
				if c.errExpected {
					t.Fatalf(c.name+": wanted an error and instead got: %f", got)
				}
				t.Fatalf(c.name+": wanted %f got an error: %v ", want, err)
			}

			if !c.errExpected && !(compareFloat64(want, got, 0.000001)) {
				t.Errorf(c.name+": want %f, got %f", want, got)
			}
		})
	}
}

func compareFloat64(a, b, tolerance float64) bool {
	if math.IsInf(a, 1) && math.IsInf(b, 1) {
		return true
	}

	diff := math.Abs(a - b)
	return diff < tolerance
}
