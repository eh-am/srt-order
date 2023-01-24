package internal_test

import (
	"testing"

	"github.com/eh-am/srt-order/internal"
)

func TestProcess(t *testing.T) {
	//	err := internal.Process("testdata/simple.srt")
	err := internal.Process("testdata/nichijou-s01e01.srt", false)
	if err != nil {
		t.Error(err)
	}
}

func TestEmptyContent(t *testing.T) {
	err := internal.Process("testdata/empty-content.srt", false)
	if err != nil {
		t.Error(err)
	}
}
