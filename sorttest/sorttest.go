package main

import (
"sort" 
"fmt"
)

func main(){
	arr := []int{9,8,7,6,3,2,1}
	sort.Ints(arr)
	fmt.Printf("%v\n", arr)
	result := sort.SearchInts(arr, 5)
	fmt.Printf("%v\n", result)
}