/*
* 	--- Day 4: Ceres Search ---
*
* 	"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it. After a brief flash, you recognize the interior of the Ceres monitoring station!
*
* 	As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if you could help her with her word search (your puzzle input). She only has to find one word: XMAS.
*
* 	This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them. Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:
*
* 	..X...
* 	.SAMX.
* 	.A..A.
* 	XMAS.S
* 	.X....
*
* 	The actual word search will be full of letters instead. For example:
*
* 	MMMSXXMASM
* 	MSAMXMSMSA
* 	AMXSXMAAMM
* 	MSAMASMSMX
* 	XMASAMXAMM
* 	XXAMMXXAMA
* 	SMSMSASXSS
* 	SAXAMASAAA
* 	MAMMMXMMMM
* 	MXMXAXMASX
*
* 	In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved in any XMAS have been replaced with .:
*
* 	....XXMAS.
* 	.SAMXMS...
* 	...S..A...
* 	..A.A.MS.X
* 	XMASAMX.MM
* 	X.....XA.A
* 	S.S.S.S.SS
* 	.A.A.A.A.A
* 	..M.M.M.MM
* 	.X.X.XMASX
*
* 	Take a look at the little Elf's word search. How many times does XMAS appear?
*
* --- Part Two ---
*
* The Elf looks quizzically at you. Did you misunderstand the assignment?
*
* Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:
*
* M.S
* .A.
* M.S
*
* Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.
*
* Here's the same example from before, but this time all of the X-MASes have been kept instead:
*
* .M.S......
* ..A..MSMS.
* .M.S.MAA..
* ..A.ASMSM.
* .M.S.M....
* ..........
* S.S.S.S.S.
* .A.A.A.A..
* M.M.M.M.M.
* ..........
*
* In this example, an X-MAS appears 9 times.
*
* Flip the word search from the instructions back over to the word search side and try again. How many times does an X-MAS appear?
*
*
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

//create directions to better identify where should we look up for next
var directions [][]int = [][]int {
	[]int{0, -1},   // up
	[]int{0, 1},  // down
	[]int{1, 0},   // right
	[]int{-1, 0},  // left
	[]int{1, -1},   // up right
	[]int{-1, -1},  // up left
	[]int{1, 1},  // down right
	[]int{-1, 1}, // down left
}
//The characters to look up	
var wordList []string = []string{"X", "M", "A", "S"}

//lines to better create final matrixs
var lines []string

//create matrix to process
var matrix[][]string


func main () {
	//Read file and create line array
	file := ReadFile("./day4.txt")
	//create 2D matrix to look for chars
	grid := createGrid(file)
	//Iterate over 2D matrix and sum the num of coincidences
	score := TraverseGrid2(grid)
 	fmt.Println("totalcount:", score)

}


func ReadFile (path string) []string {
	//Open file
	file, err := os.Open(path)
	check(err)

	//Close file
	defer file.Close()

	//Scanner to read the file
	scanner := bufio.NewScanner(file)

	//Iterate over scanner to create array of strings
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines

}

func createGrid(lines []string) [][]string {
	//iterate over lines array to create the 2D array matrix to look for the chars
	for _, line := range lines {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}
	return matrix
}

//Challenge 1
// func TraverseGrid(grid [][]string) int {
// 	score := 0
// 	for x, row := range grid {
// 		for y, char := range row {
// 			if char == wordList[0] {
// 				for _, direction := range directions {
// 					if findXMAS(x, y, 1, direction, grid) {
// 						score += 1
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return score
// }

//Challenge 1 FIND XMAS
func findXMAS(x, y, wordPosition int, direction []int, grid[][]string) bool {
	xNext := x + direction[0]
	yNext := y + direction[1]

	if wordPosition > len(wordList)-1 {
		return true
	}

	if xNext < 0 || xNext >= len(grid) || yNext < 0 || yNext >= len(grid[x]) {
		return false
	}

	if grid[xNext][yNext] == wordList[wordPosition] {
		return findXMAS(xNext, yNext, wordPosition+1, direction, grid)
	}
	return false
	
}

//Challenge 2 find MAS
func FindMAS(x, y int, grid [][]string, wordList []string) bool {
    xL, yT := x-1, y+1 // Top-left neighbor
    xR, yD := x+1, y-1 // Bottom-right neighbor

	//check if indices are within bounds
	if xL < 0 || xR >= len(grid) || yT < 0 || yD < 0 || yT >= len(grid[xL]) || yD >= len(grid[xR]) {
		return false
	}

	//declaring coordinates
	topLeft := grid[xL][yT]
	bottomRight := grid[xR][yD]
	topRight := grid[xR][yT]
	bottomLeft := grid[xL][yD]

	//getting letters M and S to look for
	word1, word3 := wordList[1], wordList[3]

	//looking
	isDiagonalMatch := (topLeft == word1 && bottomRight == word3) || (topLeft == word3 && bottomRight == word1)
    isAntiDiagonalMatch := (topRight == word1 && bottomLeft == word3) || (topRight == word3 && bottomLeft == word1)

	return isDiagonalMatch && isAntiDiagonalMatch
}

//Challenge 2
func TraverseGrid2(grid [][]string) int {
    score := 0
    for x, row := range grid {
        for y, char := range row {
            if char == wordList[2] {
                if FindMAS(x, y, grid, wordList) {
                    score += 1
                }

            }
        }
    }
    return score
}
