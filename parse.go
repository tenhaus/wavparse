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

// ChunkID returns the first 4 bytes of the file
// which should equal RIFF
func (wav *Wav) ChunkID() string {
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

// Subchunk1ID [12:4] Big
func (wav *Wav) Subchunk1ID() string {
	return string(wav.Bytes[12 : 12+4])
}

// Subchunk1Size [16:4]  little
func (wav *Wav) Subchunk1Size() uint32 {
	bytes := wav.Bytes[16 : 16+4]
	size := binary.LittleEndian.Uint32(bytes)
	return size
}

// AudioFormat [20:2] little
func (wav *Wav) AudioFormat() uint16 {
	bytes := wav.Bytes[20 : 20+2]
	format := binary.LittleEndian.Uint16(bytes)
	return format
}

// NumChannels [22:2] little
func (wav *Wav) NumChannels() uint16 {
	bytes := wav.Bytes[22 : 22+2]
	channels := binary.LittleEndian.Uint16(bytes)
	return channels
}

// SampleRate [24:4] little
func (wav *Wav) SampleRate() string {
	return "hi"
}

// ByteRate [28:4] little
func (wav *Wav) ByteRate() string {
	return "hi"
}

// BlockAlign [32:2] little
func (wav *Wav) BlockAlign() string {
	return "hi"
}

// BitsPerSample [34:2] little
func (wav *Wav) BitsPerSample() string {
	return "hi"
}

// ExtraParamSize [34+2:2]
func (wav *Wav) ExtraParamSize() string {
	return "hi"
}

// ExtraParams [34+4:ExtraParamSize]
func (wav *Wav) ExtraParams() string {
	return "hi"
}

// Subchunk2ID [36:4]
func (wav *Wav) Subchunk2ID() string {
	return "hi"
}

// Subchunk2Size [40:4]
func (wav *Wav) Subchunk2Size() string {
	return "hi"
}

// Data [44:]
func (wav *Wav) Data() string {
	return "hi"
}
