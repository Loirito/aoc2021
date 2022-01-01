package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

type Cell struct{
	x int
	y int
}

type Octopus struct{
	flashed bool
	value int
}

var octopusMap = make(map[Cell]Octopus)

func readLines() error {
	content, err := os.Open(os.Args[1])
	if err != nil{
		return err
	}

	bf := bufio.NewScanner(content)

	x:=0
	for bf.Scan(){
		line := bf.Text()
		for j:=0; j<len(line); j++{
			octopusMap[Cell{x,j}] = Octopus{false, int(line[j])-48}
		}
		x++
	}

	return nil
}

func flashStep() int{
	/*for point, octo := range octopusMap{
		octo.value++
		octopusMap[point] = octo
	}*/

	flashes := 0
	for point, _ := range octopusMap{
		flashes += incrementValue(point, 9)
	}

	for point, octo := range octopusMap{
		octo.flashed = false
		octopusMap[point] = octo
	}

	fmt.Println("Num of flashes this step=", flashes)

	return flashes
	
}

func incrementValue(point Cell, maxLen int) int{
	octopus := octopusMap[point]
	if octopus.flashed == true || point.x < 0 || point.x > maxLen || point.y < 0 || point.y > maxLen{
		return 0
	}
	octopus.value++
	octopusMap[point] = octopus
	if octopus.value > 9{
		octopus.flashed = true
		octopus.value = 0
		octopusMap[point] = octopus
		flashes := 1
		top := Cell{point.x, point.y+1}
		bot := Cell{point.x, point.y-1}
		left := Cell{point.x-1, point.y}
		right := Cell{point.x+1, point.y}
		leftTop := Cell{point.x-1, point.y+1}
		leftBot := Cell{point.x-1, point.y-1}
		rightTop := Cell{point.x+1, point.y+1}
		rightBot := Cell{point.x+1, point.y-1}
		flashes += incrementValue(top, maxLen)
		flashes += incrementValue(bot, maxLen)
		flashes += incrementValue(left, maxLen)
		flashes += incrementValue(right, maxLen)
		flashes += incrementValue(leftTop, maxLen)
		flashes += incrementValue(leftBot, maxLen)
		flashes += incrementValue(rightTop, maxLen)
		flashes += incrementValue(rightBot, maxLen)
		return flashes
	}
	return 0
}

func checkIfAllFlashed() int{
	flag := 0
	for _, octo := range octopusMap{
		if octo.value != 0{
			flag = 1
		}
	}
	return flag
}


func printMap(sortedSlice []Cell){
	for i:=0; i<len(sortedSlice); i++{
		if sortedSlice[i].y != 9{
			fmt.Printf("%d", octopusMap[sortedSlice[i]].value)
		} else {
			fmt.Printf("%d\n", octopusMap[sortedSlice[i]].value)
		}
	}
}

func main(){
	err := readLines()
	if err != nil{
		fmt.Println(fmt.Errorf("Couldn't open this file: %v", err))
		return
	}

	var sortedCells []Cell
	for point, _ := range octopusMap{
		sortedCells = append(sortedCells, Cell{point.x, point.y})
	}

	sort.Slice(sortedCells, func(i, j int) bool {
		if sortedCells[i].x != sortedCells[j].x{
			return sortedCells[i].x < sortedCells[j].x
		} else {
			return sortedCells[i].y < sortedCells[j].y
		}
	})


	fmt.Println("Before any steps:")
	printMap(sortedCells)

	flashes := 0
	for i:=0; i<1000;i++{
		flashes += flashStep()
		flag := checkIfAllFlashed()
		fmt.Println("\nAfter", i+1, " steps")
		printMap(sortedCells)
		if flag == 0{
			fmt.Printf("All the octopuses flashed after step %d!!\n", i+1)
			return
		}
	}
}
