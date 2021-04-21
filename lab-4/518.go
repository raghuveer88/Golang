package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// Open the file
	csvfile1, err := os.Open("pastproxies.csv")
	csvfile2, err := os.Open("actual_estimated.csv")
	csvfile3, err := os.Open("objectLOC.csv")
	csvfile4, err := os.Open("newobjects.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	file1 := csv.NewReader(csvfile1)
	file2 := csv.NewReader(csvfile2)
	file3 := csv.NewReader(csvfile3)
	file4 := csv.NewReader(csvfile4)
	//file1 := csv.NewReader(bufio.NewReader(csvfile))
	pastproxies_1 := []string{}
	pastproxies_2 := []string{}
	pastproxies_3 := []string{}
	pastproxies_4 := []string{}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := file1.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record[0]) > 0 {
			pastproxies_1 = append(pastproxies_1, record[0])
		}
		if len(record[1]) > 0 {
			pastproxies_2 = append(pastproxies_2, record[1])
		}
		if len(record[2]) > 0 {
			pastproxies_3 = append(pastproxies_3, record[2])
		}
		if len(record[3]) > 0 {
			pastproxies_4 = append(pastproxies_4, record[3])
		} else {

			break
		}

	}
	// fmt.Println(pastproxies_1)
	// fmt.Println(pastproxies_2)
	// fmt.Println(pastproxies_3)
	// fmt.Println(pastproxies_4)

	actual_estimated_1 := []string{}
	actual_estimated_2 := []string{}
	actual_estimated_3 := []string{}
	actual_estimated_4 := []string{}
	actual_estimated_5 := []string{}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := file2.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record[0]) > 0 {
			actual_estimated_1 = append(actual_estimated_1, record[0])
		}
		if len(record[1]) > 0 {
			actual_estimated_2 = append(actual_estimated_2, record[1])
		}
		if len(record[2]) > 0 {
			actual_estimated_3 = append(actual_estimated_3, record[2])
		}
		if len(record[3]) > 0 {
			actual_estimated_4 = append(actual_estimated_4, record[3])
		}
		if len(record[4]) > 0 {
			actual_estimated_5 = append(actual_estimated_5, record[4])
		} else {

			break
		}

	}
	// fmt.Println(actual_estimated_1)
	// fmt.Println(actual_estimated_2)
	// fmt.Println(actual_estimated_3)
	// fmt.Println(actual_estimated_4)
	// fmt.Println(actual_estimated_5)

	objectLOC_1 := []string{}
	objectLOC_2 := []string{}
	objectLOC_3 := []string{}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := file3.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record[0]) > 0 {
			objectLOC_1 = append(objectLOC_1, record[0])
		}
		if len(record[1]) > 0 {
			objectLOC_2 = append(objectLOC_2, record[1])
		}
		if len(record[2]) > 0 {
			objectLOC_3 = append(objectLOC_3, record[2])
		} else {

			break
		}

	}
	// fmt.Println(objectLOC_1)
	// fmt.Println(objectLOC_2)
	// fmt.Println(objectLOC_3)

	newobjects_1 := []string{}
	newobjects_2 := []string{}
	newobjects_3 := []string{}
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := file4.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record[0]) > 0 {
			newobjects_1 = append(newobjects_1, record[0])
		}
		if len(record[1]) > 0 {
			newobjects_2 = append(newobjects_2, record[1])
		}
		if len(record[2]) > 0 {
			newobjects_3 = append(newobjects_3, record[2])
		} else {

			break
		}

	}
	// fmt.Println(newobjects_1)
	// fmt.Println(newobjects_2)
	// fmt.Println(newobjects_3)

	pastproxies_5 := []float64{} // LOC per method
	pastproxies_5 = getLocPerMethod(pastproxies_2, pastproxies_3)
	//fmt.Println(pastproxies_5)

	cal_mean, data_mean, io_mean := mean1(pastproxies_5, pastproxies_4)
	//fmt.Println(cal_mean, data_mean, io_mean)

	cal_std, data_std, io_std := standardDeviation(pastproxies_5, pastproxies_4, cal_mean, data_mean, io_mean)
	//fmt.Println(cal_std, data_std, io_std)

	cal_categorytable := [5][3]float64{}
	cal_categorytable = sizeCategoryTable(cal_mean, cal_std)

	data_categorytable := [5][3]float64{}
	data_categorytable = sizeCategoryTable(data_mean, data_std)

	io_categorytable := [5][3]float64{}
	io_categorytable = sizeCategoryTable(io_mean, io_std)

	fmt.Println("the mid point and ranges for calculation ,Data and I/O are")
	fmt.Println("For calculation")
	fmt.Println(cal_categorytable)
	fmt.Println("For data")
	fmt.Println(data_categorytable)
	fmt.Println("For io")
	fmt.Println(io_categorytable)

	ba := similarSearch(objectLOC_1, objectLOC_2, objectLOC_3, pastproxies_1, pastproxies_4, pastproxies_5, cal_categorytable, data_categorytable, io_categorytable)
	no := similarSearch(newobjects_1, newobjects_2, newobjects_3, pastproxies_1, pastproxies_4, pastproxies_5, cal_categorytable, data_categorytable, io_categorytable)
	fmt.Println("---------------------------------------")
	fmt.Println("BA total LOC is ", ba)
	fmt.Println("NO total LOC is ", no)

	base_size := 369.0
	loc_deleted := 14.0
	loc_modified := 14.0

	e := ba + no + float64(loc_modified)
	fmt.Println("Estimated Object LOC is ", e)

	x_inputs := []float64{}
	y_inputs := []float64{}
	z_inputs := []float64{}
	// k_inputs := []float64{}
	for i := 0; i < len(actual_estimated_2); i++ {
		temp, _ := strconv.ParseFloat(actual_estimated_2[i], 8)
		temp2, _ := strconv.ParseFloat(actual_estimated_3[i], 8)
		temp3, _ := strconv.ParseFloat(actual_estimated_4[i], 8)
		//temp4, _ := strconv.ParseFloat(actual_estimated_5[i], 8)
		x_inputs = append(x_inputs, temp)
		y_inputs = append(y_inputs, temp2)
		z_inputs = append(z_inputs, temp3)
		//fmt.Println(z_inputs)
		// k_inputs = append(y_inputs, temp4)
	}

	x_finalArray := sArray{
		array: x_inputs,
		len:   len(x_inputs),
	}

	y_finalArray := sArray{
		array: y_inputs,
		len:   len(y_inputs),
	}
	z_finalArray := sArray{
		array: z_inputs,
		len:   len(z_inputs),
	}
	// k_finalArray := sArray{
	// 	array: k_inputs,
	// 	len:   len(k_inputs),
	// }
	var x_Array statistics = &x_finalArray
	var y_Array statistics = &y_finalArray
	var z_Array statistics = &z_finalArray
	//var k_Array statistics = &k_finalArray
	x_mean := x_Array.mean()
	y_mean := y_Array.mean()
	//fmt.Println("y value Mean is ", y_mean)

	denominator := x_Array.denominator(x_mean)
	//fmt.Println("denominator is ", denominator)

	numerator := 0.0
	for i := 0; i < len(x_inputs); i++ {
		x := x_inputs[i] - x_mean
		y := y_inputs[i] - y_mean
		product := x * y
		numerator = numerator + product
	}

	beta0, beta1 := findBeta(numerator, denominator, x_mean, y_mean)
	y := beta1*e + beta0
	y = math.Round(y)
	//y = math.Floor(y*10) / 10
	fmt.Println("Estimated New and Changed LOC(N)", y)

	fmt.Println("Estimated Total LOC(T) is ", y+base_size-loc_deleted-loc_modified)

	t_value := 1.067169516
	s_dev := std(x_inputs, y_inputs, beta1, beta0)
	//fmt.Println(s_dev)

	variance := variance(x_inputs, x_mean, e)
	//fmt.Println(variance)

	size_range := t_value * s_dev * variance
	fmt.Println("the range size is ", size_range)

	fmt.Println("the upper range size limit is ", y+size_range)
	fmt.Println("the lower range size limit is ", y-size_range)
	fmt.Println("-------------------------------------------")
	sum1 := 0.0
	sum2 := 0.0
	for i := 0; i < len(actual_estimated_2); i++ {
		sum1 = sum1 + x_inputs[i]
		sum2 = sum2 + z_inputs[i]
	}
	fmt.Println("productivity is ", (sum1/sum2)*60)

	z_mean := z_Array.mean()
	denominator1 := x_Array.denominator(x_mean)
	//fmt.Println("denominator is ", denominator)

	numerator1 := 0.0
	for i := 0; i < len(x_inputs); i++ {
		x := x_inputs[i] - x_mean
		z := z_inputs[i] - z_mean
		product := x * z
		numerator1 = numerator1 + product
	}

	b0, b1 := findBeta(numerator1, denominator1, x_mean, z_mean)
	y1 := b1*e + b0
	y1 = math.Round(y1)
	fmt.Println("Estimated project duration", y1)

	s_dev1 := std(x_inputs, z_inputs, b1, b0)

	//fmt.Println(s_dev1)
	duration_range := t_value * s_dev1 * variance

	fmt.Println("the duration range is ", duration_range)

	fmt.Println("the upper range suration limit is ", y1+duration_range)
	fmt.Println("the lower range suration limit is ", y1-duration_range)

}

