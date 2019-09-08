package main

import (
	"BasicDataStructure/Sort"
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var A []int

func init() {
	file, err := os.Open("1Mints.txt")
	if err != nil {
		panic(err)
	}
	A = make([]int, 0)
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		d, err := strconv.Atoi(strings.TrimSpace(string(a)))
		if err != nil {
			panic(err)
		}
		A = append(A, d)
	}

}

//func TestSelection(t *testing.T) {
//	a := make([]int, len(A))
//	copy(a, A)
//	b := &Sort.Selection{}
//	b.Sort(a)
//}
//func TestInsert(t *testing.T) {
//	a := make([]int, len(A))
//	copy(a, A)
//	b := &Sort.Insert{}
//	b.Sort(a)
//}
//func TestShell(t *testing.T) {
//	a := make([]int, len(A))
//	copy(a, A)
//	b := &Sort.Shell{}
//	b.Sort(a)
//}

func TestMerge(t *testing.T) {
	a := make([]int, len(A))
	copy(a, A)
	b := &Sort.Merge{}
	b.Sort(a)
}

func TestInMerge(t *testing.T) {
	a := make([]int, len(A))
	copy(a, A)
	b := &Sort.InMerge{}
	b.Sort(a)
}

func TestQuickSort(t *testing.T) {
	a := make([]int, len(A))
	copy(a, A)
	b := &Sort.QuickSort{}
	b.Sort(a)
}
func TestThreeQuickSort(t *testing.T) {
	a := make([]int, len(A))
	copy(a, A)
	b := &Sort.ThreeQuickSort{}
	b.Sort(a)
}

func TestHeapSort(t *testing.T) {
	a := make([]int, len(A))
	copy(a, A)
	b := &Sort.HeapSort{}
	b.Sort(a)
}

func TestSystemSort(t *testing.T) {
	a := make([]int, len(A))
	copy(a, A)
	sort.Ints(a)
}
