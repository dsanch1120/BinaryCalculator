/*
	Created by Daniel Sanchez
	July 8th, 2020
	Command Line Interface to calculate binary additions and substractions
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//Map that will clear the terminal screen
var clear map[string]func()
var addition bool

//Allows the program to clear the terminal in Windows and Unix operating systems
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd:=exec.Command("clear")
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

//Allows the user to choose their conditions
func menu() {
	fmt.Println("Welcome to the Binary Calculator")
	fmt.Println()
	fmt.Println("Select From the Following Options")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
}

//Handles the beginning stage of the program
func main() {
	//Clears screen for better readability
	CallClear()
	//Calls function to start program
	menu()
}