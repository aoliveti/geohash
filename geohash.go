package geohash

import (
	"errors"
	"math"
)

const (
	minLatitude  float64 = -90
	maxLatitude  float64 = 90
	minLongitude float64 = -180
	maxLongitude float64 = 180

	bitsPerChar = 5
	alphabet    = "0123456789bcdefghjkmnpqrstuvwxyz"
)

const (
	// Global ~5000 km × 5000 km (Global Scale)
	Global Precision = iota + 1

	// Country ~1250 km × 625 km (Country Level)
	Country

	// State ~156 km × 156 km (State or Large Region Level)
	State

	// Region ~39 km × 19.5 km (Region Level)
	Region

	// City ~4.9 km × 4.9 km (City or District Level)
	City

	// Street ~1.2 km × 0.61 km (Street or Neighborhood Level)
	Street

	// Building ~152 m × 152 m (Building Level)
	Building

	// Block ~38 m × 19 m (Block or Complex Level)
	Block

	// House ~4.8 m × 4.8 m (House Level)
	House

	// Room ~1.2 m × 0.6 m (Room-Level Precision)
	Room

	// Point ~15 cm × 15 cm (Point-Level Precision)
	Point

	// SubPoint ~1.9 cm × 1.9 cm (SubPoint-Level Precision)
	SubPoint
)

const (
	// N represents the direction North, typically used in directional or enumerated constants.
	N Direction = iota

	// NE represents the northeast direction in the set of directional constants.
	NE

	// E represents the east direction in a set of directional constants.
	E

	// SE represents the southeast direction in the directional constants enumeration.
	SE

	// S represents the south direction in the directional constants enumeration.
	S

	// SW represents the southwest direction in the constant iota enumeration for cardinal and intercardinal directions.
	SW

	// W represents the westward direction in a compass enumeration, typically used for orientation-based logic.
	W

	// NW represents the northwest direction in the set of cardinal and intercardinal directions.
	NW
)

var (
	// ErrLatitudeOutOfRange is returned when a latitude value is out of the valid range (-90 to 90 degrees).
	ErrLatitudeOutOfRange = errors.New("latitude out of range")

	// ErrLongitudeOutOfRange is returned when a longitude value is out of the valid range (-180 to 180 degrees).
	ErrLongitudeOutOfRange = errors.New("longitude out of range")

	// ErrPrecisionOutOfRange is returned when a precision value is outside the valid range.
	ErrPrecisionOutOfRange = errors.New("precision out of range")

	// ErrInvalidHashLength is returned when the length of a GeoHash string is outside the valid range (1 to 12).
	ErrInvalidHashLength = errors.New("invalid hash length")

	// ErrInvalidHashFormat is returned when a GeoHash string contains characters outside the allowed Base32 alphabet.
	ErrInvalidHashFormat = errors.New("invalid hash format")

	// ErrDirectionOutOfRange is returned when a direction value is outside the acceptable range of valid directions.
	ErrDirectionOutOfRange = errors.New("direction out of range")
)

// alphabetMap maps byte characters to corresponding index.
var alphabetMap = map[int32]uint64{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'b': 10,
	'c': 11,
	'd': 12,
	'e': 13,
	'f': 14,
	'g': 15,
	'h': 16,
	'j': 17,
	'k': 18,
	'm': 19,
	'n': 20,
	'p': 21,
	'q': 22,
	'r': 23,
	's': 24,
	't': 25,
	'u': 26,
	'v': 27,
	'w': 28,
	'x': 29,
	'y': 30,
	'z': 31,
}

type (
	// Precision GeoHash precision Levels
	Precision int

	// Direction represents cardinal or intercardinal direction
	Direction int

	// BBox defines geographical boundaries using minimum and maximum latitude and longitude values.
	BBox struct {
		MinLatitude  float64
		MaxLatitude  float64
		MinLongitude float64
		MaxLongitude float64
	}
)

// Encode generates a GeoHash string for the given latitude, longitude, and precision.
// Returns an error if the latitude, longitude, or precision is out of the valid range.
func Encode(latitude, longitude float64, precision Precision) (string, error) {
	if latitude < minLatitude || latitude > maxLatitude {
		return "", ErrLatitudeOutOfRange
	}
	if longitude < minLongitude || longitude > maxLongitude {
		return "", ErrLongitudeOutOfRange
	}
	if precision < Global || precision > SubPoint {
		return "", ErrPrecisionOutOfRange
	}

	lngBitset := encodeCoordinateBitset(minLongitude, maxLongitude, longitude, true, precision)
	latBitset := encodeCoordinateBitset(minLatitude, maxLatitude, latitude, false, precision)
	bitset := interlaceBitsets(latBitset, lngBitset, precision)

	return encodeToBase32(bitset, precision), nil
}

