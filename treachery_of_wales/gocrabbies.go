package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)


func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %s", err).Error())
		return
	}

	intSlice := readLineIntoIntSlice(content)

	avg := calculateAverage(intSlice)

	loop := false
	fuel := avg
	for !loop{
		loop, fuel = findLowestPossibleFuel(fuel, intSlice)
	}

	fmt.Println("Lowest value for fuel=", fuel)
}

func findLowestPossibleFuel(avg int, slice []int) (bool, int){
	fuel_avg:=0
	fuel_right:=0
	fuel_left:=0
	left := avg-1
	right := avg+1
	for i:=0; i<len(slice); i++{
		fuel_avg += calculateFuel(slice[i], avg)
		fuel_right += calculateFuel(slice[i], right)
		fuel_left += calculateFuel(slice[i], left)
	}
	if fuel_avg < fuel_right && fuel_avg < fuel_left{
		return true, fuel_avg
	} else if fuel_right < fuel_avg && fuel_right < fuel_left{
		return false, right
	} else{
		return false, left
	} 
}

func calculateFuel(slice_val int, comp int)int{
	if slice_val>comp{
		sum := 0
		for i:=1; i<=slice_val-comp; i++{
			sum += i
		}
		return sum
	} else{
		sum := 0
		for i:=1; i<=comp-slice_val; i++{
			sum += i
		}
		return sum
	}
}

func calculateAverage(slice []int) int{
	sum := 0
	for i:=0; i<len(slice); i++{
		sum += slice[i]
	}
	avg := sum/len(slice)

	return avg
}

func readLineIntoIntSlice(f *os.File) []int{
	bf := bufio.NewScanner(f)
	bf.Scan()
	strSlice := strings.Split(bf.Text(), ",")
	intSlice := make([]int, len(strSlice))
	for i:=0; i<len(strSlice); i++{
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}
	
	return intSlice
}
