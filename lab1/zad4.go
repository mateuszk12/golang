package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"math"
)
func randomNum(min,max int) int {
	var minNum int
	if min < 0 {
		minNum = -1*rand.Intn(int(math.Abs(float64(min))))
		return minNum+rand.Intn(max)
	}else {
		num := rand.Intn(max)
		if num > max {
			return max
		} else if num < min {
			return min
		} else {
			return num
		}
	}
	
}
func game(){
	random := int64(randomNum(-100,100))
	var odp string
	for {
		fmt.Print("podaj liczbę: ")
		fmt.Scan(&odp)
		if odp == "koniec"{
			fmt.Println("żegnaj")
			break
		} else {
			num,err := strconv.ParseInt(odp,10,32)
			 {
			if num > random {
				fmt.Println("za dużo!")
			} else if num < random {
				fmt.Println("za mało!")
			} else if num == random {
				fmt.Println("Brawo! zgadłeś")
				fmt.Println("Chcesz grać dalej?")
				var again string
				fmt.Scan(&again)
				if again == "tak"{
					random = int64(randomNum(-100,100))
				} else {
					fmt.Println("żegnaj")
					break
				}
			} else {
				fmt.Println(err)
			}
			}
		}
	}
}
func main(){
	fmt.Println("witaj")
	game()
}