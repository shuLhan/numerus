// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus_test

import (
	"github.com/shuLhan/numerus"
	"testing"
)

func BenchmarkFloats64InplaceMergesort(b *testing.B) {
	size := len(inSorts[6])
	ids := make([]int, size)

	for i := 0; i < b.N; i++ {
		numerus.Floats64InplaceMergesort(inSorts[6], ids, 0, size)
	}
}
