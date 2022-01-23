package main

import (
	"fmt"
	"os"
	"bufio"
)

type Cell struct{
	x int
	y int
}

var bestValues = make(map[Cell]int)
var unvisitedSet = make([]Cell, 0)

func readLines() ([][]int, error){
	content, err := os.Open(os.Args[1])
	if err != nil{
		var none [][]int
		return none, err
	}

	bf := bufio.NewScanner(content)

	var strMatrix []string
	for bf.Scan(){
		strMatrix = append(strMatrix, bf.Text())
	}

	var matrix [][]int

	for i:=0; i<len(strMatrix); i++{
		var slice = make([]int, len(strMatrix[i]))
		for j:=0; j<len(strMatrix[i]); j++{
			slice[j] = int(strMatrix[i][j])-48
		}
		matrix = append(matrix, slice)
	}

	return matrix, nil 
}

func findMin()(int,int, int){
	min := 999999999
	x := 0
	y := 0
	index := 0

	for i:=0; i<len(unvisitedSet); i++{
		if bestValues[unvisitedSet[i]] < min {
			min = bestValues[unvisitedSet[i]]
			x = unvisitedSet[i].x
			y = unvisitedSet[i].y
			index = i
		}
	}

	fmt.Println(x, ",", y, "\tmin =", min)

	return x, y, index
}

func removeFromSet(set []Cell, idx int) []Cell{
	set[idx] = set[len(set)-1]
	set = set[:len(set)-1]
	return set
}

func dijkstra(matrix [][]int){
	for len(unvisitedSet)>0{
		x, y, index := findMin()

		unvisitedSet = removeFromSet(unvisitedSet, index)
	
		if x == len(matrix)-1 && y == len(matrix[0])-1{
			return
		}
	
		if x > 0{
			val := bestValues[Cell{x, y}]+matrix[x-1][y]
			if val < bestValues[Cell{x-1, y}]{
				bestValues[Cell{x-1, y}] = val
				unvisitedSet = append(unvisitedSet, Cell{x-1, y})
			}
		}
		if x < len(matrix)-1{
			val := bestValues[Cell{x, y}]+matrix[x+1][y]
			if val < bestValues[Cell{x+1, y}]{
				bestValues[Cell{x+1, y}] = val
				unvisitedSet = append(unvisitedSet, Cell{x+1, y})
			}
		}
		if y > 0{
			val := bestValues[Cell{x, y}]+matrix[x][y-1]
			if val < bestValues[Cell{x, y-1}]{
				bestValues[Cell{x, y-1}] = val
				unvisitedSet = append(unvisitedSet, Cell{x, y-1})
			}
		}
		if y < len(matrix[0])-1{
			val := bestValues[Cell{x, y}]+matrix[x][y+1]
			if val < bestValues[Cell{x, y+1}]{
				bestValues[Cell{x, y+1}] = val
				unvisitedSet = append(unvisitedSet, Cell{x, y+1})
			}
		}
	}
}

func addTileHorizontally(matrix [][]int, left, right int) [][]int{
	for i:=0; i<len(matrix); i++{
		for j:=left; j<right; j++{
			val := matrix[i][j]+1
			if val > 9{
				val = 1
			}
			matrix[i] = append(matrix[i], val)
		}
	}

	return matrix
}

func addTileVertically(matrix [][]int, up, bot int) [][]int{
	for i:=up; i<bot; i++{
		sli := make([]int, len(matrix[i]))
		for j:=0; j<len(matrix[i]); j++{
			sli[j] = matrix[i][j]+1
			if sli[j] > 9{
				sli[j] = 1
			}
		}
		matrix = append(matrix, sli)
	}

	return matrix
}


func get5x5Matrix(matrix [][]int) [][]int {
	
	horizontalIncrement := len(matrix[0])
	verticalIncrement := len(matrix)

	bot := len(matrix)
	up := 0
	left := 0
	right := len(matrix[0])

	for i:=0; i<4; i++{
		matrix = addTileHorizontally(matrix, left, right)
		left += horizontalIncrement
		right += horizontalIncrement
	}

	for i:=0; i<4; i++{
		matrix = addTileVertically(matrix, up, bot)
		bot += verticalIncrement
		up += verticalIncrement
	}

	return matrix
}

func main(){
	matrix, err := readLines()
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %v", err))
		return
	}

	fmt.Printf("-----------First Tile------------\n\n")

	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]); j++{
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	matrix = get5x5Matrix(matrix)

	fmt.Printf("\n\n\n-----------Full Matrix(%d,%d)------------\n\n", len(matrix[0]), len(matrix))

	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]); j++{
			fmt.Printf("%d ", matrix[i][j])
			if i == 0 && j == 0{
				bestValues[Cell{i,j}] = matrix[0][0]
			} else{
				bestValues[Cell{i,j}] = 999999999
			}
		}
		fmt.Println()
	}

	unvisitedSet = append(unvisitedSet, Cell{0, 0})

	dijkstra(matrix)

	fmt.Println("Path Value =", bestValues[Cell{len(matrix)-1, len(matrix[0])-1}]-bestValues[Cell{0, 0}])
}
