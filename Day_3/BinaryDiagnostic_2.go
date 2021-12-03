package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func converter(diagnostic string) int {
	result, err := strconv.ParseInt(diagnostic, 2, 64)
	if err != nil {
		return 0
	}
	return int(result)
}

func oxygen_generator_rating(bits []string) string {
	pos := 0
	for len(bits) != 1 {
		bits = oxygen_filter(bits, pos)
		pos++
	}
	return bits[0]
}

func c02_scrubber_rating(bits []string) string {
	pos := 0
	for len(bits) != 1 {
		bits = c02_filter(bits, pos)
		pos++
	}
	return bits[0]
}

func oxygen_filter(bits []string, pos int) []string {
	var result []string
	var MSB uint8 //most significative bit

	x := 0
	y := 0
	for _, line := range bits {
		if line[pos] == '0' {
			x++
		} else {
			y++
		}
	}

	if y >= x {
		MSB = '1'
	} else {
		MSB = '0'
	}
	for _, diagnostic := range bits {
		if diagnostic[pos] == MSB {
			result = append(result, diagnostic)
		}
	}
	return result
}

func c02_filter(bits []string, pos int) []string {
	var result []string
	var MSB uint8

	x := 0
	y := 0
	for _, line := range bits {
		if line[pos] == '0' {
			x++
		} else {
			y++
		}
	}

	if y < x {
		MSB = '1'
	} else {
		MSB = '0'
	}
	for _, n := range bits {
		if n[pos] == MSB {
			result = append(result, n)
		}
	}
	return result
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(file)
	var bits []string
	for scanner.Scan() {
		line := scanner.Text()
		bits = append(bits, line)
	}

	err = file.Close()
	if err != nil {
		return
	}

	diagnosticOxygen := oxygen_generator_rating(bits)
	diagnosticCo2 := c02_scrubber_rating(bits)

	oxygen := converter(diagnosticOxygen)
	co2 := converter(diagnosticCo2)

	fmt.Println("Oxygen:", oxygen, "Co2:", co2)
	fmt.Println("Life support rating:", oxygen*co2)

}
