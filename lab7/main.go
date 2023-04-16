package main

import (
	. "flag"
	. "fmt"
	. "math"
)

type r struct{ a, b, c float64 }

func (l *r) q() (float64, float64, bool) {
	d := Pow(l.a*l.b, 2) - 4*l.c
	r, x := Sqrt(d), l.a*2
	return (-l.b + r) / x, (-l.b - r) / x, d >= 0
}
func main() {
	a, b, c := Float64("a", 0, ""), Float64("b", 0, ""), Float64("c", 0, "")
	Parse()
	t := r{a: *a, b: *b, c: *c}
	Println(t.q())
}
