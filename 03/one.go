package main

import (
	"bufio"
	"fmt"
	//"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	buf := scanner.Bytes()
	n := len(buf)
	counter := make([]int, len(buf))

	//use bit manipulation to transform 1 into 0, 0 int -1
	//shift left - 1
	//1*2 - 1 = 1
	//0*2 - 1 = -1
	//so we will know which is more common by if the counter is pos/neg
	//the sum should not be 0 since we should be able to decide if one
	//is more common than the other
	for i := 0; i < n; i++ {
		counter[i] += int(buf[i]-'0')<<1 - 1
	}

	for scanner.Scan() {
		buf = scanner.Bytes()
		for i := 0; i < n; i++ {
			counter[i] += int(buf[i]-'0')<<1 - 1
		}
	}

	//use the fact that 0 common is a neg counter
	//1 common is a pos counter
	// x >= 0; full arith right shift will clear
	// x < 0; full arith right shift results in -1
	// + 1 gives us the common bit
	// least is opposite -> xor 1
	// use shift and or to give us gamma and epsilon
	gamma, epsilon := 0, 0
	for i := 0; i < n; i++ {
		most := counter[n-1-i]>>63 + 1
		least := most ^ 1
		gamma |= most << i
		epsilon |= least << i
	}

	fmt.Printf("gamma = %d || epsilon = %d || product = %d\n", gamma, epsilon, gamma*epsilon)

}
