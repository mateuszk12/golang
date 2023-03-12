package main


import(
	"fmt"
)

func main(){
	const i int = 5
	switch{
	case i > 4:
		fmt.Println("git")
	case i < 4:
		fmt.Println("niegit")
	default:
		fmt.Println("cokolwiek")
	}

}