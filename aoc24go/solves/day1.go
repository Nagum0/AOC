package solves

import (
	"aoc24/solves/utils"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getLeftAndRight(splitDataString []string) (lefts []int, rights []int) {
    for idx, val := range splitDataString {
        parsedVal, err := strconv.Atoi(val)

        if err != nil {
            fmt.Printf("Error while parsing value: %v (%v)\n", val, err.Error())
            os.Exit(1)
        }

        if idx % 2 == 0 {
            lefts = append(lefts, parsedVal)
        } else {
            rights = append(rights, parsedVal)
        }
    }
    
    return lefts, rights
}

func calcDistances(lefts []int, rights []int) []int {
    sort.Ints(lefts)
    sort.Ints(rights)
       
    distances := []int{}

    for idx, val := range lefts {
        distance := math.Abs(float64(val - rights[idx]))
        distances = append(distances, int(distance))
    }
    
    return distances
}

func sumDistances(distances []int) int {
    sum := 0
    
    for _, val := range distances {
        sum += val
    }

    return sum
}

func part1(fileContents string) {
    lefts, rights := getLeftAndRight(strings.Fields(fileContents))
    distances := calcDistances(lefts, rights)
    result := sumDistances(distances)
    fmt.Printf("Part 1: %v\n", result)
}

func getMax(xs []int) int {
    m := xs[0]
    
    for _, val := range xs[1:] {
        if val > m {
            m = val
        }
    }

    return m
}

func calcSimScores(lefts []int, rights []int) []int {
    n := getMax(append(lefts, rights...))
    counts := make([]int, n)

    for _, val := range rights {
        counts[val - 1]++
    }

    simScores := []int{}

    for _, val := range lefts {
        score := val * counts[val - 1]
        simScores = append(simScores, score)
    }

    return simScores
}

func part2(fileContents string) {
    lefts, rights := getLeftAndRight(strings.Fields(fileContents))
    simScores := calcSimScores(lefts, rights)
    result := sumDistances(simScores)
    fmt.Printf("Par 2: %v\n", result)
}

func Day1() {
    fileContents := utils.ReadFileToString("../aoc24/data/day1_main.txt")
    part1(fileContents)
    part2(fileContents)
}
