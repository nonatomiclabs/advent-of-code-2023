package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nonatomiclabs/advent-of-code-2023/day_01"
	"github.com/nonatomiclabs/advent-of-code-2023/day_02"
	"github.com/nonatomiclabs/advent-of-code-2023/day_03"
	"github.com/nonatomiclabs/advent-of-code-2023/day_04"
	"github.com/nonatomiclabs/advent-of-code-2023/day_05"
	"github.com/nonatomiclabs/advent-of-code-2023/day_06"
)

func main() {
	day := flag.Int("d", 1, "the day of the calendar to run")
	secondPart := flag.Bool("s", false, "whether to run the algorithm for the second part of the solution")
	flag.Parse()

	solutions := map[int]func([]string, bool) int{
		1: day_01.Solution,
		2: day_02.Solution,
		3: day_03.Solution,
		4: day_04.Solution,
		5: day_05.Solution,
		6: day_06.Solution,
	}

	daySolution, ok := solutions[*day]
	if !ok {
		fmt.Fprintf(os.Stderr, "solution for day %d not implemented yet!\n", *day)
		os.Exit(1)
	}

	fmt.Printf("Loading solution for day %d, ", *day)
	if *secondPart {
		fmt.Printf("part two\n")
	} else {
		fmt.Printf("part one\n")
	}

	file, err := os.Open(fmt.Sprintf("day_%02d/input", *day))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open input file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var inputLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	daySolution(inputLines, *secondPart)
}
