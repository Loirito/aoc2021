package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %s", err).Error())
		return
	}
	defer content.Close()

	coord_slice, max_x, max_y := ReadLinesIntoSlice(content)
	if coord_slice == nil{
		return
	}
	var table = [][]int{}
	for i:=0; i<max_x+1; i++{
		table_length := make([]int, max_y+1)
		table = append(table, table_length)
	}

	fmt.Println("Number of overlaps =", findNumberOfOverlaps(coord_slice, table))

	/*for i:=0; i<len(table); i++{
		for j:=0; j<len(table[0]); j++{
			if table[j][i]==0{
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", table[j][i])
			}
		}
		fmt.Println()
	}*/
}

func findNumberOfOverlaps(coordinates [][]int, table [][]int) int{
	count_occ := 0
	for i:=0; i<len(coordinates); i++{
		x1 := coordinates[i][0]
		y1 := coordinates[i][1]
		x2 := coordinates[i][2]
		y2 := coordinates[i][3]

		if x1 == x2{
			if y1 > y2{
				for j:=y1; j>=y2;j--{
					table[x1][j]++
					if table[x1][j] == 2{
						count_occ++
					}
				}
			} else{
				for j:=y1; j<=y2; j++{
					table[x1][j]++
					if table[x1][j] == 2{
						count_occ++
					}
				}
			}
		} else if y1 == y2{
			if x1 > x2{
				for j:=x1; j>=x2; j--{
					table[j][y1]++
					if table[j][y1] == 2{
						count_occ++
					}
				}
			} else{
				for j:=x1; j<=x2; j++{
					table[j][y1]++
					if table[j][y1] == 2{
						count_occ++
					}
				}
			}
		} else if x1 < x2{
			it := x2-x1
			if y1 < y2{
				for j:=0; j<=it; j++{
					table[x1+j][y1+j]++
					if table[x1+j][y1+j] == 2{
						count_occ++
					}
				}
			} else {
				for j:=0; j<=it; j++{
					table[x1+j][y1-j]++
					if table[x1+j][y1-j] == 2{
						count_occ++
					}
				}
			}
		} else{
			it := x1-x2
			if y1 < y2{
				for j:=0; j<=it; j++{
					table[x1-j][y1+j]++
					if table[x1-j][y1+j]==2{
						count_occ++
					}
				}
			} else {
				for j:=0; j<=it; j++{
					table[x1-j][y1-j]++
					if table[x1-j][y1-j]==2{
						count_occ++
					}
				}
			}
		}
	}
	return count_occ
}

func ReadLinesIntoSlice(f *os.File) ([][]int, int, int){
	bf := bufio.NewScanner(f)
	var coord_slice = [][]int{}
	max_x := 0
	max_y := 0
	for bf.Scan(){
		first_split := strings.Split(bf.Text(), " -> ")
		one_end := strings.Split(first_split[0], ",")
		second_end := strings.Split(first_split[1], ",")
		x, y := convertToInt(one_end)
		x2, y2 := convertToInt(second_end)
		if x == -1 && y == -1 || x2 == -1 && y2 == -1{
			return nil, -1, -1
		}
		if x > max_x { max_x = x}
		if y > max_y { max_y = y}
		if x2 > max_x { max_x = x2}
		if y2 > max_y { max_y = y2}
		var coord_int = []int{x, y, x2, y2}
		coord_slice = append(coord_slice, coord_int)
	}
	
	return coord_slice, max_x, max_y
}

func convertToInt(sli []string) (int, int){
	x, err := strconv.Atoi(sli[0])
	if err != nil{
		fmt.Println(fmt.Errorf("Converting first elem. of string to int: %s", err).Error())
		return -1, -1
	}
	y, err := strconv.Atoi(sli[1])
	if err != nil{
		fmt.Println(fmt.Errorf("Converting second elem. of string to int: %s", err).Error())
		return -1, -1
	}
	return x, y
}
