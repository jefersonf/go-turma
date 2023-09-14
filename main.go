package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parseData()
}

func parseData() {
	turmaFile, err := os.Open("data/turma.csv")
	checkErr(err)

	turmaScanner := bufio.NewScanner(turmaFile)
	turmaScanner.Split(bufio.ScanLines)

	for turmaScanner.Scan() {
		fmt.Println(turmaScanner.Text())
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
