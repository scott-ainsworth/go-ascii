package ascii

import (
	"math/rand"
	"testing"
	"unicode"
)

//******************************************************************************
// Benchmarks
//******************************************************************************

func generateSample(n int, min, max int) []byte {
	spread := max - min
	sample := make([]byte, n)
	rs := rand.NewSource(16)
	rn := rand.New(rs)
	for i := 0; i < len(sample); i++ {
		sample[i] = byte(min + rn.Intn(spread))
	}
	return sample
}

var sample = generateSample(1000, 0, 256)

func BenchmarkIsASCII(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsASCII(sample[i])
		}
	}
}

func BenchmarkIsLetter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsLetter(sample[i])
		}
	}
}

func BenchmarkIsUpper(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsUpper(sample[i])
		}
	}
}

func BenchmarkIsLower(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsLower(sample[i])
		}
	}
}

func BenchmarkIsDigit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsDigit(sample[i])
		}
	}
}

func BenchmarkIsHexDigit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsHexDigit(sample[i])
		}
	}
}

func BenchmarkIsAlnum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsAlnum(sample[i])
		}
	}
}

func BenchmarkIsSpace(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsSpace(sample[i])
		}
	}
}

func BenchmarkIsControl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsControl(sample[i])
		}
	}
}

func BenchmarkIsGraphic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsGraph(sample[i])
		}
	}
}

func BenchmarkIsPrint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			IsPrint(sample[i])
		}
	}
}

func BenchmarkIsPrintUnicode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(sample); i++ {
			unicode.IsPrint(rune(sample[i]))
		}
	}
}

// eof
