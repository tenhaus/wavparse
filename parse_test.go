package wavparse

import (
	"fmt"
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
	fmt.Printf("ChunkID %v\n", header)
	if header != "RIFF" {
		t.Fail()
	}
}

// 226240 minus 8 bytes for ChunkID and ChunkSize which
// are not included in the total
func TestChunkSize(t *testing.T) {
	size := wav.ChunkSize()
	fmt.Printf("ChunkSize %v\n", size)
	if size != 226232 {
		t.Fail()
	}
}

func TestFormat(t *testing.T) {
	format := wav.Format()
	fmt.Printf("Format %v\n", format)
	if format != "WAVE" {
		t.Fail()
	}
}

func TestSubchunk1ID(t *testing.T) {
	data := wav.Subchunk1ID()
	fmt.Printf("SubChunk1ID %v\n", data)
	if data != "fmt " {
		t.Fail()
	}
}

func TestSubchunk1Size(t *testing.T) {
	size := wav.Subchunk1Size()
	fmt.Printf("Subchunk1Size %v\n", size)
	if size != 16 {
		t.Fail()
	}
}

func TestAudioFormat(t *testing.T) {
	format := wav.AudioFormat()
	fmt.Printf("AudioFormat %v\n", format)
	if format != 3 {
		t.Fail()
	}
}

func TestNumChannels(t *testing.T) {
	channels := wav.NumChannels()
	fmt.Printf("NumChannels %v\n", channels)
	if channels != 1 {
		t.Fail()
	}
}

func TestSampleRate(t *testing.T) {
	rate := wav.SampleRate()
	fmt.Printf("SampleRate %v\n", rate)
	if rate != 44100 {
		t.Fail()
	}
}

func TestByteRate(t *testing.T) {
	rate := wav.ByteRate()
	fmt.Printf("ByteRate %v\n", rate)
	if rate != 45328 {
		t.Fail()
	}
}

func TestBlockAlign(t *testing.T) {
	align := wav.BlockAlign()
	fmt.Printf("BlockAlign %v\n", align)
	if align != 4 {
		t.Fail()
	}
}

func TestBitsPerSample(t *testing.T) {
	bps := wav.BitsPerSample()
	fmt.Printf("BitsPerSample %v\n", bps)
	if bps != 32 {
		t.Fail()
	}
}

func TestExtraParamSize(t *testing.T) {
	extra := wav.ExtraParamSize()
	fmt.Printf("ExtraParamSize %v\n", extra)
	if extra != 24934 {
		t.Fail()
	}
}

func TestExtraParams(t *testing.T) {
	// extra := wav.ExtraParams()
	// fmt.Println(extra)
}

func TestSubchunk2ID(t *testing.T) {
	id := wav.Subchunk2ID()
	fmt.Println(id)
	if id != "data" {
		t.Fail()
	}
}
