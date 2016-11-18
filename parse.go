package wavparse

import "encoding/binary"

// Wav provided functions for working with a WAV file
type Wav struct {
	Bytes []byte
}

// New instantiates a new Wav object
func New(wav *Wav, data []byte) {
	wav.Bytes = data
}

// Header returns the first 4 bytes of the file
// which should equal RIFF
func (wav *Wav) Header() string {
	return string(wav.Bytes[0:4])
}

// ChunkSize returns the chunksize
// 4 + (8 + SubChunk1Size) + (8 + SubChunk2Size)
// or 36, which is the meta size + SubChunk2Size
func (wav *Wav) ChunkSize() uint32 {
	size := binary.LittleEndian.Uint32(wav.Bytes[4:8])
	return size
}

// Format returns the format which should be WAVE
func (wav *Wav) Format() string {
	return string(wav.Bytes[8 : 8+4])
}

func (wav *Wav) SubChunk1ID() string {
	return "hi"
}
