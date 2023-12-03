package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	rxRed   = regexp.MustCompile(`\d+\b( red)\b`)
	rxBlue  = regexp.MustCompile(`\d+\b( blue)\b`)
	rxGreen = regexp.MustCompile(`\d+\b( green)\b`)
)

func solution(fileName string) int {
	_ = parse(fileName)
	return 8
}

type M map[string]int

func parse(fileName string) any {
	lines := readLines(fileName)

	greenNum := 0
	blueNum := 0
	redNum := 0

	greens := make([][]string, len(lines))
	blues := make([][]string, len(lines))
	reds := make([][]string, len(lines))

	games := make([][3]int, len(lines))

	for i, line := range lines {
		greens[i] = rxGreen.FindAllString(line, -1)
		blues[i] = rxBlue.FindAllString(line, -1)
		reds[i] = rxRed.FindAllString(line, -1)
	}

	for i, match := range greens {
		for _, row := range match {
			greenNum = max(getNum(row), greenNum)
			games[i][0] = greenNum
		}
		greenNum = 0
	}

	for i, match := range blues {
		for _, row := range match {
			blueNum = max(getNum(row), blueNum)
			games[i][1] = blueNum
		}
		blueNum = 0
	}

	for i, match := range reds {
		for _, row := range match {
			redNum = max(getNum(row), redNum)
			games[i][2] = redNum
		}
		redNum = 0
	}

	fmt.Println(lines)
	fmt.Println(games)

	return greenNum
}

func getNum(set string) int {
	digit, _, _ := strings.Cut(set, " ")
	num, err := strconv.Atoi(digit)
	if err != nil {
		panic(err)
	}
	return num
}

func readLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
