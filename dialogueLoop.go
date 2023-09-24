package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputReader = bufio.NewReader(os.Stdin)

var toDoList []string

type dialogueState int

const QUIT dialogueState = 0
const HOME dialogueState = 1
const ADD dialogueState = 2
const DELETE dialogueState = 3

func getDialogueOption(nOptions int) int {
	var option int
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("Please enter a number. ")
		return getDialogueOption(nOptions)
	}
	input = input[:len(input)-1]
	option, err = strconv.Atoi(input)
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

func showHome() dialogueState {
	fmt.Println("Welcome, your to-do list currently contains the following items:")
	for _, v := range toDoList {
		fmt.Println("-", v)
	}
	fmt.Println("What would you like to do?")
	var chosenOption int
	chosenOption = showChoiceMenu([]string{"Add an item", "Remove an item", "Quit"})
	switch chosenOption {
	case 1:
		return ADD
	case 2:
		return DELETE
	default:
		return QUIT
	}
}

func showAdd() dialogueState {
	fmt.Println("What item would you like to add?")
	var newItem string
	newItem, err := inputReader.ReadString('\n')
	newItem = newItem[:len(newItem)-1]
	if err != nil {
		fmt.Println("Failed to add your new item, returning to home screen.")
		return HOME
	}
	toDoList = append(toDoList, newItem)
	println("Added the item \"" + newItem + "\" to your list, do you want to add another item?")
	var chosenOption int
	chosenOption = showChoiceMenu([]string{"Yes", "No"})
	switch chosenOption {
	case 1:
		return ADD
	default:
		return HOME
	}
}

func showDelete() dialogueState {
	fmt.Println("Which item would you like to delete?")
	i := showChoiceMenu(toDoList) - 1
	item := toDoList[i]
	toDoList = append(toDoList[:i], toDoList[i+1:]...)
	fmt.Println("Succesfully removed item \"", item, "\".")
	fmt.Println("Would you like to delete another element?")
	chosenOption := showChoiceMenu([]string{"Yes", "No"})
	switch chosenOption {
	case 1:
		return DELETE
	default:
		return HOME
	}
}

func main() {
	var state = HOME
	for state != QUIT {
		switch state {
		case HOME:
			state = showHome()
		case ADD:
			state = showAdd()
		case DELETE:
			state = showDelete()
		}
	}
}
