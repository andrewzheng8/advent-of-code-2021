package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	CO2 = iota
	O2
)

func streamToArray() ([2][]int, int) {
	nums := [2][]int{make([]int, 0), make([]int, 0)}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	buf := scanner.Bytes()
	n := len(buf)

	marker := 1 << (n - 1)

	num := 0
	for i := 0; i < n; i++ {
		num |= int(buf[n-1-i]-'0') << i
	}
	//if num greater than or equal to marker than
	//goes to 0 +1 = 1
	//else its neg meaning num starts with 0
	//goes to -1 + 1 = 0
	grp := (num-marker)>>63 + 1

	nums[grp] = append(nums[grp], num)

	for scanner.Scan() {
		buf = scanner.Bytes()
		num = 0
		for i := 0; i < n; i++ {
			num |= int(buf[n-1-i]-'0') << i
		}
		grp = (num-marker)>>63 + 1
		nums[grp] = append(nums[grp], num)
	}

	return nums, n

}

func getRating(n, gas, lo, hi int, nums []int) int {
  //if only one left in range or we filtered all bits, take first in range 
	if hi-lo == 1 || n == -1 {
		return nums[lo]
	}
	marker := 1 << n

	left, right := lo, hi-1

	//partition left and right by 0 and 1 in the nth bit
	for left < right {

		//get the symmetric diff between the two
		//can use to swap one for the other
		//symm diff = a xor b
		//a xor symm diff -> zeros out only in a, add only in b, leaves intersection
		//means it becomes b
		//same applies to b
		diff := nums[left] ^ nums[right]

		//left should have 1 at nth bit
		//right should have 0
		//then they can be swapped
		leftGood := (0 - (nums[left] & marker)) >> 63
		rightGood := ((nums[right] & marker) - 1) >> 63
		diff &= leftGood & rightGood

		//if diff is 0 no change
		//else diff will swap the two
		nums[left] ^= diff
		nums[right] ^= diff

		//inc left cursor if left points to 0 at nth bit
		//dec right cursor if right points to 1 at nth bit
		left += 1 - ((nums[left] & marker) >> n)
		right -= (nums[right] & marker) >> n
	}

	left += 1 - ((nums[left] & marker) >> n)
	//left points to the end of the left partition with 0s

	//cmp used to tell us to pick the left or right side
	//gas o2 = 1
	//gas co2 = 0
	cmp := (gas - 1) ^ ((hi + lo - left<<1) >> 63)

	rmask := -1 - cmp
	lmask := 0 + cmp

	return getRating(n-1, gas, (lo&lmask)|(left&rmask), (left&lmask)|(hi&rmask), nums)

}

func main() {
	nums, bits := streamToArray()

	most := (len(nums[1])-len(nums[0]))>>63 + 1
	least := 0 - ((len(nums[1]) - len(nums[0])) >> 63)

	o2Rtg := getRating(bits-2, O2, 0, len(nums[most]), nums[most])
	co2Rtg := getRating(bits-2, CO2, 0, len(nums[least]), nums[least])

	fmt.Printf("O2 Rating = %d || CO2 Rating = %d || product = %d\n", o2Rtg, co2Rtg, o2Rtg*co2Rtg)

}
