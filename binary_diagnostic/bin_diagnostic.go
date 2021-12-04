package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
)

func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(err)
		return
	}
	defer content.Close()
	
	bf := bufio.NewScanner(content)
	bf.Scan();
	first_line := bf.Text()
	var ones = make([]int, len(first_line))
	var zeros = make([]int, len(first_line))

	countBits(first_line, zeros, ones)

	for bf.Scan(){
		line := bf.Text()
		countBits(line, zeros, ones)
	}
	
	var gamma_rate_bin, epsilon_rate_bin string
	for i:=0; i<len(ones); i++{
		if ones[i] > zeros[i]{
			gamma_rate_bin = gamma_rate_bin + "1"
			epsilon_rate_bin = epsilon_rate_bin + "0"
		} else if zeros[i] > ones[i]{
			gamma_rate_bin = gamma_rate_bin + "0"
			epsilon_rate_bin = epsilon_rate_bin + "1"
		}
	}

	fmt.Println("gamma rate binary=", gamma_rate_bin)
	fmt.Println("epsilon rate binary=", epsilon_rate_bin)
	gamma_rate_dec := convertBinToDec(gamma_rate_bin)
	epsilon_rate_dec := convertBinToDec(epsilon_rate_bin)
	fmt.Println("gamma rate decimal =", gamma_rate_dec)
	fmt.Println("epsilon rate decimal =", epsilon_rate_dec)
	fmt.Println("power consumption =", gamma_rate_dec*epsilon_rate_dec)
}

func countBits(str string, zero_arr []int, ones_arr []int){
	for i:=0; i<len(str); i++{
		if str[i] == '1'{
			ones_arr[i] += 1
		} else if str[i] == '0'{
			zero_arr[i] += 1
		}
	}
}

func convertBinToDec(binary string) float64{
	var exp float64
	var val float64
	for i:=len(binary)-1; i>=0; i--{
		if binary[i] == '1'{
			val += math.Pow(2, exp)
		}
		exp += 1
	}
	return val
}
