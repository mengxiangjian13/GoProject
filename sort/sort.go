package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mengxiangjian13/GoProject/sort/bubblesort"
	"io"
	"os"
	"strconv"
)

func readValues(infile string) (values []int, err error) {
	// 打开文件
	file, err := os.Open(infile)

	if err != nil {
		fmt.Println("failed to open the input file", infile)
	}

	defer file.Close()

	br := bufio.NewReader(file) // 通过文件，创建newReader

	values = make([]int, 0)

	for {
		line, isPerfix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPerfix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed to create the output file", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

// 解析命令行传入参数
var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")

func main() {

	// 解析外部传入参数。
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("read values:", values)
		bubblesort.Bubblesort(values)
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
