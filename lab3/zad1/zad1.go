package zad1

import (
	"fmt"
)
func Collatz(min int64,max int64) {
	var number int64
	for i := min ; i <= max; i++{
		number = int64(i)
		for number != 1 {
			if number % 2 == 0{
				number = number / 2
				fmt.Print(number) 
			} else {
				number = (number * 3) + 1
				fmt.Print(number) 
			}
			fmt.Print(" ")
		}
		fmt.Print("-> ")
		fmt.Print(i)
		fmt.Println("")
	}
}