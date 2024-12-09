/*
* --- Day 3: Mull It Over ---
*
* "Our computers are having issues, so I have no idea if we have any Chief Historians in stock! You're welcome to check the warehouse, though," says the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The Historians head out to take a look.
*
* The shopkeeper turns to you. "Any chance you can see why our computers are having issues again?"
*
* The computer appears to be trying to run a program, but its memory (your puzzle input) is corrupted. All of the instructions have been jumbled up!
*
* It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.
*
* However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.
*
* For example, consider the following section of corrupted memory:
*
* xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
*
* Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).
*
* Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add up all of the results of the multiplications?
*
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Open the file
	file, err := os.Open("./day3.txt")
	check(err)

	//close the file when finisth
	defer file.Close()

	//Scanner to read the file
	scanner := bufio.NewScanner(file)

	//create the variable to sum the values
	var total int
	var totalChallenge2 int

	//create a Regex expresion for mul(X, Y)
	re := regexp.MustCompile(`\bm(?:ul|xul)\((\d{1,3}),\s*(\d{1,3})\)`)

	//create Regex expresion for do
	reDo := regexp.MustCompile(`\bdo\(\)`)

	//create Regex expresion for dont
	reDont := regexp.MustCompile(`\bdon't\(\)`)

	//Iterate over the scanner object
	for scanner.Scan() {
		//get a line from the scanner object
		line := scanner.Text()

		//apply the regex Expresion over the line
		matches := re.FindAllStringSubmatch(line, -1)

		//Go through the values to sum the total
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			total += num1 * num2
		}
	}
	fmt.Println("total", total)

	//Re open the file again for new scanner
	file2, err := os.Open("./day3.txt")
	check(err)
	defer file2.Close()

	//Reset the scanner for challenge 2
	scanner2 := bufio.NewScanner(file2)
    var mulEnabled = true

    // Loop through scanner lines (challenge 2)
	for scanner2.Scan() {
		line := scanner2.Text()

		// matches := re.FindAllStringSubmatch(line, -1)
		words:= strings.Split(line, " ")

		for _, word := range words {
			if reDo.MatchString(word) {
					mulEnabled = true
			} else if reDont.MatchString(word) {
					mulEnabled = false
			} else if re.MatchString(word) {
					if mulEnabled {
							// Extract numbers from the mul instruction
							num1, _ := strconv.Atoi(re.FindStringSubmatch(word)[1])
							num2, _ := strconv.Atoi(re.FindStringSubmatch(word)[2])
							total += num1 * num2
					}
			}
	}
}

fmt.Println("Total:", totalChallenge2)
}