package slices

// BytesToStrings converts a slice of []byte into a slice of strings.
func BytesToStrings(a [][]byte) []string {
	s := make([]string, 0, len(a))
	for _, v := range a {
		s = append(s, string(v))
	}
	return s
}

// CopyChunkedByteSlices deep-copies a [][]byte to a new [][]byte that is backed by a small number of []byte "chunks".
func CopyChunkedByteSlices(a [][]byte, chunkSize int) [][]byte {
	b := make([][]byte, len(a))
	for batchBegin := 0; batchBegin < len(a); batchBegin += chunkSize {
		batchEnd := len(a)
		if batchEnd-batchBegin > chunkSize {
			batchEnd = batchBegin + chunkSize
		}

		batchSize := 0
		for j := batchBegin; j < batchEnd; j++ {
			batchSize += len(a[j])
		}
		batch := make([]byte, batchSize)
		batchOffset := 0
		for j := batchBegin; j < batchEnd; j++ {
			copy(batch[batchOffset:batchOffset+len(a[j])], a[j])
			b[j] = batch[batchOffset : batchOffset+len(a[j])]
			batchOffset += len(a[j])
		}
	}
	return b
}
