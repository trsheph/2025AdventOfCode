package dayone

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcDayOne(inFile string) (passwd int) {
	var rowSlice []string
	var fileString string
	var pwAdd int
	currNum := 50
	passwd = 0
	instDir := make(map[int]bool) // Instruction Direction bool True is right
	instNum := make(map[int]int)  // Instruction Number int
	dat, err := os.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fileString = string(dat)
	rowSlice = strings.Split(fileString, "\n")
	for rowNum, rowA := range rowSlice {
		if len(rowA) > 1 {
			if string(rowA[0]) == "L" {
				instDir[rowNum] = false
			} else {
				instDir[rowNum] = true
			}
			moveNum, err := strconv.Atoi(string(rowA[1:]))
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			instNum[rowNum] = moveNum
		}
	}
	for i := 0; i < len(instNum); i++ {
		cN := instNum[i]
		pwAdd = 0
		if instDir[i] {
			currNum = currNum + cN
			pwAdd = currNum / 100
			currNum = (currNum%100 + 100) % 100
		} else {
			if (currNum > 0) && (currNum-cN < 0) {
				pwAdd = pwAdd + 1
			}
			currNum = currNum - cN
			pwCache := -currNum / 100
			pwAdd = pwAdd + pwCache
			currNum = (currNum%100 + 100) % 100
		}
		if (currNum == 0) && (pwAdd == 0) {
			pwAdd = 1
		}
		passwd = passwd + pwAdd
	}
	fmt.Println(passwd)
	return
}
