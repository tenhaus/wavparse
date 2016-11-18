package wavparse

// Header returns the first 4 bytes of the file
// which should equal RIFF
func Header(wav []byte) []byte {
	return wav[0:4]
}
