package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/trsheph/2025AdventOfCode/pkg/dayeight"
	"github.com/trsheph/2025AdventOfCode/pkg/dayeleven"
	"github.com/trsheph/2025AdventOfCode/pkg/dayfive"
	"github.com/trsheph/2025AdventOfCode/pkg/dayfour"
	"github.com/trsheph/2025AdventOfCode/pkg/daynine"
	"github.com/trsheph/2025AdventOfCode/pkg/dayone"
	"github.com/trsheph/2025AdventOfCode/pkg/dayseven"
	"github.com/trsheph/2025AdventOfCode/pkg/daysix"
	"github.com/trsheph/2025AdventOfCode/pkg/dayten"
	"github.com/trsheph/2025AdventOfCode/pkg/daythree"
	"github.com/trsheph/2025AdventOfCode/pkg/daytwelve"
	"github.com/trsheph/2025AdventOfCode/pkg/daytwo"
)

func main() {
	var inFile string
	var err error
	inDay := 1
	var procDos string
	if len(os.Args) == 1 {
		inFile = "inputs/one/test.txt"
	} else if len(os.Args) == 2 {
		inFile = string(os.Args[1])
	} else if len(os.Args) == 3 {
		inDay, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		inFile = string(os.Args[2])
	} else if len(os.Args) == 4 {
		inDay, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		inFile = string(os.Args[2])
		procDos = string(os.Args[3])
	} else {
		err := fmt.Errorf("too many arguments")
		panic(err)
	}
	if inDay == 1 {
		fmt.Println(inFile)
		dayone.ProcDayOne(inFile)
	} else if inDay == 2 {
		fmt.Println(inFile)
		daytwo.ProcDayTwo(inFile)
	} else if inDay == 3 {
		daythree.ProcDayThree(inFile, procDos)
	} else if inDay == 4 {
		dayfour.ProcDayFour(inFile, procDos)
	} else if inDay == 5 {
		dayfive.ProcDayFive(inFile, procDos)
	} else if inDay == 6 {
		daysix.ProcDaySix(inFile, procDos)
	} else if inDay == 7 {
		dayseven.ProcDaySeven(inFile, procDos)
	} else if inDay == 8 {
		dayeight.ProcDayEight(inFile, procDos)
	} else if inDay == 9 {
		daynine.ProcDayNine(inFile, procDos)
	} else if inDay == 10 {
		dayten.ProcDayTen(inFile, procDos)
	} else if inDay == 11 {
		dayeleven.ProcDayEleven(inFile, procDos)
	} else if inDay == 12 {
		daytwelve.ProcDayTwelve(inFile, procDos)
	}
}
