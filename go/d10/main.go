package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	fmt.Println(binNum(int64(n)))
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

func binNum(in int64) int {
	ins := strconv.FormatInt(in, 2)
	// fmt.Println(ins)
	vs := strings.Split(ins, "0")
	res := 1
	for _, v := range vs {
		if l := len(v); l > res {
			res = l
		}
	}
	return res
}
