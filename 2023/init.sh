#!/bin/bash

for day in {04..31}; do
  folder_name="day_${day}"
  file_name="${folder_name}/main.go"

  mkdir -p "${folder_name}"

  cat <<EOF >"${file_name}"
package ${folder_name}

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
    return "Day ${day} - Part 1 logic"
}

func Part2(input []string) string {
    return "Day ${day} - Part 2 logic"
}
EOF

done

