// Copyright 2016-2018, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus_test

import (
	"github.com/shuLhan/numerus"
	"testing"
)

func TestIntCreateSeq(t *testing.T) {
	exp := []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	got := numerus.IntCreateSeq(-5, 5)

	assert(t, exp, got, true)
}

func TestIntPickRandPositive(t *testing.T) {
	pickedIds := []int{0, 1, 2, 3, 4, 5, 7}
	exsIds := []int{8, 9}
	exp := 6

	// Pick random without duplicate.
	got := numerus.IntPickRandPositive(7, false, pickedIds, nil)

	assert(t, exp, got, true)

	// Pick random with exclude indices.
	got = numerus.IntPickRandPositive(9, false, pickedIds, exsIds)
	assert(t, exp, got, true)
}
