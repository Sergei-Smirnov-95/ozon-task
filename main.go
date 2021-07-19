package main

import (
	"fmt"
)

//O(logN)
func binarySearch(slice []int, subelem int) bool {
	sliceLen := len(slice)

	if sliceLen == 2 {
		return subelem == slice[0] || subelem == slice[1]
	} else if sliceLen == 1 {
		return subelem == slice[0]
	}

	medIndex := sliceLen/2
	elem := slice[medIndex]

	 if elem == subelem {
		return true
	 } else if elem > subelem {
	 	return binarySearch(slice[:medIndex],subelem)
	 } else {
		return binarySearch(slice[medIndex+1:],subelem)
	 }
}


//O(M*logN)->O(logN) [M<<N]
func isInclude(slice, subslice []int) bool {
	for _,subelem := range subslice {
		if !binarySearch(slice,subelem) {
			return false
		}
	}
	return true
}



func main() {
	slice := make([]int,1000)
	for i:=0;i<len(slice);i++ {
		slice[i] = i*2
	}
	subslice := make([]int,10)
	for i:=0;i<len(subslice);i++ {
		if i < len(subslice)/2 {
			subslice[i] = slice[i]
			continue
		}
		subslice[i] = slice[len(slice)-i]
	}
	fmt.Printf("my isInclude: %v\n",isInclude(slice,subslice))
}
