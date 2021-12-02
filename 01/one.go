package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  input := os.Stdin
  
  scanner := bufio.NewScanner(input)

  var last int
  var err error

  if scanner.Scan() {
    last, err = strconv.Atoi(scanner.Text())
    if err != nil {
      panic(err)
    }
  } else {
    fmt.Println("No input found")
    os.Exit(0)
  }

  count := 0
  for scanner.Scan() {
    num, err := strconv.Atoi(scanner.Text())

    if err != nil {
      panic(err)
    }

    if num > last {
      count++
    }
    last = num

  }

  fmt.Printf("Num increases: %d \n", count)

}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
