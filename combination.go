package main

import "fmt"

func uidCombination(uids []string, len int, r int) <-chan []string {
	c := make(chan []string)
	go func(c chan []string) {
		defer close(c)
		var result = make([]string, r)
		uidCombinationUtil(uids, result, 0, len-1, 0, r, c)

	}(c)
	return c

}

func uidCombinationUtil(uids []string, data []string, start int, end int, index int, r int, c chan []string) {
	if index == r {
		result := []string{}
		result = append(result, data...)
		c <- result

	} else {
		for i := start; i <= end && end-i+1 >= r-index; i++ {
			data[index] = uids[i]
			uidCombinationUtil(uids, data, i+1, end, index+1, r, c)

		}

	}

}

func Combination() {
	result := [][]string{}
	arr := []string{"1", "2", "3", "4", "5"}
	for i := 1; i <= 5; i++ {
		for combination := range uidCombination(arr, 5, i) {
			result = append(result, combination)

		}

	}
	fmt.Println(result)

}
