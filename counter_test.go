package util

import (
	"testing"
	. "github.com/aandryashin/matchers"
)

func TestCounter(t *testing.T) {
	counter := NewCounter()
	AssertThat(t, counter.Count(), EqualTo{uint64(0)})
	AssertThat(t, counter.Count(), EqualTo{uint64(1)})
}
