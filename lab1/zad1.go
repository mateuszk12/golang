package main

import (
	"fmt"
)
var cos int = 5
func main(){
	var age int
	var st string 
	var d float64
	fmt.Println(st,d,cos)
	fmt.Println("Podaj swój wiek: ")
	fmt.Scan(&age)
	months := age*12
	days := age*365
	fmt.Println(months,days)
	fmt.Println("lat na marsie: ",1,881*age)
}