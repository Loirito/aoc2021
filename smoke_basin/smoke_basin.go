package main

import (
	"fmt"
	"os"
	"bufio"
)

type coordinate struct{
	x int
	y int
	next *coordinate
}

var head *coordinate = nil;


func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %s", err).Error())
		return
	}

	intMatrix := readLines(content)

	fmt.Println("Risk levels sum =", findLowestPoints(intMatrix))

	var markedMatrix = make([][]int, len(intMatrix))
	for i:=0; i<len(markedMatrix); i++{
		markedMatrix[i] = make([]int, len(intMatrix[0]))
	}

	best := 0
	bestest := 0
	besterest := 0

	for aux:=head; aux != nil; aux=aux.next{
		size := findBasinSize(aux.x, aux.y, intMatrix, markedMatrix) 
		if size > best{
			best = size
			if size > bestest{
				best = bestest
				bestest = size
				if size > besterest{
					bestest = besterest
					besterest = size
				}
			}
		}
	}

	fmt.Printf("Three largest basins (%d, %d, %d) multiplied = %d\n", best, bestest, besterest, best*bestest*besterest)

}


func findLowestPoints(matrix [][]int) int{
	sum := 0
	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]); j++{
			if matrix[i][j] != 9{
				lowest := 9
				if i==0{
					if j==0{
						lowest = minimumValue(lowest, matrix[i+1][j])
						lowest = minimumValue(lowest, matrix[i][j+1])
					} else if j==len(matrix[i])-1{
						lowest = minimumValue(lowest, matrix[i+1][j])
						lowest = minimumValue(lowest, matrix[i][j-1])
					} else {
						lowest = minimumValue(lowest, matrix[i+1][j])
						lowest = minimumValue(lowest, matrix[i][j-1])
						lowest = minimumValue(lowest, matrix[i][j+1])
					}
				} else if i==len(matrix)-1{
					if j==0{
						lowest = minimumValue(lowest, matrix[i-1][j])
						lowest = minimumValue(lowest, matrix[i][j+1])
					} else if j==len(matrix[i])-1{
						lowest = minimumValue(lowest, matrix[i-1][j])
						lowest = minimumValue(lowest, matrix[i][j-1])
					} else {
						lowest = minimumValue(lowest, matrix[i-1][j])
						lowest = minimumValue(lowest, matrix[i][j-1])
						lowest = minimumValue(lowest, matrix[i][j+1])
					}
				} else{
					if j==0{
						lowest = minimumValue(lowest, matrix[i][j+1])
						lowest = minimumValue(lowest, matrix[i-1][j])
						lowest = minimumValue(lowest, matrix[i+1][j])
					} else if j==len(matrix[i])-1{
						lowest = minimumValue(lowest, matrix[i][j-1])
						lowest = minimumValue(lowest, matrix[i-1][j])
						lowest = minimumValue(lowest, matrix[i+1][j])
					} else {
						lowest = minimumValue(lowest, matrix[i][j-1])
						lowest = minimumValue(lowest, matrix[i][j+1])
						lowest = minimumValue(lowest, matrix[i-1][j])
						lowest = minimumValue(lowest, matrix[i+1][j])
					}
				}
				if matrix[i][j] < lowest{
					sum += matrix[i][j]+1
					addNode(j, i)
				}
			}
		}
	}

	return sum
}

func findBasinSize(x int, y int, matrix [][]int, marked [][]int) int{
	size := 1
	if matrix[y][x] == 9{
		return 0
	} else{
		if x > 0 && matrix[y][x-1] > matrix[y][x] && marked[y][x-1] == 0{
			marked[y][x-1] = 1
			size += findBasinSize(x-1, y, matrix, marked)
		}
		if x < len(matrix[0])-1 && matrix[y][x+1] > matrix[y][x] && marked[y][x+1] == 0{
			marked[y][x+1] = 1
			size += findBasinSize(x+1, y, matrix, marked)
		}
		if y > 0 && matrix[y-1][x] > matrix[y][x] && marked[y-1][x] == 0{
			marked[y-1][x] = 1
			size += findBasinSize(x, y-1, matrix, marked)
		}
		if y < len(matrix)-1 && matrix[y+1][x] > matrix[y][x] && marked[y+1][x] == 0{
			marked[y+1][x] = 1
			size += findBasinSize(x, y+1, matrix, marked)
		}
	}

	return size
}

func minimumValue(val1 int, val2 int) int{
	if val1 <= val2{
		return val1
	} else {
		return val2
	}
}

func addNode(x int, y int){

	node := &coordinate{x: x, y: y}

	if head == nil{
		head = node
		return
	}

	aux := head
	for ; aux.next != nil; aux = aux.next{}

	aux.next = node
}

func readLines(f *os.File) [][]int{
	bf := bufio.NewScanner(f)

	var intMatrix = make([][]int, 0)
	for bf.Scan(){
		line := bf.Text()
		var intArr = make([]int, len(line))
		for i:=0;i<len(line);i++{
			intArr[i] = int(line[i])-48
		}
		intMatrix = append(intMatrix, intArr)
	}

	return intMatrix
}
