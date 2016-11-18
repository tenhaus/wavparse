package wavparse

import (
	"io/ioutil"
	"testing"
)

var wav []byte

// http://soundfile.sapp.org/doc/WaveFormat/
func init() {
	wav, _ = ioutil.ReadFile("test/322968.wav")
}

func TestHeader(t *testing.T) {
	header := Header(wav)
	if string(header) != "RIFF" {
		t.Fail()
	}
}

// 226240 minus 8 bytes for ChunkID and ChunkSize which
// are not included in the total
func TestChunkSize(t *testing.T) {
	size := ChunkSize(wav)
	if size != 226232 {
		t.Fail()
	}
}
