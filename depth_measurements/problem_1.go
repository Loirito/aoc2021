package main;

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	file, err := os.Open("test.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()

	stdin := bufio.NewScanner(file)
	
	value := 3000
	count := 0
	var compare int
	for stdin.Scan(){
		compare, _ = strconv.Atoi(stdin.Text())
		if compare > value {
			fmt.Println(compare, "(increased)")
			count++
		} else{
			fmt.Println(compare, "(decreased)")
		}
		value = compare 
	}

	fmt.Println("number of increases = ", count)

	if err := stdin.Err(); err != nil {
		fmt.Println(err)
	}
}
