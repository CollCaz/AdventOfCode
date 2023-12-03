package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func solution(fileName string) int {
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

	var sum int

	for _, line := range lines {
		firstNum, lastNum := "", ""
		firstDigits(line, &firstNum)
		lastDigits(line, &lastNum)
		combo, err := strconv.Atoi(fmt.Sprintf("%s%s", firstNum, lastNum))
		if err != nil {
			panic(err)
		}
		sum += combo
	}
	return sum
}

func firstDigits(l string, firstNum *string) string {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var err error
	if unicode.IsNumber(rune(l[0])) {
		*firstNum = string(l[0])
		if err != nil {
			panic(err)
		}
	}

	for i, num := range nums {
		if strings.HasPrefix(l, num) {
			*firstNum = fmt.Sprint(i + 1)
			break
		}
	}

	if *firstNum == "" {
		if len(l) != 1 {
			l = l[1:]
			firstDigits(l, firstNum)
		}
	}
	return *firstNum
}

func lastDigits(l string, lastNum *string) string {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var err error
	if unicode.IsNumber(rune(l[len(l)-1])) {
		*lastNum = string(l[len(l)-1])
		if err != nil {
			panic(err)
		}
	}

	for i, num := range nums {
		if strings.HasSuffix(l, num) {
			*lastNum = fmt.Sprint(i + 1)
			break
		}
	}

	if *lastNum == "" {
		if len(l) > 0 {
			l = l[0:(len(l) - 1)]
			lastDigits(l, lastNum)
		}
	}
	return *lastNum
}
