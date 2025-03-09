package geohash

import "testing"

func BenchmarkEncodePrecisionGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Encode(37.7749, -122.4194, Global)
	}
}

func BenchmarkEncodePrecisionStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Encode(37.7749, -122.4194, Street)
	}
}

func BenchmarkEncodePrecisionHouse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Encode(37.7749, -122.4194, House)
	}
}

func BenchmarkEncodePrecisionSubPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Encode(37.7749, -122.4194, SubPoint)
	}
}

func BenchmarkDecodePrecisionGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = Decode("9")
	}
}

func BenchmarkDecodePrecisionStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = Decode("9q8yyk")
	}
}

func BenchmarkDecodePrecisionHouse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = Decode("9q8yyk8yp")
	}
}

func BenchmarkDecodePrecisionSubPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = Decode("9q8yyk8ypd23")
	}
}

func BenchmarkDecodeBBoxPrecisionGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, _ = DecodeBBox("9")
	}
}

func BenchmarkDecodeBBoxPrecisionStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, _ = DecodeBBox("9q8yyk")
	}
}

func BenchmarkDecodeBBoxPrecisionHouse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, _ = DecodeBBox("9q8yyk8yp")
	}
}

func BenchmarkDecodeBBoxPrecisionSubPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, _ = DecodeBBox("9q8yyk8ypd23")
	}
}

func BenchmarkNeighbourPrecisionGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbor("9", N)
	}
}

func BenchmarkNeighbourPrecisionCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbor("9q8yy", N)
	}
}

func BenchmarkNeighbourPrecisionBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbor("9q8yyk8y", N)
	}
}

func BenchmarkNeighbourPrecisionSubPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbor("9q8yyk8ypd23", N)
	}
}

func BenchmarkNeighboursPrecisionGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbors("9")
	}
}

func BenchmarkNeighboursPrecisionCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbors("9q8yy")
	}
}

func BenchmarkNeighboursPrecisionBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbors("9q8yyk8y")
	}
}

func BenchmarkNeighboursPrecisionSubPoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Neighbors("9q8yyk8ypd23")
	}
}
