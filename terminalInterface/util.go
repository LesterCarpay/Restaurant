package terminalInterface

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getUserInput() string {
	var inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	for err != nil {
		fmt.Println("Failed to read input, please try again.")
		input, err = inputReader.ReadString('\n')
	}
	return input[:len(input)-1]
}

func getDialogueOption(nOptions int) int {
	var option int
	input := getUserInput()
	option, err := strconv.Atoi(input)
	if err != nil || option < 1 || option > nOptions {
		fmt.Println("Invalid choice. Please choose one of the", nOptions, "options.")
		return getDialogueOption(nOptions)
	} else {
		return option
	}
}

func showChoiceMenu(options []string) int {
	for i, option := range options {
		fmt.Println(strconv.Itoa(i+1)+".", option+".")
	}
	return getDialogueOption(len(options))
}

func scanWithDefault(name string, defaultValue string) string {
	fmt.Print(name + " [" + defaultValue + "]:")
	var result string
	_, err := fmt.Scanln(&result)
	if err != nil || result == "" {
		return defaultValue
	}
	return result
}

func GetConnectionString() string {
	var password string

	host := scanWithDefault("Host", "localhost")
	dbName := scanWithDefault("Database", "restaurant")
	username := scanWithDefault("Username", "postgres")
	fmt.Print("Password:")
	_, err := fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Incorrect password.")
		log.Fatalln(err)
	}
	return "postgresql://" +
		username + ":" +
		password + "@" +
		host + "/" +
		dbName + "?sslmode=disable"
}
