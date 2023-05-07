package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	// Write your code here
	integerRepresentation, err := strconv.Atoi(n)
	if err == nil && integerRepresentation <= 9 && k == 1 {
		return int32(integerRepresentation)
	}
	repeatedString := strings.Repeat(n, int(k))
	fmt.Printf("repeated string %s", repeatedString)

	currentSuperDigit := 0
	for i := 0; i < len(repeatedString); i++ {

		convertedInt := int(repeatedString[i] - '0')
		currentSuperDigit += convertedInt
	}
	fmt.Printf("Current Super Digit is %d and the Itoa is: %s", currentSuperDigit, strconv.Itoa(currentSuperDigit))

	return superDigit(strconv.Itoa(currentSuperDigit), 1)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
