package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file - %s", err).Error())
		return
	}
	defer content.Close()
	
	inputSlice := readLineIntoIntSlice(content)

	numOfDays, err := strconv.Atoi(os.Args[2])
	if err != nil{
		fmt.Println(fmt.Errorf("Argument 2 must be an integer for the number of days: %s", err).Error())
		return
	}
	fmt.Println("Number of fish =", generateLanternfish(inputSlice, numOfDays))
}

func readLineIntoIntSlice(f *os.File) []int {
	bf := bufio.NewScanner(f)

	bf.Scan()
	strSli := strings.Split(bf.Text(), ",")
	var intSli = make([]int, len(strSli))
	for i:=0; i<len(strSli); i++{
		intSli[i], _ = strconv.Atoi(strSli[i])
	}

	return intSli
}

func generateLanternfish(slice []int, numOfDays int) int{
	var countFish int
	for i:=0; i<=numOfDays; i++{
		fmt.Println("Day", i)
		if i != 0{
			slice = decreaseTimer(slice)
		}
	}
	for i:=0; i<len(slice); i++{
		countFish++
	}
	return countFish
}

func decreaseTimer(slice []int) []int{
	for i:=0; i<len(slice); i++{
		slice[i]--
		if slice[i] == -1{
			slice[i] = 6
			slice = append(slice, 9)
		}
	}
	return slice
}
