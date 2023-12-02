package day_09

import "fmt"

func Run(input []string, mode int) {
    if mode == 1 || mode == 3 {
        fmt.Printf("Part one: %v\n", Part1(input))
    }
    if mode == 2 || mode == 3 {
        fmt.Printf("Part two: %v\n", Part2(input))
    }
}

func Part1(input []string) string {
    return "Day 09 - Part 1 logic"
}

func Part2(input []string) string {
    return "Day 09 - Part 2 logic"
}
