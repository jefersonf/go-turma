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

	// Skip first since it is a header
	turmaScanner.Scan()

	for turmaScanner.Scan() {
		fmt.Println(turmaScanner.Text())
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
