/*
* Challenge 1
*
*	This example data contains six reports each containing five levels.

*	The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:
*
*		The levels are either all increasing or all decreasing.
*		Any two adjacent levels differ by at least one and at most three.
*
*	In the example above, the reports can be found safe or unsafe by checking those rules:
*
*		7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
*		1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
*		9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
*		1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
*		8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
*		1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
*
*	So, in this example, 2 reports are safe.
*
*	Analyze the unusual data from the engineers. How many reports are safe?

* 	Challenge 2
*	Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.
*
*	More of the above example's reports are now safe:
*
*	    7 6 4 2 1: Safe without removing any level.
*	    1 2 7 8 9: Unsafe regardless of which level is removed.
*	    9 7 6 2 1: Unsafe regardless of which level is removed.
*	    1 3 2 4 5: Safe by removing the second level, 3.
*	    8 6 4 4 1: Safe by removing the third level, 4.
*	    1 3 6 7 9: Safe without removing any level.
*
*	Thanks to the Problem Dampener, 4 reports are actually safe!
*
*	Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//open the file
	file, err := os.Open("./day2.txt")
	check(err)

	//close the file when ending
	defer file.Close()

	var data [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//Read curr line as string
		line := scanner.Text()

		//split the line into fields (space separated)
		stringNumbers := strings.Fields(line)

		//convert the striong values to int
		var numbers[]int

		//iterate over file
		for _, strNum := range stringNumbers {
			num,err := strconv.Atoi(strNum)
			check(err)
			numbers = append(numbers, num)
		}
		//complet data structure with the data
		data = append(data, numbers)
	}

	res := challengeOneResult(data)
	resChall2 := challengeTwoResult(data)

	fmt.Println("Challenge One res:", res)
	fmt.Println("Challenge Two res:", resChall2)
}

//Get bool value if all values are incresing in the line
func isIncreasing (line []int) bool {
	for i:= 1; i < len(line); i++ {
		if line[i] <= line[i-1] {
			return false
		}
	}
	return true
}

//Get bool value if all values are decreasing in the Line
func isDecreasing (line []int) bool {
	for i := 1; i < len(line); i++ {
		if line[i] >= line[i-1] {
			return false
		}
	}
	return true
}

//Get bool value if the is met
func hasValidDiff (line []int) bool {
	for i:= 1 ; i < len(line); i++ {
		diff := abs(line[i] - line[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

//get Abs value of integer
func abs (x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Get challenge One Result
func challengeOneResult (data [][]int) int {
	var finalCount int

	//Iterate over data to get the number od levels that are safe
	for _, line := range data {
		increasing := isIncreasing(line)
		decreasing := isDecreasing(line)
		validDifferences := hasValidDiff(line)

		if (increasing || decreasing) && validDifferences {
			finalCount++
		}
	}
	return finalCount
}

//Get the Challenge two answer
func challengeTwoResult(data [][]int) int {
	var finalCount int

	for _, line := range data {
		if isSafe(line) {
			finalCount++
			continue
		}
		// Check if removing one level makes the line safe
		for i := 0; i < len(line); i++ {
			reducedLine := removeIndex(line, i)
			if isSafe(reducedLine) {
				finalCount++
				break
			}
		}
		
	}
	return finalCount
}

// Helper function to determine if a report is safe
func isSafe(line []int) bool {
	// Check if the line is increasing or decreasing AND has valid differences
	return (isIncreasing(line) || isDecreasing(line)) && hasValidDiff(line) 
}

// Helper function to remove an index from a slice
func removeIndex(slice []int, index int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}