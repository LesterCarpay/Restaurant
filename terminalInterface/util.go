package terminalInterface

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// getUserInput reads the standard input until a newline is found and returns it without the newline.
func getUserInput() string {
	var inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	for err != nil {
		fmt.Println("Failed to read input, please try again.")
		input, err = inputReader.ReadString('\n')
	}
	return input[:len(input)-1]
}

// getDialogueOption accepts an int nOptions, reads the standard input and, if it read a number between 
// 1 and nOptions, returns it, and if not, rereads the standard input until a valid number has been read.
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

// showChoiceMenu accepts a slice of strings, displays them as a list of options, reads the user input and
// returns the chosen option.
func showChoiceMenu(options []string) int {
	for i, option := range options {
		fmt.Println(strconv.Itoa(i+1)+".", option+".")
	}
	return getDialogueOption(len(options))
}

// scanWithDefault accepts name of the variable the user needs to input and a default value, reads the
// standard input and returns the user input or the default value if only a newline has been entered.
func scanWithDefault(varName string, defaultValue string) string {
	fmt.Print(varName + " [" + defaultValue + "]:")
	var result string
	_, err := fmt.Scanln(&result)
	if err != nil || result == "" {
		return defaultValue
	}
	return result
}

// GetConnectionString asks the user for the necessary input values in the terminal, chains the inputs together
// into a connection string and returns it.
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
