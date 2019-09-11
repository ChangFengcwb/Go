package main

import (
	"BasicDataStructure/Search"
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
	//a := []int{123, 12, 312, 31, 345, 456, 21, 4, 23523, 4, 234}
	//b := &Sort.QuickSort{}
	//b.Sort(a)
	//fmt.Println(a)
	a := Search.NewBSTSymbolTable()
	a.Put(1, "男")
	a.Put(2, "男")
	a.Put(3, "男")
	a.Put(4, "男")
	a.Put(5, "男")
	a.Put(6, "男")
	a.Put(7, "男")
	a.Put(8, "男")
	a.Put(9, "男")
	a.Put(10, "男")
	a.Put(11, "男")
	a.Put(12, "男")
	a.Put(13, "男")
	a.Put(14, "男")
	a.PrintMiddle()
}