func getLocPerMethod(lines []string, methods []string) []float64 {
	loc_per_method := []float64{}
	a := ""
	b := ""
	var div float64 = 0.0
	for i := 0; i < len(lines); i++ {
		a = lines[i]
		b = methods[i]
		//fmt.Println(a)
		temp, _ := strconv.ParseFloat(a, 8)
		temp2, _ := strconv.ParseFloat(b, 8)
		//fmt.Println(temp)
		div = float64(temp) / float64(temp2)
		loc_per_method = append(loc_per_method, div)

	}
	return loc_per_method
}

func mean1(loc_per_method []float64, pastproxies_4 []string) (float64, float64, float64) {
	cal_mean := 0.0
	data_mean := 0.0
	io_mean := 0.0
	cal_total := 0
	data_total := 0
	io_total := 0
	cal_sum := 0.0
	data_sum := 0.0
	io_sum := 0.0
	for i := 0; i < len(pastproxies_4); i++ {
		if pastproxies_4[i] == "Calculation" {
			cal_total = cal_total + 1
			//fmt.Println(math.Log10(loc_per_method[i]))
			cal_sum = cal_sum + math.Log10(loc_per_method[i])
		}
		if pastproxies_4[i] == "Data" {
			data_total = data_total + 1
			data_sum = data_sum + math.Log10(loc_per_method[i])
		}
		if pastproxies_4[i] == "I/O" {
			io_total = io_total + 1
			io_sum = io_sum + math.Log10(loc_per_method[i])
		}

	}
	cal_mean = cal_sum / float64(cal_total)
	data_mean = data_sum / float64(data_total)
	io_mean = io_sum / float64(io_total)
	return cal_mean, data_mean, io_mean
}

