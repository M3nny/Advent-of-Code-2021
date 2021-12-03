package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sync"
)

var wg sync.WaitGroup
var gamma [12]int
var epsilon [12]int

func bit_0() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[0] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[0] = 0b0
		} else {
			gamma[0] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_1() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[1] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[1] = 0b0
		} else {
			gamma[1] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_2() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[2] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[2] = 0b0
		} else {
			gamma[2] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_3() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[3] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[3] = 0b0
		} else {
			gamma[3] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_4() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[4] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[4] = 0b0
		} else {
			gamma[4] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_5() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[5] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[5] = 0b0
		} else {
			gamma[5] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_6() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[6] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[6] = 0b0
		} else {
			gamma[6] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_7() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[7] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[7] = 0b0
		} else {
			gamma[7] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_8() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[8] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[8] = 0b0
		} else {
			gamma[8] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_9() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[9] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[9] = 0b0
		} else {
			gamma[9] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_10() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[10] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[10] = 0b0
		} else {
			gamma[10] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func bit_11() {
	var x int = 0
	var y int = 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text()[11] == 48 {
			x += 1

		} else {
			y += 1
		}

		if x > y {
			gamma[11] = 0b0
		} else {
			gamma[11] = 0b1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func converter(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}

func concatenate(array [12]int) int {
	var concat int = 0
	for i := 0; i < len(array); i++ {
		concat = concat*10 + array[i]
	}
	return converter(concat)

}
func main() {
	var gammaDec int
	var epsilonDec int
	var power_consumption int

	wg.Add(1)
	go bit_0()
	wg.Add(1)
	go bit_1()
	wg.Add(1)
	go bit_2()
	wg.Add(1)
	go bit_3()
	wg.Add(1)
	go bit_4()
	wg.Add(1)
	go bit_5()
	wg.Add(1)
	go bit_6()
	wg.Add(1)
	go bit_7()
	wg.Add(1)
	go bit_8()
	wg.Add(1)
	go bit_9()
	wg.Add(1)
	go bit_10()
	wg.Add(1)
	go bit_11()

	wg.Wait()

	fmt.Println("gamma in binary:", gamma)

	for i := 0; i < len(gamma); i++ {
		if gamma[i] == 0 {
			epsilon[i] = 1
		} else {
			epsilon[i] = 0
		}
	}

	fmt.Println("epsilon in binary:", epsilon)

	gammaDec = concatenate(gamma)
	epsilonDec = concatenate(epsilon)
	power_consumption = gammaDec * epsilonDec

	fmt.Println("gamma in decimal:", gammaDec)
	fmt.Println("epsilon in decimal:", epsilonDec)
	fmt.Println("your power consumption:", power_consumption, "watt")

}
