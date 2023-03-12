package main

import (
	"math" 
	"fmt"
)

func main(){
	var a,b,c float64
	fmt.Println("podaj: wspóczynniki równania kwadratowego (ax^2+bx+x)")
	fmt.Scan(&a,&b,&c)
	var delta float64
	delta = (b*b) - 4*a*c
	if delta < 0 {
		fmt.Print("delta jest mniejsza od zera")
	}
	if delta == 0 {

	}
	if delta > 0 {
		var x1 float64
		var x2 float64
		x1 = (-b - math.Sqrt(delta))/ (2.0*a)
		x2 = (-b + math.Sqrt(delta))/ (2.0*a)
		fmt.Println("delta:",delta)
		fmt.Println("x1:",x1)
		fmt.Println("x2:",x2)
	}

}