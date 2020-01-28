package main

import (
	"bufio"
	"fmt"

	//"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) string {
	b := 0

	l := len(q)

	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l && q[j-1] > q[j]; j++ {
			if j-i > 2 {
				return "Too chaotic"
			}

			q[j-1], q[j] = q[j], q[j-1]
			b++
		}
	}

	return fmt.Sprintf("%d", b)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	var output []string

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		b := minimumBribes(q)
		output = append(output, b)
	}

	fmt.Printf(strings.Join(output, "\n"))
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
