package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	loce := []float64{255, 475, 362, 341, 469, 255, 248, 195, 223, 310, 191, 449, 258, 233, 347, 284, 339, 135, 280, 303}
	dur := []float64{600, 1440, 300, 780, 1440, 600, 780, 420, 780, 780, 360, 900, 660, 720, 780, 1020, 960, 300, 660, 420}
	//fmt.Println(len(loce))
	//fmt.Println(len(dur))
	loce_mean := mean(loce)
	dur_mean := mean(dur)
	numerator := numerator(loce, dur, loce_mean, dur_mean)
	denominator := denominator(loce, loce_mean)
	beta0, beta1 := findBeta(numerator, denominator, loce_mean, dur_mean)
	//fmt.Println("beta 0 is ", beta0)
	//fmt.Println("beta 1 is ", beta1)
	var point float64
	fmt.Print("enter the raw estimated LOC - ")
	fmt.Scanf("%f", &point)
	for true {
		if point <= 0 {
			fmt.Println("projected value is Invalid")
			fmt.Println("70% UPI -  Invalid")
			fmt.Println("70% LPI -  Invalid")
			fmt.Println("90% UPI -  Invalid")
			fmt.Println("90% LPI -  Invalid")
			fmt.Print("enter the raw estimated LOC - ")
			fmt.Scanf("%f", &point)
		} else {
			break
		}
	}
	y_predict := predict(float64(point), beta0, beta1)
	fmt.Println("projected value is ", y_predict)

	t_value := calculate_t(0.7, 18)
	//fmt.Println("t value is ", calculate_t(0.7, 18))
	s_dev := std(loce, dur, beta1, beta0)
	variance := variance(loce, loce_mean, point)
	size_range := t_value * s_dev * variance
	//fmt.Println("the range size is ", size_range)
	fmt.Println("70% UPI -  ", y_predict+size_range)
	fmt.Println("70% LPI -  ", y_predict-size_range)

	t_value = calculate_t(0.9, 18)
	size_range = t_value * s_dev * variance
	fmt.Println("90% UPI -  ", y_predict+size_range)
	fmt.Println("90% LPI -  ", y_predict-size_range)
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

func findBeta(numerator float64, denominator float64, x_mean float64, y_mean float64) (float64, float64) {
	beta1 := numerator / denominator
	if math.IsNaN(beta1) {
		error := errors.New("the number is found out to be NAN please try with someother inputs")
		fmt.Println(error)

	}
	//fmt.Println("beta1 value is ", beta1)

	beta0 := y_mean - (beta1 * x_mean)
	//fmt.Println("beta0 value is ", beta0)
	return beta0, beta1
}

func mean(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum = sum + v
	}
	return float64(sum) / float64(len(arr))
}

func predict(x float64, beta0 float64, beta1 float64) float64 {
	y := beta1*x + beta0
	return y
}

func std(x []float64, y []float64, beta1 float64, beta0 float64) float64 {
	var sum float64 = 0.0
	var temp float64 = 0.0
	for i := 0; i < len(x); i++ {
		y_hat := beta1*x[i] + beta0
		temp = y[i] - y_hat
		sum = sum + math.Pow(temp, 2)
	}
	return math.Sqrt(sum / float64(len(x)-2.0))

}

func variance(data []float64, mean float64, e float64) float64 {
	num := math.Pow(e-mean, 2)
	var temp float64 = 0.0
	var sum float64 = 0.0
	for i := 0; i < len(data); i++ {
		temp = data[i] - mean
		sum = sum + math.Pow(temp, 2)
	}
	return math.Sqrt(1.0 + float64(1/len(data)) + (num / sum))
}

func calculate_t(prediction_interval float64, dof float64) float64 {
	//interval_num := 64
	prob_least := (prediction_interval / 2) - 0.00001
	prob_high := (prediction_interval / 2) + 0.00001
	value_t := 1.0
	prob := simpsonRule(0, value_t, dof)
	for prob <= prob_least || prob >= prob_high {
		if prob >= prob_high {
			value_t = value_t / 2
			prob = simpsonRule(0, value_t, dof)
		} else {
			value_t += value_t / 2
			prob = simpsonRule(0, value_t, dof)
		}

	}
	return value_t
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

func simpsonRule(a float64, t float64, dof float64) float64 {
	b := t
	error := 1.0
	N := 4
	oldsum := 0.0
	sum := 0.0
	for math.Abs(error) > 0.0001 {
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
