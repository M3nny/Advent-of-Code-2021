package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var forward int = 0
var depth int = 0

func main() {
	var str string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str = scanner.Text()
		sub1 := strings.Contains(str, "forward")
		sub2 := strings.Contains(str, "up")
		sub3 := strings.Contains(str, "down")

		//forward
		if sub1 == true {
			var fwd string
			var intvar int

			re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
			match := re.FindAllString(str, -1)

			fwd = strings.Trim(strings.Replace(fmt.Sprint(match), " ", fwd, -1), "[]")
			intvar, err = strconv.Atoi(fwd)
			forward += intvar

		}
		//up
		if sub2 == true {
			var fwd string
			var intvar int

			re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
			match := re.FindAllString(str, -1)

			fwd = strings.Trim(strings.Replace(fmt.Sprint(match), " ", fwd, -1), "[]")
			intvar, err = strconv.Atoi(fwd)
			depth -= intvar

		}
		//down
		if sub3 == true {
			var fwd string
			var intvar int

			re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
			match := re.FindAllString(str, -1)

			fwd = strings.Trim(strings.Replace(fmt.Sprint(match), " ", fwd, -1), "[]")
			intvar, err = strconv.Atoi(fwd)
			depth += intvar

		}

	}
	fmt.Println("Your position is: ", (depth * forward))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
