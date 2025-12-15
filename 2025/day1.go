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

	//part II:

	if err := method_0x434C49434B(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func method_0x434C49434B() error {

	scanner := bufio.NewScanner(os.Stdin)
	
	position := 50
	totalZeros := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0{
			continue
		}
		
		dir := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return fmt.Errorf("invalid", line)
		}

		_, hits := applyRotation(position, dir, distance)
		totalZeros += hits

	}
	
	if err := scanner.Err(); err != nil{
		return err
	}

	fmt.Println("The actual password is:", totalZeros)
	return nil
}


func applyRotation(position int, dir byte, distance int) (int, int) {
	zeroCount := 0

	for i := 0; i < distance; i++ {
		switch dir {
		case 'L':
			position--
			if position < 0 {
				position = 99
			}
		case 'R':
			position++
			if position > 99 {
				position = 0
			}
		default:
			panic("invalid direction")
		}

		if position == 0 {
			zeroCount++
		}
	}

	return position, zeroCount
}






