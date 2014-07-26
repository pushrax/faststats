// Copyright 2014 The faststats Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.

package faststats

import (
	"math"
	"sync/atomic"
)

// Mean calculates the arithmetic mean of arbitrary float64 samples.
type Mean struct {
	value   uint64 // These bits are really a float64.
	total   float64
	samples uint64
}

func NewMean() *Mean {
	return &Mean{}
}

// Value returns the current mean of all samples.
// It is thread-safe, and may be called concurrently with AddSample.
func (p *Mean) Value() float64 {
	bits := atomic.LoadUint64(&p.value)
	return math.Float64frombits(bits)
}

// AddSample adds a single float64 sample to the data set.
// It is not thread-safe, and must not be called in parallel.
func (p *Mean) AddSample(sample float64) {
	p.samples++
	p.total += sample

	bits := math.Float64bits(p.total / float64(p.samples))
	atomic.StoreUint64(&p.value, bits)
}
