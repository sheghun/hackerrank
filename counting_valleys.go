// Copy right (c) Oladiran Segun Solomon
// 2020
// This code is written for educational purposes only
// Github url <https://github.com/sheghun/hackerrank>
// Some modifications are made to this code see line 26
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"strconv"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	// The number of valleys entered
	valley := 0

	// This the level of step taken
	// It is initialized to the sea level the steps will start from sea level
	// Any downhill step will result in -1 and any uphill step will result in +1
	// At the end of any consecutive set, depending on the number of times the step level
	// Is set 0 it means we entered and completed a valley
	stepLevel := 0

	for i := 0; int32(i) < n; i++ {
		currentStep := string(s[i])

		// Check if current step is an uphill skip to next iteration
		if currentStep == "U" {
			// Add 1 from step level
			stepLevel += 1

			// Check if step level is zero and is not first iteration
			if stepLevel == 0 && i != 0 {
				// Check if the last step is not a downhill
					// It means it's a valley increment valley
					valley++
			}
			continue
		}

		// Check if current step is a downhill
		if currentStep == "D" {
			// Subtract 1 from  step level
			stepLevel -= 1
		}
	}
	return int32(valley)
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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

	_, _ = fmt.Fprintf(writer, "%d\n", result)

	_ = writer.Flush()
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
