package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Action = iota
	Steps
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	depth, fwd, aim := 0, 0, 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		steps, _ := strconv.Atoi(tokens[Steps])

		switch tokens[Action] {
		case "forward":
			fwd += steps
      depth += aim * steps
		case "down":
			aim += steps
		case "up":
			aim -= steps
		default:
			log.Fatalf("Case not handled: %s \n", tokens[Action])
		}
	}

	fmt.Printf("At Depth [%d] and Forward [%d] \n", depth, fwd)
  fmt.Printf("Product is [%d]\n", depth * fwd)

}
