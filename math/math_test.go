package math

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestSubtract(t *testing.T){

    got := Subtract(10, 6)
    want := 4

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestMultiply(t *testing.T){

    got := Multiply(5, 5)
    want := 25

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestSquare(t *testing.T) {
    got := Square(5)
    want := 25

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}

func TestCube(t *testing.T) {
    got := Cube(5)
    want := 125

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}
