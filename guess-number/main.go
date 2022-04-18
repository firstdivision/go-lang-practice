package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	var maxNumber = 100
	myNumber := rand.Intn(maxNumber)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess The Number")
	fmt.Println("---------------------")
	fmt.Printf("I'm thinking of a number between 1 and %d, can you guess it?\n", maxNumber)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		intVar, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Integers only please!")
			continue
		} else if (intVar <= 0) || (intVar > 100) {
			fmt.Println("Out of range! (0 to 100)")
			continue
		} else if intVar < myNumber {
			fmt.Println("Too low!")
			continue
		} else if intVar > myNumber {
			fmt.Println("Too high!")
			continue
		} else {
			fmt.Println("You got it!")
			break
		}
	}

}
