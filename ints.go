// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus

//
// IntsFindMax given slice of integer, return the maximum value in slice and
// index of maximum value.
//
// If data is empty, return -1 in value and index, and false in ok.
//
// Example, given a slice of data: [0 1 2 3 4], it will return 4 as max and 4
// as index of maximum value.
//
func IntsFindMax(d []int) (maxv int, maxi int, ok bool) {
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
// IntsFindMin given slice of integer, return the minimum value in slice and
// index of minimum value.
//
// If data is empty, return -1 in value and index.
//
// Example, given a slice of data: [0 1 2 3 4], it will return 0 as min and 0
// as minimum index.
//
func IntsFindMin(d []int) (minv int, mini int, ok bool) {
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
// IntsSum return sum of all value in slice.
//
func IntsSum(d []int) (sum int) {
	for _, v := range d {
		sum += v
	}
	return
}

//
// IntsMaxCountOf will count number of occurence of each element of classes
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
func IntsMaxCountOf(d, classes []int) (int, bool) {
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

//
// IntsSwap swap two indices value of integer.
//
func IntsSwap(d []int, x, y int) {
	if x == y {
		return
	}

	tmp := d[x]
	d[x] = d[y]
	d[y] = tmp
}

//
// IntsIsExist check if integer value exist in slice of integer, return true if
// exist, false otherwise.
//
func IntsIsExist(d []int, i int) bool {
	for _, v := range d {
		if i == v {
			return true
		}
	}
	return false
}

//
// IntsTo64 convert slice of integer to 64bit values.
//
func IntsTo64(ints []int) []int64 {
	i64 := make([]int64, len(ints))
	for x, v := range ints {
		i64[x] = int64(v)
	}
	return i64
}