func standardDeviation(loc_per_method []float64, pastproxies_4 []string, cal_mean float64, data_mean float64, io_mean float64) (float64, float64, float64) {
	cal_temp := 0.0
	data_temp := 0.0
	io_temp := 0.0
	cal_sum := 0.0
	data_sum := 0.0
	io_sum := 0.0
	cal_total := 0
	data_total := 0
	io_total := 0
	for i := 0; i < len(pastproxies_4); i++ {
		if pastproxies_4[i] == "Calculation" {
			cal_temp = math.Log10(loc_per_method[i]) - cal_mean
			cal_sum = cal_sum + math.Pow(cal_temp, 2)
			cal_total = cal_total + 1
		}
		if pastproxies_4[i] == "Data" {
			data_temp = math.Log10(loc_per_method[i]) - data_mean
			data_sum = data_sum + math.Pow(data_temp, 2)
			data_total = data_total + 1
		}
		if pastproxies_4[i] == "I/O" {
			io_temp = math.Log10(loc_per_method[i]) - io_mean
			io_sum = io_sum + math.Pow(io_temp, 2)
			io_total = io_total + 1
		}
	}
	cal_std := math.Sqrt(cal_sum / float64(cal_total-1))
	data_std := math.Sqrt(data_sum / float64(data_total-1))
	io_std := math.Sqrt(io_sum / float64(io_total-1))
	return cal_std, data_std, io_std
}

