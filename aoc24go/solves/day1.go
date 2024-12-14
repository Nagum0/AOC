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

func Day1() {
    fileContents := utils.ReadFileToString("../aoc24/data/day1_main.txt")
    lefts, rights := getLeftAndRight(strings.Fields(fileContents))
    distances := calcDistances(lefts, rights)
    result := sumDistances(distances)
    fmt.Printf("%v\n", result)
}
