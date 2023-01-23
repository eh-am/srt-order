package internal_test

import (
	"testing"

	"github.com/eh-am/srt-order/internal"
)

func TestProce(t *testing.T) {
	//	err := internal.Process("testdata/simple.srt")
	err := internal.Process("testdata/nichijou-s01e01.srt")
	if err != nil {
		t.Error(err)
	}
}
