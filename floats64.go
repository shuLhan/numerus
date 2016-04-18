// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus

const (
	// SortThreshold when the data less than SortThreshold, insertion sort
	// will be used to replace current sort.
	SortThreshold = 7
)

//
// Floats64FindMax given slice of float, find the maximum value in slice and
// and return it with their index.
//
// If data is empty, return -1 in value and index, and false in ok.
//
// Example, given data: [0.0 0.1 0.2 0.2 0.4], it will return 0.4 as max and 4
// as index of maximum value.
//
func Floats64FindMax(d []float64) (maxv float64, maxi int, ok bool) {
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
// Floats64FindMin given slice of float, find the minimum value in slice and
// and return it with their index.
//
// If data is empty, return -1 in value and index, and false in ok.
//
// Example, given data: [0.0 0.1 0.2 0.2 0.4], it will return 0 as min and 0
// as index of minimum value.
//
func Floats64FindMin(d []float64) (minv float64, mini int, ok bool) {
	l := len(d)
	if l <= 0 {
		return -1, -1, false
	}

	x := 0
	minv = d[x]
	mini = x

	for x = 1; x < l; x++ {
		if d[x] > minv {
			minv = d[x]
			mini = x
		}
	}

	return minv, mini, true
}

//
// Floats64Sum return sum of slice of float64.
//
func Floats64Sum(d []float64) (sum float64) {
	for _, v := range d {
		sum += v
	}
	return
}

//
// Floats64Count will count number of class in data.
//
func Floats64Count(d []float64, class float64) (count int) {
	if len(d) <= 0 {
		return
	}

	for _, v := range d {
		if v == class {
			count++
		}
	}
	return count
}

//
// Floats64Counts will count class in data and return each of the counter.
//
// For example, if data is "[1,1,2]" and class is "[1,2]", this function will
// return "[2,1]".
//
//	idx class  count
//	0 : 1   -> 2
//	1 : 2   -> 1
//
//
func Floats64Counts(d []float64, classes []float64) (counts []int) {
	if len(classes) <= 0 {
		return
	}

	counts = make([]int, len(classes))

	for x, c := range classes {
		counts[x] = Floats64Count(d, c)
	}
	return
}

//
// Floats64Swap swap two indices value of 64bit float.
//
func Floats64Swap(d []float64, x, y int) {
	if x == y {
		return
	}

	tmp := d[x]
	d[x] = d[y]
	d[y] = tmp
}

//
// Floats64InsertionSort will sort the data using insertion-sort algorithm.
//
// Parameters:
// - `data` is slice that will be sorted.
// - `idx` is indices of data.
// - `l` is starting index of slice to be sorted.
// - `r` is end index of slice to be sorted.
//
func Floats64InsertionSort(d []float64, ids []int, l, r int) {
	for x := l; x < r; x++ {
		for y := x + 1; y < r; y++ {
			if d[x] > d[y] {
				IntsSwap(ids, x, y)
				Floats64Swap(d, x, y)
			}
		}
	}
}

//
// Floats64SortByIndex will sort the slice of float using sorted index.
//
func Floats64SortByIndex(d *[]float64, sortedIds []int) {
	newd := make([]float64, len(*d))

	for i := range sortedIds {
		newd[i] = (*d)[sortedIds[i]]
	}

	(*d) = newd
}

//
// Floats64InplaceMergesort in-place merge-sort without memory allocation.
//
// Algorithm,
//
// (0) If data length == Threshold, then
// (0.1) use insertion sort.
// (1) Divide into left and right.
// (2) Sort left.
// (3) Sort right.
// (4) Merge sorted left and right.
// (4.1) If the last element of the left is lower then the first element of the
//       right, i.e. [1 2] [3 4]; no merging needed, return immediately.
// (4.2) Let x be the first index of left-side, and y be the first index of
//       the right-side.
// (4.3) Loop until either x or y reached the maximum slice.
// (4.3.1) IF DATA[x] <= DATA[y]
// (4.3.1.1) INCREMENT x
// (4.3.1.2) IF x > y THEN GOTO 4.3
// (4.3.1.3) GOTO 4.3.4
// (4.3.2) LET YLAST := the next DATA[y] that is less DATA[x]
// (4.3.3) SWAP DATA, X, Y, YLAST
// (4.3.4) LET Y := the next DATA that has minimum value between x and r
//
func Floats64InplaceMergesort(d []float64, idx []int, l, r int) {
	// (0)
	if l+SortThreshold >= r {
		// (0.1)
		Floats64InsertionSort(d, idx, l, r)
		return
	}

	// (1)
	res := (r + l) % 2
	c := (r + l) / 2
	if res == 1 {
		c++
	}

	// (2)
	Floats64InplaceMergesort(d, idx, l, c)

	// (3)
	Floats64InplaceMergesort(d, idx, c, r)

	// (4)
	if d[c-1] <= d[c] {
		// (4.1)
		return
	}

	// (4.2)
	x := l
	y := c
	ylast := c

	// (4.3)
	for x < r && y < r {
		// (4.3.1)
		if d[x] <= d[y] {
			x++

			// (4.3.1.2)
			if x >= y {
				goto next
			}

			// (4.3.1.3)
			continue
		}

		// (4.3.2)
		ylast = movey(d, x, y, r)

		// (4.3.3)
		ylast = multiswap(d, idx, x, y, ylast)

	next:
		// (4.3.4)
		for x < r {
			y = min(d, x, r)
			if y == x {
				x++
			} else {
				break
			}
		}
	}
}

func movey(d []float64, x, y, r int) int {
	yorg := y
	y++
	for y < r {
		if d[y] >= d[x] {
			break
		}
		if d[y] < d[yorg] {
			break
		}
		y++
	}
	return y
}

func multiswap(d []float64, idx []int, x, y, ylast int) int {
	for y < ylast {
		IntsSwap(idx, x, y)
		Floats64Swap(d, x, y)
		x++
		y++
		if y >= ylast {
			return y
		}
		if d[x] <= d[y] {
			return y
		}
	}

	return y
}

func min(d []float64, l, r int) (m int) {
	min := d[l]
	m = l
	for l++; l < r; l++ {
		if d[l] <= min {
			min = d[l]
			m = l
		}
	}
	return
}

//
// Floats64IndirectSort will sort the data and return the sorted index.
//
func Floats64IndirectSort(d []float64) (sortedIdx []int) {
	dlen := len(d)

	sortedIdx = make([]int, dlen)
	for i := 0; i < dlen; i++ {
		sortedIdx[i] = i
	}

	Floats64InplaceMergesort(d, sortedIdx, 0, dlen)

	return
}
