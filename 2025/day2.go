package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(n int) bool {
	s := strconv.Itoa(n) //Interger to ASCII function takes the input and makes it a string
	length := len(s)

	if length % 2 != 0 {
		return false 
	}
	

	mid := length/2 

	return s[:mid] == s[mid:]
}

func run(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.TrimSpace(input)

	ranges := strings.Split(input, ",")

	total := 0

	for _, ranges := range ranges{
		parts := strings.Split(ranges, "-")

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		
		for num := start; num <= end; num++{
			if isValid(num){
				total += num
			}
		}
	}
	return total

}


func main(){
  
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run day2.go <input_file.txt>")
		return
	} 

	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	solution := run(string(data))

	fmt.Println("The sum of Invalid ID's are:", solution)

}
