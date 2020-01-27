package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {

	var total int32
	total = -1_000_000

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			//fmt.Println(i)
			// First row 3 numbers
			fI := arr[i][j]
			sI := arr[i][j+1]
			tI := arr[i][j+2]

			// second row mid number
			mS := arr[i+1][j+1]

			// Third row 3 numbers
			fT := arr[i+2][j]
			sT := arr[i+2][j+1]
			tT := arr[i+2][j+2]

			// Add the three items together
			newTotal := fI + sI + tI + mS + fT + sT + tT

			if newTotal > total {
				total = newTotal
			}
		}
	}

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	// The next line of code was commented out so as to make the code
	// Compatible with local machine, instead of using the hackerrank
	// Output path we are just printing directly to the terminal

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	// Write directly to Stdout as the local machine does share the same environment variables with hackerrank
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
