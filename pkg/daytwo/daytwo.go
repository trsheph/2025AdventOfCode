package daytwo

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcDayTwo(inFile string) (sumOut int) {
	var fileString string
	var rangeSlice []string
	var err error
	fmt.Println("Welcome to day 2")
	lows := make(map[int]int)
	highs := make(map[int]int)
	leftLows := make(map[int]int)
	leftHighs := make(map[int]int)
	dat, err := os.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fileString = string(dat)
	rangeSlice = strings.Split(fileString, ",")
	fmt.Println(rangeSlice)
	for rngNum, lowHigh := range rangeSlice {
		lowHighSlice := strings.Split(lowHigh, "-")
		fmt.Println(lowHighSlice)
		if len(lowHighSlice) == 2 {
			lows[rngNum], err = strconv.Atoi(lowHighSlice[0])
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			highs[rngNum], err = strconv.Atoi(lowHighSlice[1])
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			testVal := lows[rngNum]
			if testVal < 10 {
				lows[rngNum] = 10
				lowHighSlice[0] = "10"
			}
			leftLows[rngNum], err = strconv.Atoi(string(string(lowHighSlice[0])[0:(len(lowHighSlice[0]) - (len(lowHighSlice[0])+1)/2)]))
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			leftHighs[rngNum], err = strconv.Atoi(string(string(lowHighSlice[1])[0:(len(lowHighSlice[1]) - (len(lowHighSlice[0])+1)/2)]))
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
		}
	}
	for rngInd, leftLowVal := range leftLows {
		fmt.Println("----")
		fmt.Println(rngInd)
		fmt.Println(lows[rngInd])
		fmt.Println(leftLows[rngInd])
		fmt.Println(leftHighs[rngInd])
		fmt.Println(highs[rngInd])
		for i := leftLowVal; i < leftHighs[rngInd]+1; i++ {
			dblI, err := strconv.Atoi(strconv.Itoa(i) + strconv.Itoa(i))
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			if (dblI >= lows[rngInd]) && (dblI <= highs[rngInd]) {
				fmt.Println(dblI)
				sumOut = sumOut + dblI
			}
		}
	}
	fmt.Println(sumOut)
	return
}
