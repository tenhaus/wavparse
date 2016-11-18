package wavparse

import (
	"io/ioutil"
	"testing"
)

var wav Wav

// http://soundfile.sapp.org/doc/WaveFormat/
func init() {
	data, _ := ioutil.ReadFile("test/322968.wav")
	New(&wav, data)
}

func TestHeader(t *testing.T) {
	header := wav.Header()
	if header != "RIFF" {
		t.Fail()
	}
}

// 226240 minus 8 bytes for ChunkID and ChunkSize which
// are not included in the total
func TestChunkSize(t *testing.T) {
	size := wav.ChunkSize()
	if size != 226232 {
		t.Fail()
	}
}

func TestFormat(t *testing.T) {
	format := wav.Format()
	if format != "WAVE" {
		t.Fail()
	}
}
