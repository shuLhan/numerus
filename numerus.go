// Copyright 2016-2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// Package numerus provide miscellaneous functions for working with integer,
// float, slice of integer, and slice of floats.
//
// Currently it have function to,
// - sort slice of floats using in-place mergesort
// algorithm.
// - sort slice of integer/floats by predefined index
// - count number of value occurence in slice of integer/float
// - find minimum or maximum value in slice of integer/float
// - sum slice of integer/float
//
package numerus

const (
	// SortThreshold when the data less than SortThreshold, insertion sort
	// will be used to replace mergesort.
	SortThreshold = 7
)
