package main

import (
	"fmt"
	"math"
	"sort"
)

func geometricAverage(values []int) float64 {
	var sum int
	for _, v := range values {
		sum += v
	}

	solution := math.Pow(float64(sum), 1.0/float64(len(values)))
	return solution
}

func modal(values []int) int {
	m, _ := absolute(values)
	var maxFreq int

	for _, freq := range m {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	return maxFreq
}

func arithmetricAverage(values []int) float64 {
	var sum int
	for _, v := range values {
		sum += v
	}

	result := float64(sum) / float64(len(values))
	return result
}

func quantile(values []int, percentage float64) int {
	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Ints(sorted)

	position := 1 + float64(len(sorted)-1)*percentage
	index := int(position)
	if position > float64(index) {
		index++
	}
	index--

	return sorted[index]
}

func absolute(values []int) (map[int]int, int) {
	m := make(map[int]int)
	length := len(values)

	for _, v := range values {
		_, exists := m[v]

		if exists == false {
			m[v] = 1
		}

		if exists == true {
			m[v] += 1
		}
	}

	return m, length
}

func relative(values map[int]int, length int) map[int]float64 {
	relmap := make(map[int]float64)
	for key, val := range values {
		var result float64
		result = float64(val) / float64(length)
		relmap[key] = result
	}
	return relmap
}

func cumulativeabs(values map[int]int) map[int]int {
	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	cumulative := make(map[int]int)
	sum := 0

	for _, k := range keys {
		sum += values[k]
		cumulative[k] = sum
	}
	return cumulative
}

func cumulativerel(values map[int]float64) map[int]float64 {
	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	cumulative := make(map[int]float64)
	var sum float64
	sum = 0

	for _, k := range keys {
		sum += values[k]
		cumulative[k] = sum
	}
	return cumulative
}

func standDev(data []int) float64 {
	avg := arithmetricAverage(data)
	var res float64

	for _, v := range data {
		resbetween := math.Pow((float64(v) - avg), 2)
		res += resbetween
	}

	result := math.Sqrt(res / float64(len(data)-1))
	return result
}

func coVar(data1 []int, data2 []int) float64 {
	avg1 := arithmetricAverage(data1)
	avg2 := arithmetricAverage(data2)
	var res float64

	for i, v := range data1 {
		resbetween1 := float64(v) - avg1
		resbetween2 := float64(data2[i]) - avg2
		res += resbetween1 * resbetween2
	}

	result := res / float64(len(data1)-1)
	return result
}

func coRelation(data1 []int, data2 []int) float64 {
	result := coVar(data1, data2) / (standDev(data1) * standDev(data2))
	return result
}

func main() {
	var data1 = []int{5,7,2,2,8,5,10,2,6,2}
	var data2 = []int{10,1,2,1,8,2,3,1,10,5}

	fmt.Println(coRelation(data1, data2))
}
