package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("enter how many points")
	fmt.Scan(&n)
	var loce []float64
	loce = make([]float64, n, n)
	var loca []float64
	loca = make([]float64, n, n)
	var dur []float64
	dur = make([]float64, n, n)
	fmt.Println("enter LOCe values")
	for i := 0; i < n; i++ {
		fmt.Scan(&loce[i])
	}
	fmt.Println("enter LOCa values")
	for i := 0; i < n; i++ {
		fmt.Scan(&loca[i])
	}
	fmt.Println("enter Da values")
	for i := 0; i < n; i++ {
		fmt.Scan(&dur[i])
	}
	// loce := []float64{255, 475, 362, 341, 469, 255, 248, 195, 223, 310, 191, 449, 258, 233, 347, 284, 339, 135, 280, 303}
	// loca := []float64{258, 562, 121, 321, 534, 256, 323, 197, 302, 343, 151, 430, 256, 294, 359, 394, 376, 117, 232, 148}
	// dur := []float64{600, 1440, 300, 780, 1440, 600, 780, 420, 780, 780, 360, 900, 660, 720, 780, 1020, 960, 300, 660, 420}
	r1 := cal_r(loce, loca)
	r2 := cal_r(loce, dur)
	r3 := cal_r(loca, dur)
	fmt.Println("LOCe x LOCa r value is", r1)
	fmt.Println("LOCe x Da r value is", r2)
	fmt.Println("LOCe x Da r value is", r3)
	t_value1 := cal_tvalue(r1)
	t_value2 := cal_tvalue(r2)
	t_value3 := cal_tvalue(r3)
	//fmt.Println(t_value1, t_value2, t_value3)
	sigma1 := 1 - simpsonRule(t_value1, float64(n-2))
	if sigma1 < 0 {
		sigma1 = 0
	}
	sigma2 := 1 - simpsonRule(t_value2, float64(n-2))
	if sigma2 < 0 {
		sigma2 = 0
	}
	sigma3 := 1 - simpsonRule(t_value3, float64(n-2))
	if sigma3 < 0 {
		sigma3 = 0
	}
	fmt.Println("LOCe x LOCa sigma value is", sigma1)
	fmt.Println("LOCe x Da sigma value is", sigma2)
	fmt.Println("LOCe x Da sigma value is", sigma3)

	if math.Pow(r1, 2) > 0.5 && sigma1 < 0.2 {
		fmt.Println("LOCe x LOCa are correlated")
	} else {
		fmt.Println("LOCe x LOCa are not correlated")
	}

	if math.Pow(r2, 2) > 0.5 && sigma2 < 0.2 {
		fmt.Println("LOCe x Da are correlated")
	} else {
		fmt.Println("LOCe x Da are not correlated")
	}
	if math.Pow(r3, 2) > 0.5 && sigma3 < 0.2 {
		fmt.Println("LOCa x Da are correlated")
	} else {
		fmt.Println("LOCa x Da are not correlated")
	}
}

func denominator(arr []float64, mean float64) float64 {
	var sum float64 = 0.0
	var temp float64 = 0.0
	for _, v := range arr {
		temp = v - mean
		sum = sum + math.Pow(temp, 2)
	}
	return sum
}

func numerator(arr []float64, arr1 []float64, arr_mean float64, arr1_mean float64) float64 {
	numerator := 0.0
	for i := 0; i < len(arr); i++ {
		x := arr[i] - arr_mean
		y := arr1[i] - arr1_mean
		product := x * y
		numerator = numerator + product
	}
	return numerator
}

func mean(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum = sum + v
	}
	return float64(sum) / float64(len(arr))
}

func standardDeviation(arr []float64, mean float64) float64 {
	var sum float64 = 0.0
	var temp float64 = 0.0
	for _, v := range arr {
		temp = v - mean
		sum = sum + math.Pow(temp, 2)
	}
	return math.Sqrt(sum / float64(len(arr)-1.0))
}

func findBeta(numerator float64, denominator float64, x_mean float64, y_mean float64) float64 {
	beta1 := numerator / denominator
	if math.IsNaN(beta1) {
		error := errors.New("the number is found out to be NAN please try with someother inputs")
		fmt.Println(error)
	}
	return beta1
}

func cal_r(x []float64, y []float64) float64 {
	x_mean := mean(x)
	y_mean := mean(y)
	numerator := numerator(x, y, x_mean, y_mean)
	denominator := denominator(x, x_mean)
	beta1 := findBeta(numerator, denominator, x_mean, y_mean)
	r := beta1 * ((standardDeviation(x, x_mean)) / standardDeviation(y, y_mean))
	return r
}

func cal_tvalue(r float64) float64 {
	t_value := (math.Abs(r) * math.Sqrt(20-2)) / math.Sqrt(1-math.Pow(r, 2))
	return t_value
}

func gamma(x float64) float64 {
	if x == 1.0 {
		return 1.0
	} else if x == 0.5 {
		return math.Sqrt(math.Pi)
	} else {
		return (x - 1) * gamma(x-1)
	}
}

func calFunction(dof float64, x float64) float64 {
	a := (gamma((dof+1)/2) / (gamma(dof/2) * math.Sqrt(dof*math.Pi)))
	b := 1 / (math.Pow(1+(x*x/dof), (dof+1)/2))
	return a * b
}

func simpsonRule(t float64, dof float64) float64 {
	a := -t
	b := t
	error := 1.0
	N := 4
	oldsum := 0.0
	sum := 0.0
	for math.Abs(error) > 0.00001 {
		delta_x := (b - a) / float64(N)
		h := delta_x / 3
		for i := 0; i < N; i++ {
			if i == 0 {
				sum = sum + calFunction(dof, a+float64(i)*delta_x)
			} else if i == N {
				sum = sum + calFunction(dof, a+float64(i)*delta_x)
			} else if i%2 == 0 {
				sum = sum + 2*calFunction(dof, a+float64(i)*delta_x)
			} else if i%2 == 1 {
				sum = sum + 4*calFunction(dof, a+float64(i)*delta_x)
			}
		}
		sum = h * sum
		N = N * 2
		error = (oldsum - sum) / sum
		oldsum = sum
	}
	return sum
}
