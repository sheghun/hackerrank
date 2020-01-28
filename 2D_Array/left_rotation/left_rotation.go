package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {

	aCopy := make([]int32, len(a))

	l := int32(len(a))

	fmt.Println("Length", l)

	for i := int32(0); i < l; i++ {
		rotIn := (l - d) + i

		if rotIn == l {
			rotIn = 0
		}

		if rotIn > l {
			rotIn = rotIn - l
		}

		aCopy[rotIn] = a[i]

	}

	return aCopy
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

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
