package main

import (
	"fmt"
	"testing"
	"time"
)

func duration(duration int64, realisation string,t *testing.T){
	s := fmt.Sprintf("%s time(ns): %d\n", realisation, duration)
	t.Log(s)
}

//O(M*N)->O(N) [M<<N]
func isIncludeNative(slice, subslice []int) bool {
	if len(subslice) == 0 {
		return true
	} else if (len(slice) == 0) || (subslice[0] < slice[0]) ||
		(subslice[len(subslice)-1] > slice[len(slice)-1]) {
		return false
	}

	subelem := subslice[0]
	index := -1
	for i,elem := range slice {
		if elem == subelem {
			index = i
		}
	}
	if index == -1 ||
		(index + len(subslice) > len(slice)){
		return false
	}
	for i:=1;i<len(subslice);i++ {
		if subslice[i] != slice[index+i] {
			return false
		}
	}
	return true
}

func TestBinarySearchSimple(t *testing.T) {
	slice := make([]int,10)
	for i:=0;i<len(slice);i++ {
		slice[i] = i
	}
	for i,elem := range slice {
		if binarySearch(slice,elem) != i {
			t.Error("For slice: ", slice,
				" binarySearch() have not element ", elem)
		}
	}
	slice = append(slice,10)
	for i,elem := range slice {
		if binarySearch(slice,elem) != i {
			t.Error("For slice: ", slice,
				" binarySearch() have not element ", elem)
		}
	}
}
func comparator(slice,subslice []int, t *testing.T) {
	resisInclude := isInclude(slice,subslice)
	resisIncludeNative := isIncludeNative(slice,subslice)
	if resisInclude != resisIncludeNative {
		t.Error("For slice: ", slice,
			"subslice: ", subslice,
			"expected:", resisIncludeNative,
			"got:",resisInclude)
	}
}

func TestSimpleAnswers(t *testing.T){
	slice := []int{1,2,3,5,7,9,11}
	subslice := []int{3,5,9}
	comparator(slice,subslice,t)

	subslice[2] = 7
	comparator(slice,subslice,t)

	subslice[0] = 4
	comparator(slice,subslice,t)

	subslice = []int{}
	comparator(slice,subslice,t)
}

func TestHardAnswers(t *testing.T) {
	slice := make([]int,100000000)
	for i:=0;i<len(slice);i++ {
		slice[i] = i*2
	}
	subslice := make([]int,100)
	for i:=0;i<len(subslice);i++ {
		subslice[i] = slice[i]
	}
	comparator(slice,subslice,t)

	subslice[len(subslice)-1]++
	comparator(slice,subslice,t)
}

func TestTiming(t *testing.T){
	slice := make([]int,100000000)
	for i:=0;i<len(slice);i++ {
		slice[i] = i*2
	}
	subslice := make([]int,100)
	for i:=0;i<len(subslice);i++ {
		subslice[i] = slice[i]
	}

	start := time.Now()
	isIncludeNative(slice,subslice)
	durNative := time.Since(start).Nanoseconds()
	duration(durNative,"isIncludeNative",t)

	start = time.Now()
	isInclude(slice,subslice)
	dur := time.Since(start).Nanoseconds()
	duration(dur,"isInclude",t)

	if dur >= durNative {
		t.Error("Time limit error")
	}
}