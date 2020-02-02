package main

import (
	"fmt"
	"strconv"
	"os"
	"math/rand"
	"time"
)

func stoI (s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
    }
    return i
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("How many sides do you want your dice to have?")
	var sidePrompt string
	fmt.Scanln(&sidePrompt)
	fmt.Println("And how many times do you want to roll it?")
	var rollPrompt string
	fmt.Scanln(&rollPrompt)
	
	sides := stoI(sidePrompt)
	rolls := stoI(rollPrompt)
	

	fmt.Println("=================")
	for i := 0; i < rolls; i++{
		fmt.Println(rand.Intn(sides) + 1)
	}
	fmt.Println("=================")



}
