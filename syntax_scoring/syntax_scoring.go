package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"sort"
)

func readLines(path string) ([]string, error) {
	handler, err := os.Open(path)
	if err != nil{
		return nil, err
	}

	bf := bufio.NewScanner(handler)

	lines := make([]string, 0)

	for bf.Scan(){
		lines = append(lines, bf.Text())
	}

	return lines, nil
}

func parseUserInput(str string)(string, error){
	if str == "1\n" || str == "2\n"{
		return str, nil
	}
	return "", errors.New("You can only select 1 (testcase.txt) or 2 (input.txt)")
}

func checkInvalidLines(str string, line int) (int, []int){
	stack := make([]int, 0)		// this slice will act as a stack where representations are:
					// 0-par, 1-square bracks, 2-bracks, 3-gt/lt

	possibleVals := []string{")", "]", "}", ">"}
	for i:=0; i<len(str); i++{
		if str[i] == '('{
			stack = append(stack, 0)
		} else if str[i] == '['{
			stack = append(stack, 1)
		} else if str[i] == '{'{
			stack = append(stack, 2)
		} else if str[i] == '<'{
			stack = append(stack, 3)
		} else if str[i] == ')'{
			if stack[len(stack)-1] != 0{
				fmt.Printf("[%d] Expected %s but found ) instead\n", line, possibleVals[stack[len(stack)-1]])
				return 3, nil
			}
			stack = stack[:len(stack)-1]
		} else if str[i] == ']'{
			if stack[len(stack)-1] != 1{
				fmt.Printf("[%d] Expected %s but found ] instead\n", line, possibleVals[stack[len(stack)-1]])
				return 57, nil
			}
			stack = stack[:len(stack)-1]
		} else if str[i] == '}'{
			if stack[len(stack)-1] != 2{
				fmt.Printf("[%d] Expected %s but found } instead\n", line, possibleVals[stack[len(stack)-1]])
				return 1197, nil
			}
			stack = stack[:len(stack)-1]
		} else if str[i] == '>'{
			if stack[len(stack)-1] != 3{
				fmt.Printf("[%d] Expected %s but found > instead\n", line, possibleVals[stack[len(stack)-1]])
				return 25137, nil
			}
			stack = stack[:len(stack)-1]
		}
	}

	return 0, stack
}					

func scoreRemainingLines(stack []int) int{
	score := 0
	for i:=len(stack)-1; i>=0; i--{
		score *= 5
		if stack[i] == 0{
			score += 1
		} else if stack[i] == 1{
			score += 2
		} else if stack[i] == 2{
			score += 3
		} else if stack[i] == 3{
			score += 4
		}
	}
	
	return score
}


func main(){
	test_path := "./testcase.txt"
	input_path := "./input.txt"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("1. Run for testcase.txt\n2. Run for input.txt")
	input, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println(fmt.Errorf("Couldn't read line: %v", err))
	}
	option, err := parseUserInput(input)
	if err != nil{
		fmt.Println(fmt.Errorf("Expected 1 or 2: %v", err))
	}
	
	var path string
	if option == "1\n"{
		path = test_path
	} else {
		path = input_path
	}

	lines, err := readLines(path)
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %v", err))
	}

	totalSyntaxErrorScore := 0
	autocompleteScore := make([]int, 0)
	for i:=0; i<len(lines); i++{
		val, stack := checkInvalidLines(lines[i], i)
		totalSyntaxErrorScore += val
		if stack != nil{
			autocompleteScore = append(autocompleteScore, scoreRemainingLines(stack))
			fmt.Println("[", i,"] Autocomplete score =", autocompleteScore[len(autocompleteScore)-1])
		}
	}

	sort.Ints(autocompleteScore)

	fmt.Println("Syntax error score =", totalSyntaxErrorScore)
	fmt.Println("Autocomplete middle score =", autocompleteScore[len(autocompleteScore)/2])
}
