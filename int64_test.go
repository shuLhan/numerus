// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus_test

import (
	"github.com/shuLhan/numerus"
	"testing"
)

func TestInt64CreateSeq(t *testing.T) {
	exp := []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	got := numerus.Int64CreateSeq(-5, 5)

	assert(t, exp, got, true)
}
