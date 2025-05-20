package utilFolder

import (
	"fmt"
)

//-- completely useless because I am supposed to be asking the client lol
// func AskForMangaName() string {
// 	var manga string = ""
// 	fmt.Println("What manga would you like to look at? (make sure if you want to type with spaces,, instead of spaces use dashes, '-'")
// 	fmt.Scanln(&manga) //gets user input lol
// 	return manga
// }

func AskPlan() {
	fmt.Println("Okay so what would you like to do?")
	var option int = 0
	showListOfThings()
	fmt.Scanln(&option)

	//use a switch case here maybe lol
}

// figure out a better way to call this ngl lol
func showListOfThings() {
	fmt.Println("(1). Show Synopsis")
	fmt.Println("(2). Give Recommendations")
	fmt.Println("(any other integers). just simply echo the title right now lol")
}
