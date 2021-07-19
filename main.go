package main

import (
	"fmt"
)

//O(logN)
func binarySearch(slice []int, subelem int) int {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return -1
	}
	if sliceLen == 1 {
		if subelem == slice[0] {
			return 0
		}
		return -1
	}

	medIndex := sliceLen/2
	elem := slice[medIndex]

	 if elem == subelem {
		return medIndex
	 } else if elem > subelem {
	 	return binarySearch(slice[:medIndex],subelem)
	 } else {
		return medIndex+1 +binarySearch(slice[medIndex+1:],subelem)
	 }
}


//O(M+logN)->O(logN) [M<<N]
func isInclude(slice, subslice []int) bool {
	if len(subslice) == 0 {
		return true
	} else if (len(slice) == 0) || (subslice[0] < slice[0]) ||
	(subslice[len(subslice)-1] > slice[len(slice)-1]) {
		return false
	}

	firstIndex := binarySearch(slice,subslice[0])
	if firstIndex == -1 ||
		(firstIndex + len(subslice) > len(slice)){
		return false
	}

	for i:=1;i<len(subslice);i++ {
		if subslice[i] != slice[firstIndex+i] {
			return false
		}
	}
	return true
}



func main() {
	slice := []int{1,2,3,5,7,9,11}
	subslice := []int{3,5,7}
	fmt.Printf("my isInclude: %v\n",isInclude(slice,subslice))
}
