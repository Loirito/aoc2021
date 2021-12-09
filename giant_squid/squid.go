package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
	"strings"
)

func main(){
	drawn_values , main_matrix, err := readLines()
	if err != nil{
		return
	}

	var won_boards = make([]int, len(main_matrix))
	for i:=0; i<len(drawn_values); i++{
		mark := drawn_values[i]
		for idx:=0; idx<len(main_matrix); idx++{
			var ret_val int
			if won_boards[idx] != -1{
				main_matrix[idx], ret_val = markValuesInMatrix(main_matrix[idx], mark)
			}
			if ret_val != -1{
				won_boards[idx] = -1
				all_boards_done := 0
				for j:=0; j<len(won_boards); j++{
					if won_boards[j] != -1{
						break
					} else{
						all_boards_done++
					}
				}
				if all_boards_done == len(won_boards){
					sum := getResult(main_matrix[idx])
					fmt.Println("Score =", sum*ret_val)
					return
				}
			}
		}
	}
}

func getResult(matrix [][]int) int{
	sum := 0
	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[0]); j++{
			if matrix[i][j] != -1{
				sum += matrix[i][j]
			}
		}
	}
	return sum
}

func markValuesInMatrix(matrix [][]int, mark int)([][]int, int){
	for i:=0; i<len(matrix); i++{
		line := 0
		for j:=0; j<len(matrix[i]); j++{
			if matrix[i][j] == -1{
				line++
			}
			if matrix[i][j] == mark{
				matrix[i][j] = -1
				line++
			}
			if line == len(matrix[i]){
				return matrix, mark
			}
		}
	}
	for j:=0; j<len(matrix[0]); j++{
		col := 0
		for i:=0; i<len(matrix); i++{
			if matrix[i][j] == -1{
				col++
			}
			if col == len(matrix){
				return matrix, mark
			}
		}
	}
	return matrix, -1
}

func readLines() ([]int, [][][]int, error){
	content, err := os.Open(os.Args[1])
	if err != nil{
		err_msg := fmt.Errorf("Opening file: %s", err)
		fmt.Println(err_msg.Error())
		return nil, nil, err
	}

	bf := bufio.NewScanner(content)
	bf.Scan()
	drawn_values := strings.Split(bf.Text(), ",")
	var drawn_values_int = []int{}
	for i:=0; i<len(drawn_values); i++{
		if drawn_values[i] != ""{
			val, err := strconv.Atoi(drawn_values[i])
			if err != nil{
				fmt.Println(fmt.Errorf("Converting marked value to int - %s", err).Error())
			}
			drawn_values_int = append(drawn_values_int, val)
		}
	}
	main_idx := -1
	var main_matrix = [][][]int{}
	var aux_matrix = [][]int{}
	for bf.Scan(){
		if bf.Text() == ""{
			if main_idx != -1{
				main_matrix = append(main_matrix, aux_matrix)
			}
			main_idx++
			aux_matrix = nil
		} else{
			aux_matrix2 := strings.Split(bf.Text(), " ")
			var aux_matrix2_int = []int{}
			for i:=0; i<len(aux_matrix2); i++{
				if aux_matrix2[i] != ""{
					val, err := strconv.Atoi(aux_matrix2[i])
					if err != nil{
						fmt.Println(fmt.Errorf("Converting marked value to int - %s", err).Error())
					}
					aux_matrix2_int = append(aux_matrix2_int, val) 
				}
			}
			aux_matrix = append(aux_matrix, aux_matrix2_int)
		}
	}
	main_matrix = append(main_matrix, aux_matrix)
	return drawn_values_int, main_matrix, nil
}
