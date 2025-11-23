# GeoHash
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aoliveti/geohash)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/aoliveti/geohash/go.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/aoliveti/geohash)](https://pkg.go.dev/github.com/aoliveti/geohash)
[![codecov](https://codecov.io/gh/aoliveti/geohash/graph/badge.svg?token=TINZl3wzTv)](https://codecov.io/gh/aoliveti/geohash)
[![Go Report Card](https://goreportcard.com/badge/github.com/aoliveti/geohash)](https://goreportcard.com/report/github.com/aoliveti/geohash)
![GitHub License](https://img.shields.io/github/license/aoliveti/geohash)

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

### Neighbors
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
pkg: github.com/aoliveti/geohash
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkEncodePrecisionGlobal-4                45100946                23.69 ns/op            0 B/op          0 allocs/op
BenchmarkEncodePrecisionStreet-4                 8274196               141.5 ns/op             8 B/op          1 allocs/op
BenchmarkEncodePrecisionHouse-4                  5921347               199.3 ns/op            16 B/op          1 allocs/op
BenchmarkEncodePrecisionSubPoint-4               4427629               280.4 ns/op            16 B/op          1 allocs/op

BenchmarkDecodePrecisionGlobal-4                45866013                21.86 ns/op            0 B/op          0 allocs/op
BenchmarkDecodePrecisionStreet-4                11939612                99.55 ns/op            0 B/op          0 allocs/op
BenchmarkDecodePrecisionHouse-4                  8022915               149.3 ns/op             0 B/op          0 allocs/op
BenchmarkDecodePrecisionSubPoint-4               5940921               201.7 ns/op             0 B/op          0 allocs/op

BenchmarkDecodeBBoxPrecisionGlobal-4            53141138                22.25 ns/op            0 B/op          0 allocs/op
BenchmarkDecodeBBoxPrecisionStreet-4            11948359                98.76 ns/op            0 B/op          0 allocs/op
BenchmarkDecodeBBoxPrecisionHouse-4              7960542               148.7 ns/op             0 B/op          0 allocs/op
BenchmarkDecodeBBoxPrecisionSubPoint-4           5864090               202.6 ns/op             0 B/op          0 allocs/op

BenchmarkNeighbourPrecisionGlobal-4             18967950                61.04 ns/op            0 B/op          0 allocs/op
BenchmarkNeighbourPrecisionCity-4                5318504               224.0 ns/op             5 B/op          1 allocs/op
BenchmarkNeighbourPrecisionBlock-4               3469327               340.6 ns/op             8 B/op          1 allocs/op
BenchmarkNeighbourPrecisionSubPoint-4            2381121               497.5 ns/op            16 B/op          1 allocs/op

BenchmarkNeighborsPrecisionGlobal-4             3441564               351.8 ns/op           128 B/op          1 allocs/op
BenchmarkNeighborsPrecisionCity-4                756705              1347 ns/op             170 B/op          9 allocs/op
BenchmarkNeighborsPrecisionBlock-4               536829              1970 ns/op             192 B/op          9 allocs/op
BenchmarkNeighborsPrecisionSubPoint-4            399649              2808 ns/op             256 B/op          9 allocs/op

PASS
ok      github.com/aoliveti/geohash     26.838s
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file.