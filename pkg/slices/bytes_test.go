package slices

import (
	"reflect"
	"testing"
)

func TestCopyChunkedByteSlices_oneChunk(t *testing.T) {
	src := [][]byte{
		[]byte("influx"),
		[]byte("data"),
	}

	dst := CopyChunkedByteSlices(src, 3)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("destination should match source src: %v dst: %v", src, dst)
	}

	dst[0][1] = 'z'
	if reflect.DeepEqual(src, dst) {
		t.Error("destination should not match source")
	}
}

func TestCopyChunkedByteSlices_multipleChunks(t *testing.T) {
	src := [][]byte{
		[]byte("influx"),
		[]byte("data"),
		[]byte("is"),
		[]byte("the"),
		[]byte("best"),
		[]byte("time"),
		[]byte("series"),
		[]byte("database"),
	}

	dst := CopyChunkedByteSlices(src, 3)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("destination should match source src: %v dst: %v", src, dst)
	}

	dst[0][5] = 'z'
	if reflect.DeepEqual(src, dst) {
		t.Error("destination should not match source")
	}
}
