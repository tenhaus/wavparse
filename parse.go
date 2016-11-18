package wavparse

import "encoding/binary"

// Header returns the first 4 bytes of the file
// which should equal RIFF
func Header(wav []byte) string {
	return string(wav[0:4])
}

// ChunkSize returns the chunksize
// 4 + (8 + SubChunk1Size) + (8 + SubChunk2Size)
// or 36, which is the meta size + SubChunk2Size
func ChunkSize(wav []byte) uint32 {
	size := binary.LittleEndian.Uint32(wav[4:8])
	return size
}

// Format returns the format which should be WAVE
func Format(wav []byte) string {
	return string(wav[8 : 8+4])
}
