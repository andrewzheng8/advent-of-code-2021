package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "strconv"
)

func main() {
  input := os.Stdin

  var err error
  scanner := bufio.NewScanner(input)

  var window [3]int

  for i := 0; i < 3; i++ {
    if !scanner.Scan() {
      log.Fatal("No input found")
    }
    window[i], err = strconv.Atoi(scanner.Text())
    check(err)
  }

  count := 0
  for scanner.Scan() {
    num, err := strconv.Atoi(scanner.Text())

    if err != nil {
      panic(err)
    }

    if num > window[0] {
      count++
    }
    window[0] = num
    
    if !scanner.Scan() {
      break
    }

    num, err = strconv.Atoi(scanner.Text())
    if err != nil {
      panic(err)
    }
 
    if num > window[1] {
      count++
    }
    window[1] = num

    if !scanner.Scan() {
      break
    }
    
    num, err = strconv.Atoi(scanner.Text())
    if err != nil {
      panic(err)
    }
 
    if num > window[2] {
      count++
    }
    window[2] = num
 
  }

  fmt.Printf("Num increases: %d \n", count)

}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
