package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(err)
		return
	}
	defer content.Close()
	
	bf := bufio.NewScanner(content)

	var cmd string
	var val, horizontal, depth int

	for bf.Scan(){
		split := strings.Split(bf.Text(), " ")
		cmd = split[0]
		val, _ = strconv.Atoi(split[1])
		if cmd == "forward"{
			horizontal += val
		} else if cmd == "down"{
			depth += val
		} else if cmd == "up"{
			depth -= val
		}
	}

	fmt.Println("multiplied value =", horizontal*depth)
}
