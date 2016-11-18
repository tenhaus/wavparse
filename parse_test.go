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

func TestChunkID(t *testing.T) {
	header := wav.ChunkID()
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

func TestSubchunk1ID(t *testing.T) {
	data := wav.Subchunk1ID()
	if data != "fmt " {
		t.Fail()
	}
}

func TestSubchunk1Size(t *testing.T) {
	size := wav.Subchunk1Size()
	if size != 16 {
		t.Fail()
	}
}

func TestAudioFormat(t *testing.T) {
	format := wav.AudioFormat()
	if format != 3 {
		t.Fail()
	}
}

func TestNumChannels(t *testing.T) {
	channels := wav.NumChannels()
	if channels != 1 {
		t.Fail()
	}
}

func TestSampleRate(t *testing.T) {
	rate := wav.SampleRate()
	if rate != 44100 {
		t.Fail()
	}
}

func TestByteRate(t *testing.T) {
	rate := wav.ByteRate()
	if rate != 45328 {
		t.Fail()
	}
}

func TestBlockAlign(t *testing.T) {
	align := wav.BlockAlign()
	if align != 4 {
		t.Fail()
	}
}

func TestBitsPerSample(t *testing.T) {
	bps := wav.BlockAlign()
	if bps != 4 {
		t.Fail()
	}
}
