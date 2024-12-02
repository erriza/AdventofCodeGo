package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//open the file
	file, err := os.Open("./day1.txt")
	check(err)

	//close file when ending
	defer file.Close()

	//create slices and variables to store data
	var data1 []int
	var data2 []int
	var sum int

	//scanner ??
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//split the line into two numbers
		line := scanner.Text()
		numbers := strings.Fields(line)

		if len(numbers) == 2 {
			var num1, num2 int
			fmt.Sscanf(numbers[0], "%d", &num1)
			fmt.Sscanf(numbers[1], "%d", &num2)

			//fill slices with some data
			data1 = append(data1, num1)
			data2 = append(data2, num2)
		}
	}

	//check if error in Scanner
	if err:= scanner.Err(); err !=  nil {
		fmt.Println("Error reading file", err)
		return
	}
	
	//Sort slices
	sort.Slice(data1, func(i, j int) bool {
		return data1[i] < data1[j]
	})

	sort.Slice(data2, func(i, j int) bool {
		return data2[i] < data2[j]
	})

	//Slices sorted
	// fmt.Println("Sorted Data 1:")
	// for _, val := range data1 {
	// 	fmt.Println(val)
	// }

	// fmt.Println("Sorted Data 2:")
	// for _, val := range data2 {
	// 	fmt.Println(val)
	// }

	//validate the same length
	if len(data1) == len(data2) {
		for i := range data1 {
			diff := diffValues(data1[i], data2[i])
			//sum diff
			sum += diff
		}		
	}

	//result of part1
	fmt.Println(sum)

	partTwo(data1, data2)

}
//Get difference between the numbers
func diffValues(a, b int) int  {
	if a < b {
		return b-a
	}
	return a-b
}

/* 
* Part 2
*/
func partTwo(data1, data2 []int) int {
	/*
	* Create variables for part2
	* similarityScore is the final answer
	* countRepeat is the number of times the number in the left data1 repeats in the right data2
	* */
	var similarityScore int 
	var countRepet int

	if len(data1) == len(data2) {
		for i := range data1 {
			countRepet = checkExist(data1[i], data2)
			totalNumber := countRepet * data1[i]
			similarityScore+= totalNumber
		}
		//Final Result for day 2
		fmt.Println(similarityScore)
	}
	return similarityScore
}

/*
* Function that returns the number of times a value repeats in a slice
*/
func checkExist(value1 int, data2 []int) int {
	var count int
	for _, v := range data2 {
		if value1 == v {
			count++
		}
	}
	return count
}