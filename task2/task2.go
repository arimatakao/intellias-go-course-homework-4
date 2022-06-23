package task2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Task2() {
	nums := "1 9 3 4 -5"
	var result string
	var max int
	var min int

	arr := strings.Split(nums, " ")
	for _, v := range arr {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	result = fmt.Sprintf("%d %d", max, min)
	fmt.Println(nums)
	fmt.Println(result)
}
