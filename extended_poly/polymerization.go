package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

var rulesMap = make(map[string]string)
var countOccMap = make(map[string]uint64)
var currentRules = make(map[string]uint64)


func readLines() (string, error){
	content, err := os.Open(os.Args[1])
	if err != nil{
		return "", err
	}

	defer content.Close()

	bf := bufio.NewScanner(content)

	bf.Scan()
	polymer := bf.Text()

	bf.Scan() // go over empty line

	for bf.Scan(){
		pair := strings.Split(bf.Text(), " -> ")
		rulesMap[pair[0]] = pair[1]
		currentRules[pair[0]] = 0
	}

	return polymer, nil
}

func evolvePolymer(currRules map[string]uint64) map[string]uint64{
	var rules = make(map[string]uint64)
	for key, value := range currRules{
		if value != 0{
			if val, ok := rulesMap[key]; ok{
				if _, ok2 := countOccMap[val]; !ok2{
					countOccMap[val] = value
				} else{
					countOccMap[val] += value
				}
				str1 := string(key[0])+val
				str2 := val+string(key[1])
				rules[str1] += value
				rules[str2] += value
			}
		}
	}

	return rules
}

func countInitialOccs(poly string){
	for i:=0; i<len(poly); i++{
		if i != len(poly)-1{
			str := string(poly[i])+string(poly[i+1])
			currentRules[str] += 1
		}
		if val, ok := countOccMap[string(poly[i])]; ok{
			countOccMap[string(poly[i])] = val+1
		} else{
			countOccMap[string(poly[i])] = 1
		}
	}
}

func main(){
	polymer, err := readLines()
	if err != nil{
		fmt.Println(fmt.Errorf("[error] %v\n", err))
		return
	}

	fmt.Println("Template:\t", polymer)

	countInitialOccs(polymer)

	fmt.Println("initial state")
	for key, val := range countOccMap{
		fmt.Printf("%s --> %d\n", key, val)
	}

	fmt.Println("rules initial")
	for key, val := range currentRules{
		fmt.Printf("%s --> %d\n", key, val)
	}

	for i:=1; i<=40; i++{
		currentRules = evolvePolymer(currentRules)
		fmt.Printf("-----step %d-----\n", i)
		for key, val := range countOccMap{
			fmt.Printf("%s --> %d\n", key, val)
		}
	}

	max := uint64(0)
	min := uint64(999999999999999999)
	var maxChar, minChar string

	fmt.Println("\n\nFinal Values:")
	fmt.Println()

	for elem, num := range countOccMap{
		fmt.Println(elem, "-->", num)
		if num < min{
			min = num
			minChar = elem
		}
		if num > max{
			max = num
			maxChar = elem
		}
	}

	fmt.Printf("Element with highest num. of occurences is %s with %d appearances in the polymer\n", maxChar, max)

	fmt.Printf("Element with lowest num. of occurences is %s with %d appearances in the polymer\n", minChar, min)

	fmt.Println("Difference =", max-min)
}
