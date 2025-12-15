package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	position := 50
	zeroCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		dir := line[0]
		distStr := line[1:]
		distance, err := strconv.Atoi(distStr)

		if err != nil {
			fmt.Println("invalid rotation:", line)
			return
		}

		if dir == 'L' {
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		}else if dir == 'R' {
			position = (position + distance) % 100
		}else {
			fmt.Println("Invalid Direction:", line)
			return
		}

		if position == 0 {
			zeroCount++
		}
	}

	fmt.Println("The real code is:", zeroCount)
}
