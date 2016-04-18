// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus

//
// Ints64FindMax given a slice of integer, return the maximum value in slice
// and its index.
//
// If data is empty, return -1 in value and index, and false in ok.
//
func Ints64FindMax(d []int64) (maxv int64, maxi int, ok bool) {
	l := len(d)
	if l <= 0 {
		return -1, -1, false
	}

	x := 0
	maxv = d[x]
	maxi = x

	for x = 1; x < l; x++ {
		if d[x] > maxv {
			maxv = d[x]
			maxi = x
		}
	}

	return maxv, maxi, true
}

//
// Ints64FindMin given a slice of integer, return the minimum value in slice
// and its index.
//
// If data is empty, return -1 in value and index; and false in ok.
//
func Ints64FindMin(d []int64) (minv int64, mini int, ok bool) {
	l := len(d)
	if l <= 0 {
		return -1, -1, false
	}

	x := 0
	minv = d[x]
	mini = x

	for x = 1; x < l; x++ {
		if d[x] < minv {
			minv = d[x]
			mini = x
		}
	}

	return minv, mini, true
}

//
// Ints64Sum return sum of all integer value.
//
func Ints64Sum(d []int64) (sum int64) {
	for _, v := range d {
		sum += v
	}
	return sum
}

//
// Ints64MaxCountOf will count number of occurence of each element of classes
// in data and return the class with maximum count.
//
// If `classes` is empty, it will return -1 and false.
// If `data` is empty, it will return -2 and false.
// If classes has the same count value, then the first max in the class will be
// returned.
//
// For example, given a data [0, 1, 0, 1, 0] and classes [0, 1], the function
// will count 0 as 3, 1 as 2; and return 0.
//
func Ints64MaxCountOf(d, classes []int64) (int64, bool) {
	if len(classes) == 0 {
		return -1, false
	}
	if len(d) == 0 {
		return -2, false
	}

	counts := make([]int, len(classes))

	for x, c := range classes {
		for _, d := range d {
			if c == d {
				counts[x]++
			}
		}
	}

	_, maxi, _ := IntsFindMax(counts)

	return classes[maxi], true
}
