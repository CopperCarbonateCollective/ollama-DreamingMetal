//go:build !cgo
// +build !cgo

package ggml

import (
	fsggml "github.com/CopperCarbonateCollective/ollama-DreamingMetal/fs/ggml"
)

// ConvertToF32 is a noop fallback used when cgo is disabled.
// It returns a zeroed slice of the requested length so callers won't panic during analysis.
func ConvertToF32(data []byte, dtype uint32, nelements uint64) []float32 {
	// return zeroed slice to keep callers safe during static analysis / non-cgo builds
	return make([]float32, int(nelements))
}

// Quantize is a noop fallback used when cgo is disabled.
func Quantize(newType fsggml.TensorType, f32s []float32, shape []uint64) []byte {
	return nil
}

// QuantizationVersion returns 0 in non-cgo builds.
func QuantizationVersion() uint32 {
	return 0
}
