package internal

import (
	"m/day_01"
	"m/day_02"
	"m/day_03"
	"m/day_04"
	"m/day_05"
	"m/day_06"
	"m/day_07"
	"m/day_08"
	"m/day_09"
	"m/day_10"
	"m/day_11"
	"m/day_12"
	"m/day_13"
	"m/day_14"
	"m/day_15"
	"m/day_16"
	"m/day_17"
	"m/day_18"
	"m/day_19"
	"m/day_20"
	"m/day_21"
	"m/day_22"
	"m/day_23"
	"m/day_24"
	"m/day_25"
)

// RunChallenge executes the challenge of a specific day with the provided input.
func RunChallenge(day int, inputPath string, mode int) {
	input := LoadInputLines(inputPath)
	mapping := map[int]func([]string, int){
		1:  day_01.Run,
		2:  day_02.Run,
		3:  day_03.Run,
		4:  day_04.Run,
		5:  day_05.Run,
		6:  day_06.Run,
		7:  day_07.Run,
		8:  day_08.Run,
		9:  day_09.Run,
		10:  day_10.Run,
		11:  day_11.Run,
		12:  day_12.Run,
		13:  day_13.Run,
		14:  day_14.Run,
		15:  day_15.Run,
		16:  day_16.Run,
		17:  day_17.Run,
		18:  day_18.Run,
		19:  day_19.Run,
		20:  day_20.Run,
		21:  day_21.Run,
		22:  day_22.Run,
		23:  day_23.Run,
		24:  day_24.Run,
		25:  day_25.Run,
	}

	mapping[day](input, mode)
}