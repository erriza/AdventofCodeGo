package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	//create a Regex expresion 
	re := regexp.MustCompile(`\bm(?:ul|xul)\((\d{1,3}),\s*(\d{1,3})\)`)

	//Iterate over the scanner object
	for scanner.Scan() {
		//get a line from the scanner object
		line := scanner.Text()
		fmt.Println("this is line", line)

		//apply the regex Expresion over the line
		matches := re.FindAllStringSubmatch(line, -1)
		fmt.Println("this is matches", matches)

		//Go through the values to sum the total
		for _, match := range matches {

			fmt.Println("this is match", match)

			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
	
		fmt.Println("this is num1", num1)
		fmt.Println("this is num2", num2)

			total += num1 * num2
		}
	}
	fmt.Println("total", total)
}