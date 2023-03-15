package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fName := "test.txt.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

//func main() {
//	file, err := os.Open("products2.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	var col1, col2, col3 []string
//	for {
//		var v1, v2, v3 string
//		_, err := fmt.Fscanln(file, &v1, &v2, &v3) // scans until newline
//		if err != nil {
//			break
//		}
//		col1 = append(col1, v1)
//		col2 = append(col2, v2)
//		col3 = append(col3, v3)
//	}
//	fmt.Println(col1)
//	fmt.Println(col2)
//	fmt.Println(col3)
//}

//func main() {
//	inputFile := "products.txt"
//	buf, err := ioutil.ReadFile(inputFile)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
//	}
//	fmt.Printf("%s\n", string(buf))
//}

//
//func main() {
//	inputFile, inputError := os.Open("input.dat")
//	if inputError != nil {
//		fmt.Printf("An error occurred on opening the inputfile\n" +
//			"Does the file exist?\n" +
//			"Have you got access to it?\n")
//		return // exit the function on error
//	}
//	defer inputFile.Close()
//	inputReader := bufio.NewReader(inputFile)
//	for {
//		inputString, readerError := inputReader.ReadString('\n')
//		if readerError == io.EOF {
//			return
//		}
//		fmt.Printf("The input was: %s", inputString)
//	}
//}
