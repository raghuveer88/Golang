package main

import (
	"fmt"
	"math"
)

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

func simpsonRule(t float64, dof float64, tails float64) float64 {
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
	if tails == 1 {
		left := 1 - sum
		left_l := left / 2
		sum = left_l + sum
	}
	return sum

}

func main() {
	var t float64
	var dof float64
	var tails float64
	for true {
		fmt.Println("Please enter t value")
		fmt.Scan(&t)
		if t >= 0 {
			break
		} else {
			fmt.Println("t value should be greater than 0")
		}
	}
	for true {
		fmt.Println("Enter Degree of Freedom value")
		fmt.Scan(&dof)
		if dof >= 0 {
			break
		} else {
			fmt.Println("dof should be >=0")
		}
	}
	for true {
		fmt.Println("Enter Tails value 1 or 2")
		fmt.Scan(&tails)
		if tails == 2 || tails == 1 {
			break
		} else {
			fmt.Println("(tails value should be 1 or 2)")
		}
	}
	fmt.Println(simpsonRule(t, dof, tails))
}
