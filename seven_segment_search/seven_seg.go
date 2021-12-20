package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type SevenSegmentDisplay struct{
	fullDigit string
}

type KnownChars struct{
	a string
	b string
	c string
	d string
	e string
	f string
	g string
}

var (
	zero = SevenSegmentDisplay{}
	one = SevenSegmentDisplay{}
	two = SevenSegmentDisplay{}
	three = SevenSegmentDisplay{}
	four = SevenSegmentDisplay{}
	five = SevenSegmentDisplay{}
	six = SevenSegmentDisplay{}
	seven = SevenSegmentDisplay{}
	eight = SevenSegmentDisplay{}
	nine = SevenSegmentDisplay{}
	known = KnownChars{}
)

func main(){
	content, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println(fmt.Errorf("Opening file: %s", err).Error())
	}

	digits, output:= readLines(content)

	finalSum := 0
	for i:=0; i<len(digits); i++{
		getInitialDigits(digits[i])
		getChars(digits[i])
		results := getRemainingDigits()
		outputSplit := strings.Split(output[i], " ")
		
		var finalVal string
		for j:=0; j<len(outputSplit); j++{
			for k:=0; k<len(results); k++{
				if len(outputSplit[j]) == len(results[k]){
					if verifyNumber(outputSplit[j], results[k]){
						finalVal += strconv.Itoa(k)
					}
				}
			}
		}
		fmt.Printf("Value %d = %s\n", i+1, finalVal)
		val, _ := strconv.Atoi(finalVal)
		finalSum += val
	}

	fmt.Println("Sum of all values =", finalSum)
}

func verifyNumber(firstStr string, secondStr string) bool{
	for i:=0; i<len(firstStr); i++{
		for j:=0; j<len(secondStr); j++{
			if firstStr[i] == secondStr[j]{
				break
			}
			if j == len(secondStr)-1{
				return false
			}
		}
	}
	return true
}

func getRemainingDigits() []string {
	two.fullDigit = known.a + known.c + known.d + known.e + known.g
	three.fullDigit = known.a + known.c + known.d + known.f + known.g
	five.fullDigit = known.a + known.b + known.d + known.f + known.g
	six.fullDigit = known.a + known.b + known.d + known.e + known.f + known.g
	nine.fullDigit = known.a + known.b + known.c + known.d + known.f + known.g
	zero.fullDigit = known.a + known.b + known.c + known.e + known.f + known.g

	//build string array for easier handling of results (should have done this from the beginning...)
	var results = []string{zero.fullDigit, one.fullDigit, two.fullDigit, three.fullDigit, four.fullDigit, five.fullDigit, six.fullDigit, seven.fullDigit, eight.fullDigit, nine.fullDigit}

	return results
}

func getInitialDigits(digits string){
	digitSplit := strings.Split(digits, " ")
	for i:=0; i<len(digitSplit); i++{
		if len(digitSplit[i]) == 2{
			one.fullDigit = digitSplit[i]
		} else if len(digitSplit[i]) == 4{
			four.fullDigit = digitSplit[i]
		} else if len(digitSplit[i]) == 3{
			seven.fullDigit = digitSplit[i]
		} else if len(digitSplit[i]) == 7{
			eight.fullDigit = digitSplit[i]
		}
	}
}

func getChars(digits string) {
	digitSplit := strings.Split(digits, " ")
	var countChars = []int{0, 0, 0, 0, 0, 0, 0} 
	var chars = []string{"a", "b", "c", "d", "e", "f", "g"}

	// get char correspondent to the a segment first
	countCharsInString(seven.fullDigit, countChars)
	countCharsInString(eight.fullDigit, countChars)
	countCharsInString(one.fullDigit, countChars)

	var aChar int;
	for i:=0; i<len(countChars); i++{
		if countChars[i] == 2{
			aChar = i
			retrieveKnownChar(chars[i], "a")
		}
		countChars[i] = 0
	}

	// now we can get the remaining chars
	for i:=0; i<len(digitSplit); i++{
		countCharsInString(digitSplit[i], countChars)
	}

	for i:=0; i<len(countChars); i++{
		if countChars[i] == 4{
			retrieveKnownChar(chars[i], "e")
		} else if countChars[i] == 6{
			retrieveKnownChar(chars[i], "b")
		} else if countChars[i] == 8 && i != aChar{
			retrieveKnownChar(chars[i], "c")
		} else if countChars[i] == 9{
			retrieveKnownChar(chars[i], "f")
		} else if countChars[i] == 7{
			dFlag := -1
			for j:=0; j<len(four.fullDigit); j++{
				if string(four.fullDigit[j]) == chars[i]{
					dFlag = i
					break
				}
			}
			if dFlag == -1{
				retrieveKnownChar(chars[i], "g")
			} else if dFlag == i{
				retrieveKnownChar(chars[i], "d")
			}
		}
	}
}

func countCharsInString(str string, charCount []int){
	for i:=0; i<len(str); i++{
		if str[i] == 'a'{
			charCount[0]++
		} else if str[i] == 'b'{
			charCount[1]++
		} else if str[i] == 'c'{
			charCount[2]++
		} else if str[i] == 'd'{
			charCount[3]++
		} else if str[i] == 'e'{
			charCount[4]++
		} else if str[i] == 'f'{
			charCount[5]++
		} else if str[i] == 'g'{
			charCount[6]++
		}
	}
}

func retrieveKnownChar(element string, char string){
	if char == "a"{
		known.a = element
	} else if char == "b"{
		known.b = element
	} else if char == "c"{
		known.c = element
	} else if char == "d"{
		known.d = element
	} else if char == "e"{
		known.e = element
	} else if char == "f"{
		known.f = element
	} else if char == "g"{
		known.g = element
	}
}

func readLines(f *os.File)([]string, []string){
	bf := bufio.NewScanner(f)

	var digits = make([]string, 0)
	var output = make([]string, 0)

	for bf.Scan(){
		digitsAndOutput := strings.Split(bf.Text(), " | ")
		digits = append(digits, digitsAndOutput[0])
		output = append(output, digitsAndOutput[1])
	}

	return digits, output
}
