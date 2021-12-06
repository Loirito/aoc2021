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
	var oxy_gen_arr = []string{first_line}
	var co2_arr = []string{first_line}

	countBits(first_line, zeros, ones)

	for bf.Scan(){
		line := bf.Text()
		oxy_gen_arr = append(oxy_gen_arr, line)
		co2_arr = append(co2_arr, line)
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

	for i:=0; i<len(first_line); i++{
		oxy_ones, oxy_zeros := countRelevantBit(oxy_gen_arr, i)
		co2_ones, co2_zeros := countRelevantBit(co2_arr, i)
		if oxy_ones > oxy_zeros {
			oxy_gen_arr = deleteUnwantedValues(oxy_gen_arr, i, '1')
		} else if oxy_zeros > oxy_ones{
			oxy_gen_arr = deleteUnwantedValues(oxy_gen_arr, i, '0')
		} else{
			oxy_gen_arr = deleteUnwantedValues(oxy_gen_arr, i, '1')
		}
		if co2_ones < co2_zeros{
			co2_arr = deleteUnwantedValues(co2_arr, i, '1')
		} else if co2_zeros < co2_ones{
			co2_arr = deleteUnwantedValues(co2_arr, i, '0')
		} else{
			co2_arr = deleteUnwantedValues(co2_arr, i, '0')
		}
	}

	fmt.Println("gamma rate binary=", gamma_rate_bin)
	fmt.Println("epsilon rate binary=", epsilon_rate_bin)
	gamma_rate_dec := convertBinToDec(gamma_rate_bin)
	epsilon_rate_dec := convertBinToDec(epsilon_rate_bin)
	fmt.Println("gamma rate decimal =", gamma_rate_dec)
	fmt.Println("epsilon rate decimal =", epsilon_rate_dec)
	fmt.Println("power consumption =", gamma_rate_dec*epsilon_rate_dec)
	
	// part 2 
	fmt.Println("\noxygen gen. binary=", oxy_gen_arr[0])
	fmt.Println("co2 scrubber binary=", co2_arr[0])
	oxygen_gen_dec := convertBinToDec(oxy_gen_arr[0])
	co2_scrubber_dec := convertBinToDec(co2_arr[0])
	fmt.Println("oxygen gen. decimal =", oxygen_gen_dec)
	fmt.Println("co2 scrubber decimal =", co2_scrubber_dec)
	fmt.Printf("life support rating = %.0f\n", oxygen_gen_dec*co2_scrubber_dec)
	
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

// only used for part 2
func countRelevantBit(arr []string, index int) (int, int){
	ones := 0
	zeros := 0
	for i:=0; i<len(arr); i++{
		if int32(arr[i][index]) == '1'{
			ones++
		} else{
			zeros++
		}
	}
	return ones, zeros
}

// value corresponds to the bit value we want to preserve
func deleteUnwantedValues(arr []string, index int, value int32) []string{
	i:=0
	if len(arr) == 1{
		return arr
	}
	for i<len(arr){
		array_value := arr[i]
		if int32(array_value[index]) == value{
			i++
		} else {
			arr[i] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		}
	}
	return arr
}
