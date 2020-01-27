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

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	jumps := 0

	for i := 0; i < len(c); {
		currentCloud := c[i]

		if currentCloud == 0 && i != 0 {
			jumps++
		}

		// Check to see if the next two clouds are safe then jump one
		if i+1 != len(c) && i+2 != len(c) {
			nextCloud := c[i+1]
			nextSecondCloud := c[i+2]

			if currentCloud == 0 && nextCloud == 0 && nextSecondCloud == 0 {
				i += 2
				continue
			}

		}

		i++
	}

	return int32(jumps)
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

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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
