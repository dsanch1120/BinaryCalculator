/*
	Created by Daniel Sanchez
	July 8th, 2020
	Command Line Interface to calculate binary additions and substractions
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

//Map that will clear the terminal screen
var clear map[string]func()
var addition bool

//Allows the program to clear the terminal in Windows and Unix operating systems
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//Clears the Screen
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

//Gets binary strings
func getStrings(ind int) string {
	if ind == 1 {
		fmt.Println("Enter First Binary String")
	} else if ind == 2 {
		fmt.Println("Enter Second Binary String")
	} else {
		return ""
	}
	for {
		reader := bufio.NewReader(os.Stdin)

		//Gets user's input
		temp, _ := reader.ReadString('\n')
		temp = strings.Replace(temp, "\n", "", -1)

		for i := 0; i < len(temp); i++ {
			if temp[i] != '0' && temp[i] != '1' {
				break
			}

			if i == len(temp) - 1 {
				CallClear()
				return temp
			}
		}

		//Clears terminal for readability and displays updated message to user
		CallClear()
		fmt.Println(temp + " is not a Valid Input!")
		fmt.Println("Enter an unsigned Binary number")
		fmt.Println()
	}

}

//Calculates for unsigned binary
func unsigned() {
	//Message to be displayed at top
	CallClear()
	if addition {
		fmt.Println("Unsigned Binary Addition")
	} else {
		fmt.Println("Unsigned Binary Subtraction")
	}

	fmt.Println()

	bin1 := getStrings(1)
	bin2 := getStrings(2)

	fmt.Println(bin1)
	fmt.Println(bin2)
}

//Calculates for signed binary
func signed() {
	CallClear()
	if addition {
		fmt.Println("Signed Binary Addition")
	} else {
		fmt.Println("Signed Binary Subtraction")
	}

	fmt.Println()

	bin1 := getStrings(1)
	bin2 := getStrings(2)

	fmt.Println(bin1)
	fmt.Println(bin2)
}

//Allows the user to choose their conditions
func menu() {
	//Variables to be used throughout the function
	reader := bufio.NewReader(os.Stdin)
	var temp int
	var err error
	var funcCall map[int]func()
	funcCall = make(map[int]func()) //Allocated funcCall
	//Has funcCall call the various functions depending on the user's input
	funcCall[1] = func() {
		unsigned()
	}
	funcCall[2] = func() {
		signed()
	}

	//Displays welcome message and prompts user to choose between addition and subtraction
	fmt.Println("Welcome to the Binary Calculator")
	fmt.Println()
	fmt.Println("Select From the Following Options")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")

	//Loops until user gives a correct input
	for {
		//Gets user's input
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		temp, err = strconv.Atoi(text)

		//Ensures user gives a correct input
		if err == nil && temp == 1 {
			addition = true
			break
		}
		if err == nil && temp == 2 {
			addition = false
			break
		}

		//Clears terminal for readability and displays updated message to user
		CallClear()
		fmt.Println(text + " is not a Valid Input!")
		fmt.Println("Select From the Following Options")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
	}

	//Prompts user to choose type of binary string
	CallClear()
	if addition {
		fmt.Println("What type of binary strings are being added?")
	} else {
		fmt.Println("What type of binary strings are being subtracted?")

	}
	fmt.Println()
	fmt.Println("Select From the Following Options")
	fmt.Println("1. Unsigned")
	fmt.Println("2. Signed")

	//Loops until user gives a correct input
	for {
		//Gets user's input
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		temp, err = strconv.Atoi(text)

		//Ensures user gives a correct input
		if err == nil && (0 < temp && temp <= 2) {
			value, ok := funcCall[temp]
			if ok {
				value()
			}
			break
		}

		//Clears terminal for readability and displays updated message to user
		CallClear()
		fmt.Println(text + " is not a Valid Input!")
		fmt.Println("Select From the Following Options")
		fmt.Println("1. Unsigned")
		fmt.Println("2. Signed")
	}
}

//Handles the beginning stage of the program
func main() {
	//Clears screen for better readability
	CallClear()
	//Calls function to start program
	menu()
}
