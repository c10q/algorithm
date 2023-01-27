package scan

import "testing"

var testBytes = []byte{50, 51, 52}

func BenchmarkBytesToIntV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToIntV1(testBytes)
	}
}

func BenchmarkBytesToIntV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToIntV2(testBytes)
	}
}

func BenchmarkBytesToIntV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToIntV3(testBytes)
	}
}

func BenchmarkBytesToIntV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToIntV4(testBytes)
	}
}
