package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	fmt.Println("enter length of values (note - length of x & y should be equal)")
	var x_leng int
	fmt.Scan(&x_leng)
	x_inputs := make([]float64, x_leng)
	fmt.Println("enter the values of x")
	for i := 0; i < x_leng; i++ {
		fmt.Scan(&x_inputs[i])
	}

	x_finalArray := sArray{
		array: x_inputs,
		len:   x_leng,
	}

	var y_leng int = x_leng	
	y_inputs := make([]float64, y_leng)
	fmt.Println("enter the values of y")
	for i := 0; i < y_leng; i++ {
		fmt.Scan(&y_inputs[i])
	}

	y_finalArray := sArray{
		array: y_inputs,
		len:   y_leng,
	}

	var x_Array statistics = &x_finalArray
	var y_Array statistics = &y_finalArray
	x_mean := x_Array.mean()
	//fmt.Println("x value Mean is ", x_mean)

	y_mean := y_Array.mean()
	//fmt.Println("y value Mean is ", y_mean)

	denominator := x_Array.denominator(x_mean)
	//fmt.Println("denominator is ", denominator)

	numerator := 0.0
	for i:=0;i<x_leng;i++{
		x := x_inputs[i] - x_mean
		y := y_inputs[i] - y_mean
		product := x*y
		numerator = numerator + product
	}

	findBeta(numerator,denominator,x_mean,y_mean)

}

type statistics interface {
	mean() float64
	denominator(float64) float64
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



func (arr *sArray) denominator(mean float64) float64 {
	var sum float64 = 0.0
	var temp float64 = 0.0
	for _, v := range (*arr).array {
		temp = v - mean
		sum = sum + math.Pow(temp, 2)
	}

	return sum

}

func findBeta(numerator float64,denominator float64,x_mean float64,y_mean float64){
	beta1 := numerator/denominator
	if math.IsNaN(beta1){
		error := errors.New("the number is found out to be NAN please try with someother inputs")
		fmt.Println(error)
		return 
	}
	fmt.Println("beta1 value is ",beta1)

	beta0 := y_mean - (beta1*x_mean)
	fmt.Println("beta0 value is ",beta0)
}

