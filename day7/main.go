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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dirSizes := make(map[string]int)
	currentDir := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		cmdRegex := regexp.MustCompile(`^\$ (\w+)\s?(.*)$`)
		cmdMatch := cmdRegex.FindStringSubmatch(line)

		sizeRegex := regexp.MustCompile(`^(\d+)`)
		sizeMatch := sizeRegex.FindStringSubmatch(line)

		if len(cmdMatch) > 0 && cmdMatch[1] == "cd" {
			currentDir = changeDir(cmdMatch[2], currentDir)
		} else if len(sizeMatch) > 0 {
			size, _ := strconv.Atoi(sizeMatch[1])
			addSizetoDir(currentDir, dirSizes, size)
		}
	}

	fmt.Println("Part1: Total of dirs", totalOfDirsUnder(100001, dirSizes))

	totalUsed := dirSizes[""]
	totalCurrentlyFree := 70000000 - totalUsed
	freeSpaceNeeded := 30000000 - totalCurrentlyFree

	fmt.Println("Part 2: size of dir to delete: ", SmallestDirOver(freeSpaceNeeded, dirSizes))
}

func changeDir(newDir string, currentDir string) string {
	switch newDir {
	case "/":
		return ""
	case "..":
		dirArr := strings.Split(currentDir, "/")
		return strings.Join(dirArr[:len(dirArr)-1], "/")
	}
	return currentDir + "/" + newDir
}

func addSizetoDir(dir string, dirSizes map[string]int, size int) {
	dirArr := strings.Split(dir, "/")
	for i := 0; i < len(dirArr); i++ {
		dir := strings.Join(dirArr[:i+1], "/")
		currentSize := dirSizes[dir]
		dirSizes[dir] = currentSize + size
	}
}

func totalOfDirsUnder(size int, dirSizes map[string]int) int {
	total := 0
	for _, dirSize := range dirSizes {
		if dirSize < size {
			total += dirSize
		}
	}
	return total
}

func SmallestDirOver(size int, dirSizes map[string]int) int {
	smallest := 0
	for _, dirSize := range dirSizes {
		if dirSize > size && (smallest == 0 || dirSize < smallest) {
			smallest = dirSize
		}
	}
	return smallest
}
