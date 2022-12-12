package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	stream := string(input)

	monkeyInputs := strings.Split(stream, "\n\n")

	doMonkeyBusiness(monkeyInputs, 20, 3)
	doMonkeyBusiness(monkeyInputs, 10000, 1)
}

func doMonkeyBusiness(monkeyInputs []string, rounds int, worryReducer int) {

	monkeys := make([]Monkey, 0, len(monkeyInputs))

	divisibilities := make([]int, 0, len(monkeyInputs))

	for _, monkeyInput := range monkeyInputs {
		monkey := configureMonkey(monkeyInput)
		monkeys = append(monkeys, monkey)
		divisibilities = append(divisibilities, monkey.test)
	}

	commonDivisibility := 1
	for _, divisibility := range divisibilities {
		commonDivisibility *= divisibility
	}

	for round := 1; round <= rounds; round++ {
		for monkeyNumber := 0; monkeyNumber < len(monkeys); monkeyNumber++ {
			monkey := &monkeys[monkeyNumber]
			monkey.performInspection(monkeys, worryReducer, commonDivisibility)

		}
		fmt.Println("Round", round, "complete", monkeys)
	}

	monkeyInspections := make([]int, 0, len(monkeys))
	for _, monkey := range monkeys {
		monkeyInspections = append(monkeyInspections, monkey.inspectionCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyInspections)))

	monkeyBusiness := monkeyInspections[0] * monkeyInspections[1]
	fmt.Println("Monkey business:", monkeyBusiness)
}

type Monkey struct {
	items           []int
	operation       []string
	test            int
	pass            int
	fail            int
	inspectionCount int
}

func configureMonkey(input string) Monkey {
	return Monkey{
		items:           parseItems(input),
		operation:       parseOperation(input),
		test:            parseTest(input),
		pass:            parsePass(input),
		fail:            parseFail(input),
		inspectionCount: 0,
	}
}

func parseItems(input string) []int {
	itemsRegex := regexp.MustCompile(`Starting items: (.*)`)

	itemsString := itemsRegex.FindStringSubmatch(input)[1]
	stringItems := strings.Split(itemsString, ", ")
	items := make([]int, 0, len(stringItems))
	for _, item := range stringItems {
		item, _ := strconv.Atoi(item)
		items = append(items, item)
	}
	return items
}

func parseOperation(input string) []string {
	operationRegex := regexp.MustCompile(`Operation: new = (\w+) ([+*]) (\w+)`)

	return operationRegex.FindStringSubmatch(input)[1:]
}

func parseTest(input string) int {
	testRegex := regexp.MustCompile(`Test: divisible by (.*)`)
	testString := testRegex.FindStringSubmatch(input)[1]
	test, _ := strconv.Atoi(testString)
	return test
}

func parsePass(input string) int {
	passRegex := regexp.MustCompile(`If true: throw to monkey (.*)`)
	passString := passRegex.FindStringSubmatch(input)[1]
	pass, _ := strconv.Atoi(passString)
	return pass
}

func parseFail(input string) int {
	failRegex := regexp.MustCompile(`If false: throw to monkey (.*)`)
	failString := failRegex.FindStringSubmatch(input)[1]
	fail, _ := strconv.Atoi(failString)
	return fail
}

func (m *Monkey) performInspection(monkeys []Monkey, worryReducer int, commonDivisibility int) {
	for _, item := range m.items {
		m.inspectionCount++

		worryLevel := m.calculateWorryLevel(item)

		if worryLevel > commonDivisibility {
			fmt.Println("Worry level:", worryLevel, "common divisibility:", commonDivisibility)
			worrySubtractor := (worryLevel / commonDivisibility) * commonDivisibility
			worryLevel -= worrySubtractor
			fmt.Println("Worry subtractor:", worrySubtractor)
		} else {
			worryLevel = worryLevel / worryReducer
		}

		testResult := worryLevel%m.test == 0
		if testResult {
			passItemTo(&(monkeys[m.pass]), worryLevel)
		} else {
			passItemTo(&(monkeys[m.fail]), worryLevel)
		}
	}

	m.items = []int{}
}

func (m Monkey) calculateWorryLevel(item int) int {

	var worryLevel int
	left, right := oldOrValue(m.operation[0], item), oldOrValue(m.operation[2], item)
	switch m.operation[1] {
	case "+":
		worryLevel = left + right
	case "*":
		worryLevel = left * right
	}
	fmt.Println(left, m.operation[1], right, "=", worryLevel)
	return worryLevel
}

func oldOrValue(value string, old int) int {
	if value == "old" {
		return old
	}
	new, _ := strconv.Atoi(value)
	return new
}

func passItemTo(monkey *Monkey, item int) {
	monkey.items = append(monkey.items, item)
}
