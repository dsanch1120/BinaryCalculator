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
var bits int

func pow(base int, exp int) int{
	output := 1
	for exp != 0 {
		output *= base
		exp -= 1
	}

	return output;
}

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

//Gets number of bits
func getBits() {
	//Clears Terminal
	CallClear()

	//Prompts user to input desired number of bits
	fmt.Println("Enter Number of Bits")

	//Loop that iterates until user gives correct input
	for {
		reader := bufio.NewReader(os.Stdin)

		//Gets user's input
		temp, _ := reader.ReadString('\n')
		temp = strings.Replace(temp, "\n", "", -1)
		int, err := strconv.Atoi(temp)

		//Ensures user gives a correct input
		if err == nil && (0 < int) {
			bits = int
			break
		} else {
			CallClear()
			fmt.Println("Invalid input")
			fmt.Println("Enter Number of Bits as an Integer")
		}
	}
	CallClear()
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

func decimalToBinary(num int, option int) string{
	var bin []int
	switch option {
	case 1:				//Unsigned binary
		for {
			if num <= 0 {
				break
			}
			bin = append(bin, num%2)
			num = num / 2
		}
		if len(bin) == 0 {
			return "000"
		}
		break
	case 2:
		if num < 0 {
			bin = append(bin, 1)
			num *= -1
		} else {
			bin = append(bin, 0)
		}
		for {
			if num <= 0 {
				break
			}
			bin = append(bin, num%2)
			num = num / 2
		}
		break
	}
	output := ""
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == 1 {
			output += string('1')
		} else {
			output += string('0')
		}
	}

	return output
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

	getBits()

	//Gets values of binary strings to be used
	bin1 := getStrings(1)
	bin2 := getStrings(2)

	bitCounter := 0
	dec := 0
	val1 := 0
	val2 := 0

	//Converts from binary to base 10
	for i := len(bin1) - 1; i >= 0; i-- {
		if bin1[i] == '1' {
			val1 += pow(2, bitCounter)
		}
		bitCounter++
	}
	bitCounter = 0
	for i := len(bin2) - 1; i >= 0; i-- {
		if bin2[i] == '1' {
			val2 += pow(2, bitCounter)
		}
		bitCounter++
	}

	//Performs the appropriate arithmetic
	if addition {
		dec = val1 + val2
	} else {
		dec = val1 - val2
	}

	//Converts from base 10 to binary
	output := decimalToBinary(dec, 1)

	//Displays output to user
	fmt.Println(output)
}

//Calculates for signed binary
func signed() {
	//Message to be displayed
	CallClear()
	if addition {
		fmt.Println("Signed Binary Addition")
	} else {
		fmt.Println("Signed Binary Subtraction")
	}

	fmt.Println()

	getBits()

	//Gets binary values to be used throughout function
	bin1 := getStrings(1)
	bin2 := getStrings(2)

	//Converts from signed binary to base 10
	bitCounter := 0
	dec := 0
	sign := 0
	val1 := 0
	val2 := 0

	//Converts from binary to base 10
	for i := len(bin1) - 1; i >= 0; i-- {
		if i == 0 {
			if bin1[0] == '0' {
				sign = 1
			} else {
				sign = -1
			}
			break
		}
		if bin1[i] == '1' {
			val1 += pow(2, bitCounter)
		}
		bitCounter++
	}
	val1 *= sign
	bitCounter = 0
	for i := len(bin2) - 1; i >= 0; i-- {
		if i == 0 {
			if bin2[0] == '1' {
				sign = 1
			} else {
				sign = -1
			}
			break
		}
		if bin2[i] == '1' {
			val2 += pow(2, bitCounter)
		}
		bitCounter++
	}
	val2 *= sign
	//Performs the appropriate arithmetic
	if addition {
		dec = val1 + val2
	} else {
		dec = val1 - val2
	}

	//Converts from base 10 to signed binary
	output := decimalToBinary(dec, 2)

	//Displays output to user
	fmt.Println(output)
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
