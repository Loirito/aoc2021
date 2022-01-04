package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

var coordinates = make([][]int, 0)
var instructions = make([]string, 0)


func readLines() (int, int, error){
	content, err := os.Open(os.Args[1])
	if err != nil{
		return -1, -1, err
	}

	defer content.Close()

	bf := bufio.NewScanner(content)

	instruction_flag := 0
	maxX := 0
	maxY := 0
	for bf.Scan(){
		if bf.Text() == ""{
			instruction_flag = 1
		} else if instruction_flag == 0{
			coordinateSplit := strings.Split(bf.Text(), ",")
			x, _ := strconv.Atoi(coordinateSplit[0])
			y, _ := strconv.Atoi(coordinateSplit[1])
			values := []int{x, y}
			if x > maxX{
				maxX = x
			}
			if y > maxY{
				maxY = y
			}
			coordinates = append(coordinates, values)
		} else if instruction_flag == 1{
			instructionSplit := strings.Split(bf.Text(), " ")
			instruction := instructionSplit[2]
			instructions = append(instructions, instruction)
		}
	}

	return maxX, maxY, nil
}

func makePaper(maxX, maxY int) [][]string{
	matrix := make([][]string, maxY+1)
	for i:=0; i<len(matrix); i++{
		matrix[i] = make([]string, maxX+1)
	}

	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[0]); j++{
			matrix[i][j] = "."
		}
	}

	for i:=0; i<len(coordinates); i++{
		x := coordinates[i][0]
		y := coordinates[i][1]
		matrix[y][x] = "#"
	}

	return matrix
}

func foldPaper(matrix [][]string, lineOrCol string, value int) [][]string{
	if lineOrCol == "x"{ 		// If you want to fold by x
		maxVal := len(matrix[0])
		diff := 1
		for j:=value+1; j<maxVal; j++{
			for i:=0; i<len(matrix); i++{
				if matrix[i][j] != matrix[i][value-diff]{
					matrix[i][value-diff] = "#"
				}
			}
			diff++
		}

		for i:=0; i<len(matrix); i++{
			matrix[i] = matrix[i][:value]
		}

	} else if lineOrCol == "y"{	// If you want to fold by y
		maxVal := len(matrix)
		diff := 1
		for i:=value+1; i<maxVal; i++{
			for j:=0; j<len(matrix[0]); j++{
				if matrix[i][j] != matrix[value-diff][j]{
					matrix[value-diff][j] = "#"
				}
			}
			diff++
		}
		matrix = matrix[:value][:]
	}

	return matrix
}

func main(){
	maxX, maxY, err := readLines()
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %v\n", err))
		return
	}

	matrix := makePaper(maxX, maxY)

	fmt.Println("\nFolded paper")

	for i:=0; i<len(instructions); i++{
		firstInstruction := strings.Split(instructions[i], "=")
		value, _ := strconv.Atoi(firstInstruction[1])
		matrix = foldPaper(matrix, firstInstruction[0], value)
	}

	count := 0
	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[0]); j++{
			fmt.Printf("%s", matrix[i][j])
			if matrix[i][j] == "#"{
				count++
			}
		}
		fmt.Println()
	}


	fmt.Println("\nNumber of dots =", count)
}
