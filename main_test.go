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
	res := false
	for _,subelem := range subslice {
		res = false
		for _,elem := range slice {
			if elem == subelem {
				res = true
				break
			}
		}
		if !res {
			return false
		}
	}
	return true
}

func TestBinarySearchSimple(t *testing.T) {
	slice := make([]int,100000)
	for i:=0;i<len(slice);i++ {
		slice[i] = i
	}
	for _,elem := range slice {
		if !binarySearch(slice,elem) {
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
	subslice := []int{3,5,7}
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
		if i < len(subslice)/2 {
			subslice[i] = slice[i]
			continue
		}
		subslice[i] = slice[len(slice)-i]
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
		if i < len(subslice)/2 {
			subslice[i] = slice[i]
			continue
		}
		subslice[i] = slice[len(slice)-i]
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