// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numerus

import (
	"math/rand"
	"time"
)

//
// IntPickRandPositive return random integer value from 0 to maximum value
// `maxVal`.
//
// The random value is checked with already picked index: `pickedIds`.
//
// If `dup` is true, allow duplicate value in `pickedIds`, otherwise only
// single unique value allowed in `pickedIds`.
//
// If excluding index `exsIds` is not empty, do not pick the integer value
// listed in there.
//
func IntPickRandPositive(maxVal int, dup bool, pickedIds, exsIds []int) (
	idx int,
) {
	rand.Seed(time.Now().UnixNano())

	excluded := false
	picked := false
	for {
		idx = rand.Intn(maxVal)

		// Check in exclude indices.
		excluded = false
		for _, v := range exsIds {
			if idx == v {
				excluded = true
				break
			}
		}
		if excluded {
			continue
		}

		if dup {
			// Allow duplicate idx.
			return
		}

		// Check if its already picked.
		picked = false
		for _, v := range pickedIds {
			if idx == v {
				picked = true
				break
			}
		}

		if picked {
			// Get another random idx again.
			continue
		}

		// Bingo, we found unique idx that has not been picked.
		return
	}
}
