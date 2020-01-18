package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

func main() {
	route := os.Args[1]
	file, error := os.Open(route)
	if error != nil {
		fmt.Println("Error con el archivo")
	}
	scanner := bufio.NewScanner(file)
	var mat [][]float64
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "LocationDelegate.swift#100") {
			substring1 := line[78:89]
			substring2 := line[90:102]
			//fmt.Println(substring1 + "," + substring2)
			f1, _ := strconv.ParseFloat(substring1, 64)
			f2, _ := strconv.ParseFloat(substring2, 64)
			coord := []float64{f1, f2}
			mat = append(mat, coord)
		}
	}
	file.Close()
	fmt.Printf("%s\n", polyline.EncodeCoords(mat))
}
