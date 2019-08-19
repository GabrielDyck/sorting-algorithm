package main

import (
	"../main/utils"
	"fmt"
	"time"
)

/* Bubble sort
[5*,3,2,4,1] -> [3,5*,2,4,1] -> [3,2,5*,4,1] -> [3,2,4,5*,1] -> [3,2,4,1,5*] ->
[2*,3,4,1,5] -> [2,3*,4,1,5] -> [2,3,4*,1,5] -> [2,3,1,4*,5] -> [2,3,1,4,5*] ->
[2*,3,1,4,5] -> [2,3*,1,4,5] -> [2,1,3*,4,5] -> [2,1,3,4*,5] -> [2,1,3,4,5*] ->
[2*,1,3,4,5] -> [1,2*,3,4,5] -> [1,2,3*,4,5] -> [1,2,3,4*,5] -> [1,2,3,4,5*] ->
[1*,2,3,4,5] -> [1,2*,3,4,5] -> [1,2,3*,4,5] -> [1,2,3,4*,5] -> [1,2,3,4,5*] ->
*/


var comparitions =0
var fullScans = 1
var startTime= time.Now()
func main () {
	values:= utils.GenerateUnsortedValues(1000)
	//values:= []int{5,3,2,4,1}
	fmt.Println("Init state of values: ", values)
	unsorted:= true
	for unsorted {
		fmt.Println("Iteration number: ",fullScans, ". State: ", values)
		sorted:= true
		for firstIndex := 0 ; firstIndex < len(values); firstIndex++{
			fmt.Println("Using index: ", firstIndex, ". State: ",values)
			if firstIndex + 1 >= len(values){continue}
			sorted= !compareAndSort(values,firstIndex) && sorted
		}
		unsorted=!sorted
		fullScans++

	}
	endTime:= time.Now()
	fmt.Println("Comparitions: ", comparitions)
	fmt.Println("FullScans: ", fullScans)
	fmt.Println("Elapsed time: ", endTime.Sub(startTime).Seconds())
}


func compareAndSort(values []int, index int)(hasToSort bool){
	comparitions++
	if values[index] > values[index+1]{
		fmt.Println("Sorting ",values[index] , " and ", values[index+1] )
		temp:= values[index+1]
		values[index+1]=values[index]
		values[index]=temp
		hasToSort=true
	}
	return
}