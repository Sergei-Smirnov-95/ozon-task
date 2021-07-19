package main

import (
	"fmt"
)

//O(logN)
func binarySearch(slice []int, subelem int) int {
	sliceLen := len(slice)

	if sliceLen == 2 {
		switch subelem {
		case slice[0]:
			return 0
		case slice[1]:
			return 1
		default:
			return -1
		}
	} else if sliceLen == 1 {
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
	} else if len(subslice) != 0 && len(slice) == 0 {
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
	slice := make([]int,1000)
	for i:=0;i<len(slice);i++ {
		slice[i] = i*2
	}
	var subslice =make([]int,10)
	for i:=0;i<len(subslice);i++ {
		if i < len(subslice)-1 {
			subslice[i] = slice[i]
			continue
		}
		subslice[i] = slice[len(slice)-1]
	}
	fmt.Printf("my isInclude: %v\n",isInclude(slice,subslice))
}
