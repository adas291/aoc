package day_01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// func test() {

// 	m:=regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`).FindStringSubmatch("2023-12-02")

// 	println(m[1])
// }

func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		panic("not implemented")
		// fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func Part1(input []string) string {
	file, err := os.Open("input/1.txt")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d`)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := re.FindAllString(line, -1)
		first, _ := strconv.Atoi(nums[0])
		last, _ := strconv.Atoi(nums[len(nums)-1])
		sum += (first*10 + last)
	}
	return fmt.Sprint(sum)
}