func sizeCategoryTable(mean float64, std float64) [5][3]float64 {
	categorytable := [5][3]float64{}
	categorytable[0][0] = mean - 2*std
	categorytable[0][1] = -99999
	categorytable[0][2] = mean - 1.5*std

	categorytable[1][0] = mean - std
	categorytable[1][1] = mean - 1.5*std
	categorytable[1][2] = mean - 0.5*std

	categorytable[2][0] = mean
	categorytable[2][1] = mean - 0.5*std
	categorytable[2][2] = mean + 0.5*std

	categorytable[3][0] = mean + std
	categorytable[3][1] = mean + 0.5*std
	categorytable[3][2] = mean + 1.5*std

	categorytable[4][0] = mean + 2*std
	categorytable[4][1] = mean + 1.5*std
	categorytable[4][2] = 99999

	return categorytable
}

func similarSearch(objectLOC_1 []string, objectLOC_2 []string, objectLOC_3 []string, pastproxies_1 []string, pastproxies_4 []string, pastproxies_5 []float64, cal_categorytable [5][3]float64, data_categorytable [5][3]float64, io_categorytable [5][3]float64) float64 {
	total_lines := 0.0
	//fmt.Println(pastproxies_5)
	for i := 0; i < len(objectLOC_3); i++ {
		for j := 0; j < len(pastproxies_1); j++ {
			if objectLOC_3[i] == pastproxies_1[j] {
				if pastproxies_4[j] == "Calculation" {
					temp := math.Log10(pastproxies_5[j])
					temp1, _ := strconv.ParseFloat(objectLOC_2[i], 8)
					est_loc := temp1 * pastproxies_5[j]
					total_lines = total_lines + est_loc
					for k := 0; k < 5; k++ {
						if temp >= cal_categorytable[k][1] && temp < cal_categorytable[k][2] {
							fmt.Println("-----------------------------------------")
							if k == 0 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is very small")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 1 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is small")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 2 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is medium")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 3 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is large")
								fmt.Println("estimated lines of code is ", est_loc)
								//fmt.Println(temp)
							}
							if k == 4 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is very large")
								fmt.Println("estimated lines of code is ", est_loc)

							}
						}
					}
				}

				if pastproxies_4[j] == "Data" {
					temp := math.Log10(pastproxies_5[j])
					temp1, _ := strconv.ParseFloat(objectLOC_2[i], 8)
					est_loc := temp1 * pastproxies_5[j]
					total_lines = total_lines + est_loc
					for k := 0; k < 5; k++ {
						if temp >= data_categorytable[k][1] && temp < data_categorytable[k][2] {
							fmt.Println("-----------------------------------------")
							if k == 0 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is very small")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 1 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is small")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 2 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is medium")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 3 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is large")
								fmt.Println("estimated lines of code is ", est_loc)
								//fmt.Println(temp)
							}
							if k == 4 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is very large")
								fmt.Println("estimated lines of code is ", est_loc)

							}
						}
					}
				}

				if pastproxies_4[j] == "I/O" {
					temp := math.Log10(pastproxies_5[j])
					//fmt.Println(temp)
					temp1, _ := strconv.ParseFloat(objectLOC_2[i], 8)
					est_loc := temp1 * pastproxies_5[j]
					total_lines = total_lines + est_loc
					for k := 0; k < 5; k++ {
						if temp >= io_categorytable[k][1] && temp < io_categorytable[k][2] {
							fmt.Println("-----------------------------------------")
							if k == 0 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is very small")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 1 {
								//fmt.Println(cal_categorytable[k][1])
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is small")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 2 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is medium")
								fmt.Println("estimated lines of code is ", est_loc)
							}
							if k == 3 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is large")
								fmt.Println("estimated lines of code is ", est_loc)
								//fmt.Println(temp)
							}
							if k == 4 {
								fmt.Println(objectLOC_1[i], "type ", pastproxies_4[j], "and estimated size is very large")
								fmt.Println("estimated lines of code is ", est_loc)

							}
						}
					}
				}
			}
		}
	}
	return total_lines
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

func findBeta(numerator float64, denominator float64, x_mean float64, y_mean float64) (float64, float64) {
	beta1 := numerator / denominator
	if math.IsNaN(beta1) {
		error := errors.New("the number is found out to be NAN please try with someother inputs")
		fmt.Println(error)

	}
	fmt.Println("beta1 value is ", beta1)

	beta0 := y_mean - (beta1 * x_mean)
	fmt.Println("beta0 value is ", beta0)
	return beta0, beta1
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
