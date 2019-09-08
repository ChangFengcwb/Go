package main

import (
	"BasicDataStructure/Sort"
	"fmt"
)

func main() {
	//file, err := os.Open("1Mints.txt")
	//if err != nil {
	//	panic(err)
	//}
	//a := make([]int, 0)
	//defer file.Close()
	//br := bufio.NewReader(file)
	//for {
	//	b, _, c := br.ReadLine()
	//	if c == io.EOF {
	//		break
	//	}
	//
	//	d, err := strconv.Atoi(strings.TrimSpace(string(b)))
	//	if err != nil {
	//		panic(err)
	//	}
	//	a = append(a, d)
	//}
	a := []int{123, 12, 312, 31, 345, 456, 21, 4, 23523, 4, 234}
	b := &Sort.QuickSort{}
	b.Sort(a)
	fmt.Println(a)
}
