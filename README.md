# GeoHash

A lightweight, high-performance Go package for GeoHash operations:

- 🌍 Encode latitude/longitude with adjustable precision.
- 🔄 Decode a GeoHash back to coordinates.
- 🗺️ Bounding Box support for quick region queries.
- 🧭 Neighbor calculation for directional lookups.
- ⚡ Minimal allocations, high performance.

## Install

```bash
go get github.com/aoliveti/geohash
```

## Methods

### Encode
```go
func Encode(latitude, longitude float64, precision Precision) (string, error)
```
Generates a geohash from coordinates.  
See the [Precision Levels](#precision-levels) table for an overview of available spatial resolutions.

### Decode
```go
func Decode(hash string) (latitude, longitude float64, err error)
```
Retrieves coordinates from a geohash.

### DecodeBBox
```go
func DecodeBBox(hash string) (latitude float64, longitude float64, bbox BBox, err error)
```
Returns the bounding box for a geohash.

### Neighbour
```go
func Neighbor(hash string, direction Direction) (string, error)
```
Finds the adjacent geohash in a given direction.  

### Neighbours
```go
func Neighbors(hash string) ([]string, error)
```
Finds all adjacent geohashes in the eight main directions (N, NE, E, SE, S, SW, W, NW).  
See the [Directions](#directions) table for details.

---

## Precision Levels

Supports various precision levels, each offering different spatial resolutions.

| **No.** | **Precision** | **Approximate Coverage**    |
|---------|--------------|------------------------------|
| 1       | Global       | ~5000 km × 5000 km          |
| 2       | Country      | ~1250 km × 625 km           |
| 3       | State        | ~156 km × 156 km            |
| 4       | Region       | ~39 km × 19.5 km            |
| 5       | City         | ~4.9 km × 4.9 km            |
| 6       | Street       | ~1.2 km × 0.61 km           |
| 7       | Building     | ~152 m × 152 m              |
| 8       | Block        | ~38 m × 19 m                |
| 9       | House        | ~4.8 m × 4.8 m              |
| 10      | Room         | ~1.2 m × 0.6 m              |
| 11      | Point        | ~15 cm × 15 cm              |
| 12      | SubPoint     | ~1.9 cm × 1.9 cm            |

## Directions
Indexes the eight principal directions (N, NE, E, SE, S, SW, W, NW) for quick neighbor lookups.

| No. | Direction |
|-----|-----------|
| 0   | N         |
| 1   | NE        |
| 2   | E         |
| 3   | SE        |
| 4   | S         |
| 5   | SW        |
| 6   | W         |
| 7   | NW        |

## Usage

Below is a minimal usage example demonstrating how to encode and decode GeoHash values:

```go
package main

import (
    "fmt"
    "github.com/aoliveti/geohash"
)

func main() {
    // Encode a latitude and longitude with 'City' precision, for example.
    hash, err := geohash.Encode(37.7749, -122.4194, geohash.City)
    if err != nil {
        panic(err)
    }
    fmt.Println("Encoded GeoHash:", hash)

    // Decode the GeoHash string back to latitude/longitude.
    lat, lng, err := geohash.Decode(hash)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Decoded: Latitude=%.4f, Longitude=%.4f\n", lat, lng)
}
```

## Benchmarks
```text
goos: darwin
goarch: amd64
pkg: geohash
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkEncodePrecisionGlobal-4                46509464                24.95 ns/op            0 B/op          0 allocs/op
BenchmarkEncodePrecisionStreet-4                 7434326               148.2 ns/op             8 B/op          1 allocs/op
BenchmarkEncodePrecisionHouse-4                  5502320               205.9 ns/op            16 B/op          1 allocs/op
BenchmarkEncodePrecisionSubPoint-4               4425508               278.1 ns/op            16 B/op          1 allocs/op

BenchmarkDecodePrecisionGlobal-4                38196273                32.83 ns/op            0 B/op          0 allocs/op
BenchmarkDecodePrecisionStreet-4                 8365387               152.7 ns/op             0 B/op          0 allocs/op
BenchmarkDecodePrecisionHouse-4                  5479942               226.1 ns/op             0 B/op          0 allocs/op
BenchmarkDecodePrecisionSubPoint-4               3926744               315.1 ns/op             0 B/op          0 allocs/op

BenchmarkDecodeBBoxPrecisionGlobal-4            37767790                33.75 ns/op            0 B/op          0 allocs/op
BenchmarkDecodeBBoxPrecisionStreet-4             7945538               140.8 ns/op             0 B/op          0 allocs/op
BenchmarkDecodeBBoxPrecisionHouse-4              5773412               217.3 ns/op             0 B/op          0 allocs/op
BenchmarkDecodeBBoxPrecisionSubPoint-4           3792211               298.8 ns/op             0 B/op          0 allocs/op

BenchmarkNeighbourPrecisionGlobal-4             13688127                92.74 ns/op            0 B/op          0 allocs/op
BenchmarkNeighbourPrecisionCity-4                4459608               267.5 ns/op             5 B/op          1 allocs/op
BenchmarkNeighbourPrecisionBlock-4               2901146               421.7 ns/op             8 B/op          1 allocs/op
BenchmarkNeighbourPrecisionSubPoint-4            1776609               638.6 ns/op            16 B/op          1 allocs/op

BenchmarkNeighboursPrecisionGlobal-4             1344180               830.5 ns/op           128 B/op          1 allocs/op
BenchmarkNeighboursPrecisionCity-4                406406              2632 ns/op             170 B/op          9 allocs/op
BenchmarkNeighboursPrecisionBlock-4               323782              3684 ns/op             192 B/op          9 allocs/op
BenchmarkNeighboursPrecisionSubPoint-4            194656              6200 ns/op             256 B/op          9 allocs/op
PASS
ok      geohash 29.995s
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file.