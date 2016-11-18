package wavparse

import (
	"io/ioutil"
	"testing"
)

var wav []byte

func init() {
	wav, _ = ioutil.ReadFile("test/322968.wav")
}

func TestHeader(t *testing.T) {
	header := Header(wav)
	if string(header) != "RIFF" {
		t.Fail()
	}
}
