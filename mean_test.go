// Copyright 2014 The faststats Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.

package faststats

import (
	"math/rand"
	"testing"
	"time"
)

func TestMean(t *testing.T) {
	rand.Seed(time.Now().Unix())

	testMean(t, uniform(10000, 1), 0.5)
	testMean(t, uniform(10000, 50), 25)
}

func testMean(t *testing.T, numbers []float64, expected float64) {
	p := NewMean()

	for i := 0; i < len(numbers); i++ {
		p.AddSample(numbers[i])
	}

	got := p.Value()

	if got < expected*0.99 || got > expected*1.01 {
		t.Errorf("Mean out of range\n  actual: %f\nexpected: %f\n", got, expected)
	}
}
