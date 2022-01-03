package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Cave struct{
	caveType int // 0 for small cave 1 for big cave
	visited int // number of times a cave has been visited
	caveName string // name of the cave
	depth int	// depth of current path
	paths []Cave // paths to other caves adjacent to this cave to make traversal easier
	next *Cave // node for cave linked list for input reading at program start
}

var start *Cave = &Cave{caveType: 0, visited: 0, caveName: "start"}
var visitedSmallCave int = 0


func readLines() error{
	content, err := os.Open(os.Args[1])
	if err != nil{
		return err
	}

	bf := bufio.NewScanner(content)

	for bf.Scan(){
		path := strings.Split(bf.Text(), "-")
		leftCave := searchForNode(path[0])
		rightCave := searchForNode(path[1])
		if leftCave == nil{
			leftCave = addNode(path[0])
		}
		if rightCave == nil{
			rightCave = addNode(path[1])
		}

		leftCave.paths = append(leftCave.paths, *rightCave)
		rightCave.paths = append(rightCave.paths, *leftCave)
	}

	return nil
}

func searchForNode(name string) *Cave {
	for aux:=start; aux != nil; aux=aux.next{
		if aux.caveName == name{
			return aux
		}
	}
	return nil
}

func addNode(name string) *Cave{
	node := &Cave{}	
	node.caveType = 0
	if int(name[0]) < 97{
		node.caveType = 1
	}
	node.caveName = name
	node.visited = 0
	node.depth = 0
	node.paths = make([]Cave, 0)
	node.next = nil

	aux:=start

	for ; aux.next != nil; aux=aux.next{}

	aux.next = node

	return node
}

func recursiveStep(cave *Cave, flag int, str string) int{

	if cave.caveName == "end"{
		str += ", end"
		fmt.Println(str)
		return 1
	}

	internal_flag := 0

	if cave.caveType == 0{
		if cave.visited > 0{
			if flag == 0 && cave.caveName != "start"{
				internal_flag = 1
				flag = 1
			} else{
				return 0
			}
		}
	}

	numOfCompletedPaths := 0
	cave.visited++

	str += ", " + cave.caveName

	for i:=0; i<len(cave.paths); i++{
		caveSearch := searchForNode(cave.paths[i].caveName)
		numOfCompletedPaths += recursiveStep(caveSearch, flag, str)
	}

	cave.visited--

	if flag == 1 && internal_flag == 1{
		flag = 0
	}

	return numOfCompletedPaths
}

func main(){
	err := readLines()
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %v\n", err))
		return
	}

	for aux:=start; aux != nil; aux=aux.next{
		fmt.Printf("node name: %s\tneighbours: ", aux.caveName)
		for i:=0; i<len(aux.paths); i++{
			fmt.Printf("%s ", aux.paths[i].caveName)
		}
		fmt.Println()
	}

	numOfCompletedPaths := recursiveStep(start, 0, "")

	fmt.Println("Number of different paths =", numOfCompletedPaths)

}
