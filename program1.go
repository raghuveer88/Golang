package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hom many numbers do you want to enter(Range should be from 1 to 200)")
	//call with manoj
	var leng int
	fmt.Scan(&leng)
	if leng < 0 || leng > 200 {
		fmt.Println("sorry the length is not in the range")
	}
	inptArray := make([]float64, leng)
	fmt.Println("enter the values")
	for i := 0; i < leng; i++ {
		fmt.Scan(&inptArray[i])
	}

	// fmt.Println(inptArray)
	sortedArray := make([]float64, leng)
	sortedArray = sort(inptArray, leng)
	//fmt.Println(sortedArray)

	finalArray := sArray{
		array: sortedArray,
		len:   leng,
	}
	var opArray statistics = &finalArray
	mean := opArray.mean()
	fmt.Println("Mean is ", mean)

	standardDeviation := opArray.standardDeviation(mean)
	fmt.Println("standard deviation is ", standardDeviation)

	median := opArray.median()
	fmt.Println("Median is ", median)

}

func sort(sortedArray []float64, len int) []float64 {
	for i := 1; i < len; i++ {
		temp := sortedArray[i]
		j := i - 1
		for ; j >= 0 && sortedArray[j] > temp; j = j - 1 {
			sortedArray[j+1] = sortedArray[j]
		}
		sortedArray[j+1] = temp
	}
	return sortedArray
}

type statistics interface {
	mean() float64
	median() float64
	standardDeviation(float64) float64
}

type sArray struct {
	array []float64
	len   int
}

func (arr *sArray) mean() float64 {
	sum := 0.0
	for _, v := range (*arr).array {
		sum += v
	}
	return float64(sum) / float64((*arr).len)

}

func (arr *sArray) median() float64 {
	mid := (*arr).len / 2
	if (*arr).len%2 == 0 {
		return ((*arr).array[mid-1] + (*arr).array[mid]) / 2.0
	}
	mid = int(mid)
	return (*arr).array[mid]
}

func (arr *sArray) standardDeviation(mean float64) float64 {
	var sum float64 = 0.0
	var temp float64 = 0.0
	for _, v := range (*arr).array {
		temp = v - mean
		sum = sum + math.Pow(temp, 2)
	}
	//fmt.Println(sum)

	return math.Sqrt(sum / float64((*arr).len-1.0))

}
