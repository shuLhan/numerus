// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus_test

import (
	"github.com/shuLhan/numerus"
	"testing"
)

var (
	dInts = [][]int{
		{},
		{5, 6, 7, 8, 9, 0, 1, 2, 3, 4},
		{0, 1, 0, 1, 0, 1, 0, 1, 0},
	}
)

func TestIntsFindMaxEmpty(t *testing.T) {
	maxv, maxi, ok := numerus.IntsFindMax(dInts[0])

	assert(t, -1, maxv, true)
	assert(t, -1, maxi, true)
	assert(t, false, ok, true)
}

func TestIntsFindMax(t *testing.T) {
	maxv, maxi, ok := numerus.IntsFindMax(dInts[1])

	assert(t, 9, maxv, true)
	assert(t, 4, maxi, true)
	assert(t, true, ok, true)
}

func TestIntsFindMinEmpty(t *testing.T) {
	minv, mini, ok := numerus.IntsFindMin(dInts[0])

	assert(t, -1, minv, true)
	assert(t, -1, mini, true)
	assert(t, false, ok, true)
}

func TestIntsFindMin(t *testing.T) {
	minv, mini, ok := numerus.IntsFindMin(dInts[1])

	assert(t, 0, minv, true)
	assert(t, 5, mini, true)
	assert(t, true, ok, true)
}

func TestIntsSum(t *testing.T) {
	got := numerus.IntsSum(dInts[1])

	assert(t, 45, got, true)
}

func TestIntsMaxCountOf(t *testing.T) {
	classes := []int{0, 1}
	exp := int(0)
	got, _ := numerus.IntsMaxCountOf(dInts[2], classes)

	assert(t, exp, got, true)

	// Swap the class values.
	classes = []int{1, 0}
	got, _ = numerus.IntsMaxCountOf(dInts[2], classes)

	assert(t, exp, got, true)
}

func TestIntsIsExist(t *testing.T) {
	var s bool

	// True positive.
	for _, d := range dInts {
		for _, v := range d {
			s = numerus.IntsIsExist(d, v)

			assert(t, true, s, true)
		}
	}

	// False positive.
	for _, d := range dInts {
		s = numerus.IntsIsExist(d, -1)
		assert(t, false, s, true)
		s = numerus.IntsIsExist(d, 10)
		assert(t, false, s, true)
	}
}
