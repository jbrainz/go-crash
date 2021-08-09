package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	confees := make(map[int]string)
	confees[1] = "Cappucino"
	confees[2] = "Latte"
	confees[3] = "Americano"
	confees[4] = "Mocha"
	confees[5] = "Machiatto"
	confees[6] = "Expresso"

	fmt.Println("MENU")
	fmt.Println("----------")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Machiatto")
	fmt.Println("6 - Expresso")
	fmt.Println("Q - Quit the program")

	fmt.Println("Press any key on the ketboard.")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}
		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You choose %s", confees[i]))

	}

	fmt.Println("Program Exiting...")

}