// MustEncode generates a GeoHash string for given latitude, longitude, and precision or panics if an error occurs.
func MustEncode(latitude, longitude float64, precision Precision) string {
	hash, err := Encode(latitude, longitude, precision)
	if err != nil {
		panic(err)
	}
	return hash
}

// Decode takes a GeoHash string and returns its decoded latitude and longitude.
// Returns an error if the GeoHash string is invalid or cannot be decoded.
func Decode(hash string) (latitude, longitude float64, err error) {
	if len(hash) < int(Global) || len(hash) > int(SubPoint) {
		return 0, 0, ErrInvalidHashLength
	}

	bitset, precision, err := decodeFromBase32(hash)
	if err != nil {
		return 0, 0, err
	}

	latBitset, lngBitset := splitBitset(bitset, precision)
	_, _, latitude = decodeCoordinateBitset(latBitset, false, precision)
	_, _, longitude = decodeCoordinateBitset(lngBitset, true, precision)

	return latitude, longitude, nil
}

// MustDecode decodes a GeoHash string into latitude and longitude or panics if an error occurs.
func MustDecode(hash string) (latitude, longitude float64) {
	lat, lon, err := Decode(hash)
	if err != nil {
		panic(err)
	}
	return lat, lon
}

// DecodeBBox decodes a GeoHash string into its center coordinates with relative bounding box (BBox).
// Returns an error if the GeoHash string is invalid or cannot be decoded.
func DecodeBBox(hash string) (latitude float64, longitude float64, bbox BBox, err error) {
	if len(hash) < int(Global) || len(hash) > int(SubPoint) {
		return 0, 0, BBox{}, ErrInvalidHashLength
	}

	bitset, precision, err := decodeFromBase32(hash)
	if err != nil {
		return 0, 0, BBox{}, err
	}

	latBitset, lngBitset := splitBitset(bitset, precision)
	minLat, maxLat, latitude := decodeCoordinateBitset(latBitset, false, precision)
	minLng, maxLng, longitude := decodeCoordinateBitset(lngBitset, true, precision)

	return latitude, longitude, BBox{
		MinLatitude:  minLat,
		MaxLatitude:  maxLat,
		MinLongitude: minLng,
		MaxLongitude: maxLng,
	}, nil
}

// MustDecodeBBox decodes a GeoHash string to coordinates and bounding box or panics if an error occurs.
func MustDecodeBBox(hash string) (latitude float64, longitude float64, bbox BBox) {
	lat, lon, bbox, err := DecodeBBox(hash)
	if err != nil {
		panic(err)
	}
	return lat, lon, bbox
}

// Neighbor returns the neighbor of a given GeoHash in one enumerated Direction.
// Returns an error if the input is invalid.
func Neighbor(hash string, direction Direction) (string, error) {
	latitude, longitude, err := Decode(hash)
	if err != nil {
		return "", err
	}

	if direction < N || direction > NW {
		return "", ErrDirectionOutOfRange
	}

	totalBits := uint(len(hash)) * bitsPerChar
	latitudeBits := totalBits / 2
	longitudeBits := totalBits / 2
	if totalBits%2 != 0 {
		longitudeBits++
	}

	deltaLat := 180.0 / float64(uint(1)<<latitudeBits)
	deltaLng := 360.0 / float64(uint(1)<<longitudeBits)

	// Direction deltas in the order: N, NE, E, SE, S, SW, W, NW.
	directionDeltas := []struct {
		lat float64
		lon float64
	}{
		{+deltaLat, 0},         // N
		{+deltaLat, +deltaLng}, // NE
		{0, +deltaLng},         // E
		{-deltaLat, +deltaLng}, // SE
		{-deltaLat, 0},         // S
		{-deltaLat, -deltaLng}, // SW
		{0, -deltaLng},         // W
		{+deltaLat, -deltaLng}, // NW
	}

	d := directionDeltas[direction]
	lat, lng := wrapCoordinates(latitude+d.lat, longitude+d.lon)

	neighborHash := MustEncode(lat, lng, Precision(len(hash)))
	return neighborHash, nil
}

// MustNeighbor returns the neighboring GeoHash in a given direction, panicking if an error occurs during lookup.
func MustNeighbor(hash string, direction Direction) string {
	n, err := Neighbor(hash, direction)
	if err != nil {
		panic(err)
	}
	return n
}

// Neighbors returns the eight neighboring GeoHash values for the given GeoHash string.
// The returned slice is indexed by the Direction enumeration in the order: N, NE, E,
// SE, S, SW, W, NW.
// Returns an error if the input is invalid.
func Neighbors(hash string) ([]string, error) {
	results := make([]string, 8)
	for dir := N; dir <= NW; dir++ {
		n, err := Neighbor(hash, dir)
		if err != nil {
			return nil, err
		}
		results[dir] = n
	}
	return results, nil
}

