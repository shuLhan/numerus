// Copyright 2016-2018, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus_test

import (
	"github.com/shuLhan/numerus"
	"testing"
)

func TestFloat64Round(t *testing.T) {
	data := []float64{0.553, -0.553, 0.49997, -0.49997, 0.4446, -0.4446}
	exps := [][]float64{
		{1, 0.6, 0.55},
		{-1, -0.6, -0.55},
		{0.0, 0.5, 0.5},
		{0.0, -0.5, -0.5},
		{0, 0.4, 0.44},
		{0, -0.4, -0.44},
	}

	for x := range data {
		for y, exp := range exps[x] {
			got := numerus.Float64Round(data[x], y)

			assert(t, exp, got, true)
		}
	}
}
