package day_02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type set struct {
	red   int
	green int
	blue  int
}

func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func Part1(input []string) string {

	R := 12
	G := 13
	B := 14

	file, err := os.Open("input/2.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	validGames := 0

outer:
	for scanner.Scan() {
		line := scanner.Text()
		parts := regexp.MustCompile(`Game (\d+): (.+)`).FindStringSubmatch(line)
		gameId, _ := strconv.Atoi(parts[1])
		sets := strings.Split(parts[2], `;`)

		for _, set := range sets {
			r, g, b := 0, 0, 0
			gameData := strings.Split(set, `,`)

			for _, item := range gameData {
				m := regexp.MustCompile(`(\d+) (\w)`).FindStringSubmatch(item)
				cnt, _ := strconv.Atoi(m[1])
				switch m[2] {
				case "r":
					r += cnt
				case "b":
					b += cnt
				case "g":
					g += cnt
				}
				if r > R || g > G || b > B {
					continue outer
				}
			}
		}
		validGames += gameId
	}

	return fmt.Sprint(validGames)
}

func Part2(input []string) string {
	sum := 0
	for _, game := range input {
		sum += calculateMinPower(game)
	}
	return strconv.Itoa(sum)
}

// calculateMinPower calculates the power of the smallest set of cubes which still make the game possible
func calculateMinPower(game string) int {
	sets := getSets(game)
	cubes := calculateMinCubes(sets)
	return cubes.red * cubes.green * cubes.blue
}

// calculateMinCubes finds the smallest required value of each color to make the game possible.
// The smallest required value is always the largest number in the input for the respective color.
func calculateMinCubes(sets []set) set {
	var m set
	for _, s := range sets {
		if s.red > m.red {
			m.red = s.red
		}
		if s.green > m.green {
			m.green = s.green
		}
		if s.blue > m.blue {
			m.blue = s.blue
		}
	}
	return m
}

// getSets parses a single input game and creates a slice of sets out of it
func getSets(game string) []set {
	gameWithoutId := strings.Split(game, ": ")[1]
	setStrings := strings.Split(gameWithoutId, "; ")
	var sets []set
	for _, s := range setStrings {
		sets = append(sets, getSet(s))
	}
	return sets
}

// getSet receives a set as string and creates a set struct from it
func getSet(setString string) set {
	var s set
	cubes := strings.Split(setString, ", ")
	for _, cube := range cubes {
		count, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		if strings.HasSuffix(cube, "red") {
			s.red = count
		} else if strings.HasSuffix(cube, "green") {
			s.green = count
		} else if strings.HasSuffix(cube, "blue") {
			s.blue = count
		}
	}
	return s
}