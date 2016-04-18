// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus_test

import (
	"github.com/shuLhan/numerus"
	"testing"
)

var (
	dataInts64 = [][]int64{
		{},
		{5, 6, 7, 8, 9, 0, 1, 2, 3, 4},
		{0, 1, 0, 1, 0, 1, 0, 1, 0},
	}
)

func TestInts64FindMaxEmpty(t *testing.T) {
	gotv, goti, gotok := numerus.Ints64FindMax(dataInts64[0])

	assert(t, int64(-1), gotv, true)
	assert(t, -1, goti, true)
	assert(t, false, gotok, true)
}

func TestInts64FindMax(t *testing.T) {
	gotv, goti, gotok := numerus.Ints64FindMax(dataInts64[1])

	assert(t, int64(9), gotv, true)
	assert(t, 4, goti, true)
	assert(t, true, gotok, true)
}

func TestInts64FindMinEmpty(t *testing.T) {
	gotv, goti, gotok := numerus.Ints64FindMin(dataInts64[0])

	assert(t, int64(-1), gotv, true)
	assert(t, -1, goti, true)
	assert(t, false, gotok, true)
}

func TestInts64FindMin(t *testing.T) {
	gotv, goti, gotok := numerus.Ints64FindMin(dataInts64[1])

	assert(t, int64(0), gotv, true)
	assert(t, 5, goti, true)
	assert(t, true, gotok, true)
}

func TestInts64Sum(t *testing.T) {
	got := numerus.Ints64Sum(dataInts64[1])

	assert(t, int64(45), got, true)
}

func TestInts64MaxCountOf(t *testing.T) {
	classes := []int64{0, 1}
	exp := int64(0)
	got, _ := numerus.Ints64MaxCountOf(dataInts64[2], classes)

	assert(t, exp, got, true)

	// Swap the class values.
	classes = []int64{1, 0}
	got, _ = numerus.Ints64MaxCountOf(dataInts64[2], classes)

	assert(t, exp, got, true)
}