// MustNeighbors returns the eight neighboring GeoHash values for the given hash. It panics if an error occurs.
func MustNeighbors(hash string) []string {
	neighbors, err := Neighbors(hash)
	if err != nil {
		panic(err)
	}
	return neighbors
}

// encodeCoordinateBitset generates a bitset for a coordinate using binary partitioning within given bounds and precision.
func encodeCoordinateBitset(leftBound, rightBound, value float64, isLongitude bool, precision Precision) uint64 {
	totalBits := int(precision) * bitsPerChar
	cycles := totalBits / 2
	if isLongitude && totalBits%2 != 0 {
		cycles++
	}

	var bitset uint64
	for i := 0; i < cycles; i++ {
		avg := (leftBound + rightBound) / 2.0

		bitset <<= 1

		if value >= avg {
			bitset |= 1
			leftBound = avg
			continue
		}

		rightBound = avg
	}

	return bitset
}

func decodeCoordinateBitset(bitset uint64, isLongitude bool, precision Precision) (max, min, center float64) {
	leftBound := minLatitude
	rightBound := maxLatitude

	totalBits := int(precision) * bitsPerChar
	bitsToProcess := totalBits / 2
	if isLongitude {
		leftBound = minLongitude
		rightBound = maxLongitude

		if totalBits%2 != 0 {
			bitsToProcess++
		}
	}

	var mid float64
	for i := 0; i < bitsToProcess; i++ {
		msb := (bitset >> (bitsToProcess - 1 - i)) & 1
		mid = (leftBound + rightBound) / 2.0

		if msb == 1 {
			leftBound = mid
		} else {
			rightBound = mid
		}
	}

	return leftBound, rightBound, (leftBound + rightBound) / 2.0
}

// interlaceBitsets interlaces latitude and longitude bitsets at the specified precision to generate a combined GeoHash bitset.
func interlaceBitsets(latBitset, lngBitset uint64, precision Precision) uint64 {
	totalBits := int(precision) * bitsPerChar
	lngBitCount := (totalBits + 1) / 2
	latBitCount := totalBits / 2

	var bitset uint64
	for i := 0; i < totalBits; i++ {
		bitset <<= 1

		if i%2 == 0 {
			bitset |= (lngBitset >> (lngBitCount - 1 - i/2)) & 1
			continue
		}

		bitset |= (latBitset >> (latBitCount - 1 - i/2)) & 1
	}

	return bitset
}

// splitBitset splits a GeoHash bitset into separate longitude and latitude bitsets based on the provided precision.
func splitBitset(bitset uint64, precision Precision) (latBitset uint64, lngBitset uint64) {
	totalBits := int(precision) * bitsPerChar
	bitset <<= 64 - totalBits

	for i := 0; i < totalBits; i++ {
		msb := (bitset >> 63) & 1

		if i%2 == 0 {
			lngBitset = (lngBitset << 1) | msb
		} else {
			latBitset = (latBitset << 1) | msb
		}

		bitset <<= 1
	}

	return latBitset, lngBitset
}

// encodeToBase32 encodes a given bitset into a Base32 GeoHash string using the specified precision level.
func encodeToBase32(bitset uint64, precision Precision) string {
	const mask = 0x1F // 0b11111

	p := int(precision)
	shift := precision * bitsPerChar
	bitset <<= 64 - shift

	var buf [SubPoint]byte
	for i := 0; i < p; i++ {
		index := (bitset >> (64 - bitsPerChar)) & mask
		buf[i] = alphabet[index]
		bitset <<= bitsPerChar
	}

	return string(buf[:p])
}

// decodeFromBase32 decodes a Base32 GeoHash string into a bitset, its precision, or an error on invalid input.
func decodeFromBase32(hash string) (uint64, Precision, error) {
	var bitset uint64
	for _, char := range hash {
		index, ok := alphabetMap[char]
		if !ok {
			return 0, 0, ErrInvalidHashFormat
		}

		bitset <<= bitsPerChar
		bitset |= index
	}

	return bitset, Precision(len(hash)), nil
}

// wrapCoordinates adjusts latitude and longitude values to ensure they fall within their valid ranges: [-90, 90] and [-180, 180].
func wrapCoordinates(lat, lon float64) (float64, float64) {
	lat = math.Mod(lat+maxLatitude, 2*maxLatitude)
	if lat < 0 {
		lat += 2 * maxLatitude
	}
	lat -= maxLatitude

	lon = math.Mod(lon+maxLongitude, 2*maxLongitude)
	if lon < 0 {
		lon += 2 * maxLongitude
	}
	lon -= maxLongitude

	return lat, lon
}
