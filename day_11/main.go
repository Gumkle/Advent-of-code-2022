package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monke struct {
	items          []int
	op             func(int) int
	test           func([]*monke, int) *monke
	itemsInspected int
	divisibleBy    int
}

func main() {
	input := bufio.NewReader(os.Stdin)
	monkeys := make([]*monke, 0)
	divider := 1
	for {
		_, _, err := input.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		itemsLine, _, err := input.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		rawItems := strings.Split(strings.Split(string(itemsLine), ": ")[1], ", ")
		items := make([]int, 0)
		for _, v := range rawItems {
			item, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			items = append(items, item)
		}

		funcLine, _, err := input.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		rawFunc := strings.Split(string(funcLine), "= ")
		funcElements := strings.Split(rawFunc[1], " ")
		op := func(item int) int {
			var firstElement, secondElement int
			if funcElements[0] == "old" {
				firstElement = item
			} else {
				firstElementConverted, err := strconv.Atoi(funcElements[0])
				if err != nil {
					log.Fatalln(err)
				}
				firstElement = firstElementConverted
			}
			if funcElements[2] == "old" {
				secondElement = item
			} else {
				secondElementConverted, err := strconv.Atoi(funcElements[2])
				if err != nil {
					log.Fatalln(err)
				}
				secondElement = secondElementConverted
			}
			if funcElements[1] == "*" {
				return firstElement * secondElement
			}
			return firstElement + secondElement
		}

		testLine, _, err := input.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		divisibleByRaw := strings.Split(string(testLine), "divisible by ")[1]
		divisibleBy, err := strconv.Atoi(divisibleByRaw)
		if err != nil {
			log.Fatalln(err)
		}

		trueLine, _, err := input.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		trueNumber, err := strconv.Atoi(string(trueLine[len(trueLine)-1]))
		if err != nil {
			log.Fatalln(err)
		}

		falseLine, _, err := input.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		falseNumber, err := strconv.Atoi(string(falseLine[len(falseLine)-1]))
		if err != nil {
			log.Fatalln(err)
		}
		test := func(monkeys []*monke, item int) *monke {
			if item%divisibleBy == 0 {
				return monkeys[trueNumber]
			}
			return monkeys[falseNumber]
		}
		monkeys = append(monkeys, &monke{
			items:       items,
			test:        test,
			op:          op,
			divisibleBy: divisibleBy,
		})
		divider *= divisibleBy
		_, _, err = input.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
	}
	for i := 0; i < 10000; i++ {
		for _, mo := range monkeys {
			mo.itemsInspected += len(mo.items)
			for _, item := range mo.items {
				newItemValue := mo.op(item)
				nextMonke := mo.test(monkeys, newItemValue)
				newItemValue %= divider
				//primes := PrimeFactors(newItemValue)
				//newItemValue = lo.Reduce[int, int](lo.Uniq[int](primes), func(agg int, item int, index int) int {
				//	return agg * item
				//}, 1)
				nextMonke.items = append(nextMonke.items, newItemValue)
			}
			mo.items = make([]int, 0)
		}
	}

	//for _, v := range monkeys {
	//	fmt.Println(v.itemsInspected)
	//}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].itemsInspected > monkeys[j].itemsInspected
	})
	fmt.Printf("monkey business: %d\n", monkeys[0].itemsInspected*monkeys[1].itemsInspected)
}

func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
