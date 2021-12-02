package main;

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
)

func main(){
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil{
		fmt.Println(err)
	}

	values := strings.Split(string(file), "\n")

	var conversion_arr = []int{}

	for idx, value := range values{
		if idx!=len(values)-1{
			intval, err := strconv.Atoi(value)
			if err != nil{
				panic(err)
			}
			conversion_arr = append(conversion_arr, intval)
		}
	}

	var sum1, sum2 int
	count := 0
	for idx, _ := range conversion_arr{
		if idx == len(conversion_arr)-3{
			break
		}
		sum1 = conversion_arr[idx]+conversion_arr[idx+1]+conversion_arr[idx+2]
		sum2 = conversion_arr[idx+1]+conversion_arr[idx+2]+conversion_arr[idx+3]
		if sum2 > sum1{
			fmt.Println(idx, ":", sum2, "(increased)")
			count++
		} else if sum2 == sum1{
			fmt.Println(idx, ":", sum2, "(no change)")
		} else{
			fmt.Println(idx, ":", sum2, "(decreased)")
		}
	}

	fmt.Println("Number of increases = ", count)

}
