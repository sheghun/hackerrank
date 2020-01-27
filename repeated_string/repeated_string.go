// Copy right (c) Oladiran Segun Solomon
// 2020
// This code is written for educational purposes only
// Github url <https://github.com/sheghun/hackerrank>
// Some modifications are made to this code see line 26
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	// Check if number of strings to consider is greater than 10 ^ 12
	if n > int64(math.Pow10(12)) {
		panic(errors.New("integer is greater than 10^12"))
	}

	var count int

	if len(s) == 1 {
		if s == "a" {
			return n
		}
		return 0
	}

	if n < int64(len(s)) {
		str := s[:n]

		count = strings.Count(str, "a")

	}

	if n >= int64(len(s)) {
		// Get the reminder
		rem := n % int64(len(s))
		fac := int(math.Floor(float64(n) / float64(len(s))))

		a := strings.Count(s, "a")

		count = a * fac

		if rem != 0 {
			str := s[:rem]

			count += strings.Count(str, "a")

		}

	}

	return int64(count)

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

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
