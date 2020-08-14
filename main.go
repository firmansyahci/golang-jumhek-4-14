package main

import (
    "log"
    "os"
    "bufio"
	"strings"
	"strconv"
	"errors"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	dataInput := strings.Split(scanner.Text(), " ")

	items, err := createItems(dataInput)
	if (err != nil) {
		log.Println(err)
		os.Exit(1)
	}

	max := sum(items) / 2;
	sortItems := sortDesc(items)

	cart1, cart2, diff := inputToCarts(sortItems, max)
	log.Printf(`cart 1: %v, cart2: %v, diff: %v`, cart1, cart2, diff)
}

func inputToCarts(items []int, max int) ([]int, []int, int) {
	cart1 := []int{}
	cart2 := []int{}
	sumCart1 := 0
	sumCart2 := 0

	for i, _ := range items {
		if (sumCart1 + items[i] <= max) {
			sumCart1 += items[i]
			cart1 = append(cart1, items[i])
		} else {
			sumCart2 += items[i]
			cart2 = append(cart2, items[i])
		}
	}
	
	diff := sumCart2 - sumCart1

	return cart1, cart2, diff
}

func createItems(dataInput []string) ([]int, error) {
	result := []int{}

	for _, v := range dataInput {
		i, err := strconv.Atoi(v)
		if (err != nil) {
			return nil, errors.New("error: isNaN")
		}
		result = append(result, i)
	}

	if len(result) < 2 || len(result) > 100 {
		return nil, errors.New("error: invalid length input")
	}
	return result, nil
}

func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func sortDesc(nums []int) []int {
	for n := len(nums); n > 1; n-- {
		for j := 1; j < n; j++ {
			if nums[j-1] < nums[j] {
				swap(nums, j-1, j)
			}
		}
	}
	return nums
}