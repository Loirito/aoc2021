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
	
	numOfFishForEachDaySli := readLineIntoIntSlice(content)

	numOfDays, err := strconv.Atoi(os.Args[2])
	if err != nil{
		fmt.Println(fmt.Errorf("Argument 2 must be an integer for the number of days: %s", err).Error())
		return
	}

	numOfFishForEachDaySli = generateLanternfish(numOfFishForEachDaySli, numOfDays)

	countFish := 0
	for i:=0; i<len(numOfFishForEachDaySli); i++{
		countFish += numOfFishForEachDaySli[i]
	}
	fmt.Println("Number of fish=", countFish)

}

func readLineIntoIntSlice(f *os.File) []int {
	bf := bufio.NewScanner(f)

	bf.Scan()
	strSli := strings.Split(bf.Text(), ",")
	var intSli = make([]int, len(strSli))
	for i:=0; i<len(strSli); i++{
		intSli[i], _ = strconv.Atoi(strSli[i])
	}

	var newSlice = make([]int, 9)
	for i:=0; i<len(intSli); i++{
		newSlice[intSli[i]]++
	}
	
	return newSlice
}

func generateLanternfish(slice []int, numOfDays int) []int{
	for i:=1; i<=numOfDays; i++{
		numOfNewFish := 0
		for i:=0; i<len(slice); i++{
			if i==0{
				numOfNewFish = slice[i]
			} else{
				slice[i-1] = slice[i]
				if i == 7{
					slice[i-1] += numOfNewFish
				}
				if i == 8{
					slice[i] = numOfNewFish
				}
			}
		}

	}
	return slice
}
