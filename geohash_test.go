package geohash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Acceptable difference up to 6 decimal places (±1 meter)
const tolerance = 1e-6

func TestEncode(t *testing.T) {
	tests := []struct {
		name      string
		latitude  float64
		longitude float64
		precision Precision
		want      string
		wantErr   error
	}{
		{
			name:      "Latitude out of range - too low",
			latitude:  -91.0,
			longitude: 0.0,
			precision: 7,
			want:      "",
			wantErr:   ErrLatitudeOutOfRange,
		},
		{
			name:      "Latitude out of range - too high",
			latitude:  91.0,
			longitude: 0.0,
			precision: 7,
			want:      "",
			wantErr:   ErrLatitudeOutOfRange,
		},
		{
			name:      "Longitude out of range - too low",
			latitude:  0.0,
			longitude: -181.0,
			precision: 7,
			want:      "",
			wantErr:   ErrLongitudeOutOfRange,
		},
		{
			name:      "Longitude out of range - too high",
			latitude:  0.0,
			longitude: 181.0,
			precision: 7,
			want:      "",
			wantErr:   ErrLongitudeOutOfRange,
		},
		{
			name:      "Invalid precision - too low",
			latitude:  0.0,
			longitude: 0.0,
			precision: 0,
			want:      "",
			wantErr:   ErrPrecisionOutOfRange,
		},
		{
			name:      "Invalid precision - too high",
			latitude:  0.0,
			longitude: 0.0,
			precision: 13,
			want:      "",
			wantErr:   ErrPrecisionOutOfRange,
		},
		{
			name:      "Edge case - min latitude and longitude",
			latitude:  -90.0,
			longitude: -180.0,
			precision: 5,
			want:      "00000",
			wantErr:   nil,
		},
		{
			name:      "Edge case - max latitude and longitude",
			latitude:  90.0,
			longitude: 180.0,
			precision: 5,
			want:      "zzzzz",
			wantErr:   nil,
		},
		{
			name:      "Edge case - zero latitude and zero longitude",
			latitude:  0.0,
			longitude: 0.0,
			precision: 5,
			want:      "s0000",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Global",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Global,
			want:      "9",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Country",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Country,
			want:      "9q",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - State",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: State,
			want:      "9q8",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Region",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Region,
			want:      "9q8y",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - City",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: City,
			want:      "9q8yy",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Street",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Street,
			want:      "9q8yyk",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Building",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Building,
			want:      "9q8yyk8",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Block",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Block,
			want:      "9q8yyk8y",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - House",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: House,
			want:      "9q8yyk8yt",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Room",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Room,
			want:      "9q8yyk8ytp",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - Point",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: Point,
			want:      "9q8yyk8ytpx",
			wantErr:   nil,
		},
		{
			name:      "Valid Precision - SubPoint",
			latitude:  37.7749,
			longitude: -122.4194,
			precision: SubPoint,
			want:      "9q8yyk8ytpxr",
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.latitude, tt.longitude, tt.precision)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMustEncode(t *testing.T) {
	type args struct {
		latitude  float64
		longitude float64
		precision Precision
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{
			name: "Latitude out of range - too low",
			args: args{
				latitude:  -91.0,
				longitude: 0.0,
				precision: 7,
			},
			wantPanic: true,
		},
		{
			name: "Latitude out of range - too high",
			args: args{
				latitude:  91.0,
				longitude: 0.0,
				precision: 7,
			},
			wantPanic: true,
		},
		{
			name: "Longitude out of range - too low",
			args: args{
				latitude:  0.0,
				longitude: -181.0,
				precision: 7,
			},
			wantPanic: true,
		},
		{
			name: "Longitude out of range - too high",
			args: args{
				latitude:  0.0,
				longitude: 181.0,
				precision: 7,
			},
			wantPanic: true,
		},
		{
			name: "Invalid precision - too low",
			args: args{
				latitude:  0.0,
				longitude: 0.0,
				precision: 0,
			},
			wantPanic: true,
		},
		{
			name: "Invalid precision - too high",
			args: args{
				latitude:  0.0,
				longitude: 0.0,
				precision: 13,
			},
			wantPanic: true,
		},
		{
			name: "Valid Precision - City",
			args: args{
				latitude:  37.7749,
				longitude: -122.4194,
				precision: City,
			},
			want:      "9q8yy",
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { MustEncode(tt.args.latitude, tt.args.longitude, tt.args.precision) })
			} else {
				assert.Equalf(t, tt.want, MustEncode(tt.args.latitude, tt.args.longitude, tt.args.precision), "MustEncode(%v, %v, %v)", tt.args.latitude, tt.args.longitude, tt.args.precision)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name          string
		args          args
		wantLatitude  float64
		wantLongitude float64
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "Invalid hash - too short",
			args: args{
				hash: "",
			},
			wantLatitude:  0,
			wantLongitude: 0,
			wantErr:       assert.Error,
		},
		{
			name: "Invalid hash - too long",
			args: args{
				hash: "9q8yyk8ytpxrs",
			},
			wantLatitude:  0,
			wantLongitude: 0,
			wantErr:       assert.Error,
		},
		{
			name: "Invalid hash - contains invalid characters",
			args: args{
				hash: "9q8yy!",
			},
			wantLatitude:  0,
			wantLongitude: 0,
			wantErr:       assert.Error,
		},
		{
			name: "GeoHash at min latitude and longitude",
			args: args{
				hash: "00000",
			},
			wantLatitude:  -89.978027,
			wantLongitude: -179.978027,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash at max latitude and longitude",
			args: args{
				hash: "zzzzz",
			},
			wantLatitude:  89.978027,
			wantLongitude: 179.978027,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash at zero latitude and longitude",
			args: args{
				hash: "s0000",
			},
			wantLatitude:  0.021972,
			wantLongitude: 0.021972,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Global",
			args: args{
				hash: "9",
			},
			wantLatitude:  22.5,
			wantLongitude: -112.5,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Country",
			args: args{
				hash: "9q",
			},
			wantLatitude:  36.5625,
			wantLongitude: -118.125,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision State",
			args: args{
				hash: "9q8",
			},
			wantLatitude:  37.265625,
			wantLongitude: -123.046875,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Region",
			args: args{
				hash: "9q8y",
			},
			wantLatitude:  37.705078125,
			wantLongitude: -122.51953125,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision City",
			args: args{
				hash: "9q8yy",
			},
			wantLatitude:  37.770996,
			wantLongitude: -122.409667,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Street",
			args: args{
				hash: "9q8yyk",
			},
			wantLatitude:  37.773742,
			wantLongitude: -122.415162,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Building",
			args: args{
				hash: "9q8yyk8",
			},
			wantLatitude:  37.774429,
			wantLongitude: -122.419967,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Block",
			args: args{
				hash: "9q8yyk8y",
			},
			wantLatitude:  37.774858,
			wantLongitude: -122.419452,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision House",
			args: args{
				hash: "9q8yyk8yp",
			},
			wantLatitude:  37.774794,
			wantLongitude: -122.419302,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Room",
			args: args{
				hash: "9q8yyk8ypd",
			},
			wantLatitude:  37.774786,
			wantLongitude: -122.419297,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision Point",
			args: args{
				hash: "9q8yyk8ypd2",
			},
			wantLatitude:  37.774785,
			wantLongitude: -122.419301,
			wantErr:       assert.NoError,
		},
		{
			name: "GeoHash - Precision SubPoint",
			args: args{
				hash: "9q8yyk8ypd23",
			},
			wantLatitude:  37.7747849,
			wantLongitude: -122.419301,
			wantErr:       assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLatitude, gotLongitude, err := Decode(tt.args.hash)
			if !tt.wantErr(t, err, fmt.Sprintf("Decode(%v)", tt.args.hash)) {
				return
			}
			assert.InDeltaf(t, tt.wantLatitude, gotLatitude, tolerance, "Decode(%v)", tt.args.hash)
			assert.InDeltaf(t, tt.wantLongitude, gotLongitude, tolerance, "Decode(%v)", tt.args.hash)
		})
	}
}

func TestMustDecode(t *testing.T) {
	tests := []struct {
		name      string
		hash      string
		wantLat   float64
		wantLon   float64
		wantPanic bool
	}{
		{
			name:      "Invalid GeoHash - Too Short",
			hash:      "",
			wantLat:   0,
			wantLon:   0,
			wantPanic: true,
		},
		{
			name:      "Invalid GeoHash - Too Long",
			hash:      "9q8yyk8ytpxrs",
			wantLat:   0,
			wantLon:   0,
			wantPanic: true,
		},
		{
			name:      "Invalid GeoHash - Invalid Characters",
			hash:      "9q8yy!",
			wantLat:   0,
			wantLon:   0,
			wantPanic: true,
		},
		{
			name:      "Valid GeoHash - Precision Global",
			hash:      "9",
			wantLat:   22.5,
			wantLon:   -112.5,
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { MustDecode(tt.hash) })
			} else {
				lat, lon := MustDecode(tt.hash)
				assert.InDeltaf(t, tt.wantLat, lat, tolerance, "MustDecode(%v)", tt.hash)
				assert.InDeltaf(t, tt.wantLon, lon, tolerance, "MustDecode(%v)", tt.hash)
			}
		})
	}
}

func TestDecodeBBox(t *testing.T) {
	tests := []struct {
		name     string
		hash     string
		wantLat  float64
		wantLon  float64
		wantBBox BBox
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:    "Invalid GeoHash - Too Short",
			hash:    "",
			wantLat: 0, wantLon: 0,
			wantBBox: BBox{},
			wantErr:  assert.Error,
		},
		{
			name:    "Invalid GeoHash - Too Long",
			hash:    "9q8yyk8ytpxrs",
			wantLat: 0, wantLon: 0,
			wantBBox: BBox{},
			wantErr:  assert.Error,
		},
		{
			name:    "Invalid GeoHash - Invalid Characters",
			hash:    "9q8yy!",
			wantLat: 0, wantLon: 0,
			wantBBox: BBox{},
			wantErr:  assert.Error,
		},
		{
			name:    "Valid GeoHash - Precision Global",
			hash:    "9",
			wantLat: 22.5, wantLon: -112.5,
			wantBBox: BBox{
				MinLatitude:  0,
				MaxLatitude:  45,
				MinLongitude: -135,
				MaxLongitude: -90,
			},
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - Precision City",
			hash:    "9q8yy",
			wantLat: 37.77099609375, wantLon: -122.40966796875,
			wantBBox: BBox{
				MinLatitude:  37.7490234375,
				MaxLatitude:  37.79296875,
				MinLongitude: -122.431640625,
				MaxLongitude: -122.3876953125,
			},
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - Precision House",
			hash:    "9q8yyk8yp",
			wantLat: 37.77479410171509, wantLon: -122.4193024635315,
			wantBBox: BBox{
				MinLatitude:  37.77477264404297,
				MaxLatitude:  37.77481555938721,
				MinLongitude: -122.41932392120361,
				MaxLongitude: -122.41928100585938,
			},
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - Precision SubPoint",
			hash:    "9q8yyk8ypd23",
			wantLat: 37.774784, wantLon: -122.419302,
			wantBBox: BBox{
				MinLatitude:  37.774785,
				MaxLatitude:  37.774785,
				MinLongitude: -122.419302,
				MaxLongitude: -122.419301,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLat, gotLon, gotBBox, err := DecodeBBox(tt.hash)
			if !tt.wantErr(t, err, fmt.Sprintf("DecodeBBox(%v)", tt.hash)) {
				return
			}

			assert.InDeltaf(t, tt.wantLat, gotLat, tolerance, "Latitude DecodeBBox(%v)", tt.hash)
			assert.InDeltaf(t, tt.wantLon, gotLon, tolerance, "Longitude DecodeBBox(%v)", tt.hash)
			assert.InDeltaf(t, tt.wantBBox.MinLatitude, gotBBox.MinLatitude, tolerance, "MinLatitude DecodeBBox(%v)", tt.hash)
			assert.InDeltaf(t, tt.wantBBox.MaxLatitude, gotBBox.MaxLatitude, tolerance, "MaxLatitude DecodeBBox(%v)", tt.hash)
			assert.InDeltaf(t, tt.wantBBox.MinLongitude, gotBBox.MinLongitude, tolerance, "MinLongitude DecodeBBox(%v)", tt.hash)
			assert.InDeltaf(t, tt.wantBBox.MaxLongitude, gotBBox.MaxLongitude, tolerance, "MaxLongitude DecodeBBox(%v)", tt.hash)
		})
	}
}

func TestMustDecodeBBox(t *testing.T) {
	tests := []struct {
		name      string
		hash      string
		wantLat   float64
		wantLon   float64
		wantBBox  BBox
		wantPanic bool
	}{
		{
			name:      "Invalid GeoHash - Too Short",
			hash:      "",
			wantPanic: true,
		},
		{
			name:      "Invalid GeoHash - Too Long",
			hash:      "9q8yyk8ytpxrs",
			wantPanic: true,
		},
		{
			name:      "Invalid GeoHash - Invalid Characters",
			hash:      "9q8yy!",
			wantPanic: true,
		},
		{
			name:    "Valid GeoHash - Precision Global",
			hash:    "9",
			wantLat: 22.5, wantLon: -112.5,
			wantBBox: BBox{
				MinLatitude:  0,
				MaxLatitude:  45,
				MinLongitude: -135,
				MaxLongitude: -90,
			},
			wantPanic: false,
		},
		{
			name:    "Valid GeoHash - Precision City",
			hash:    "9q8yy",
			wantLat: 37.77099609375, wantLon: -122.40966796875,
			wantBBox: BBox{
				MinLatitude:  37.7490234375,
				MaxLatitude:  37.79296875,
				MinLongitude: -122.431640625,
				MaxLongitude: -122.3876953125,
			},
			wantPanic: false,
		},
		{
			name:    "Valid GeoHash - Precision House",
			hash:    "9q8yyk8yp",
			wantLat: 37.774794, wantLon: -122.419302,
			wantBBox: BBox{
				MinLatitude:  37.774772,
				MaxLatitude:  37.774815,
				MinLongitude: -122.419323,
				MaxLongitude: -122.419281,
			},
			wantPanic: false,
		},
		{
			name:    "Valid GeoHash - Precision SubPoint",
			hash:    "9q8yyk8ypd23",
			wantLat: 37.774784, wantLon: -122.419301,
			wantBBox: BBox{
				MinLatitude:  37.774785,
				MaxLatitude:  37.774785,
				MinLongitude: -122.419302,
				MaxLongitude: -122.4193009,
			},
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_, _, _ = MustDecodeBBox(tt.hash)
				}, "Expected panic for MustDecodeBBox(%v)", tt.hash)
			} else {
				assert.NotPanics(t, func() {
					gotLat, gotLon, gotBBox := MustDecodeBBox(tt.hash)
					assert.InDeltaf(t, tt.wantLat, gotLat, tolerance, "Latitude MustDecodeBBox(%v)", tt.hash)
					assert.InDeltaf(t, tt.wantLon, gotLon, tolerance, "Longitude MustDecodeBBox(%v)", tt.hash)
					assert.InDeltaf(t, tt.wantBBox.MinLatitude, gotBBox.MinLatitude, tolerance, "MinLatitude MustDecodeBBox(%v)", tt.hash)
					assert.InDeltaf(t, tt.wantBBox.MaxLatitude, gotBBox.MaxLatitude, tolerance, "MaxLatitude MustDecodeBBox(%v)", tt.hash)
					assert.InDeltaf(t, tt.wantBBox.MinLongitude, gotBBox.MinLongitude, tolerance, "MinLongitude MustDecodeBBox(%v)", tt.hash)
					assert.InDeltaf(t, tt.wantBBox.MaxLongitude, gotBBox.MaxLongitude, tolerance, "MaxLongitude MustDecodeBBox(%v)", tt.hash)
				})
			}
		})
	}
}

func TestNeighbourhoods(t *testing.T) {
	tests := []struct {
		name      string
		hash      string
		want      []string
		wantError assert.ErrorAssertionFunc
	}{
		{
			name:      "Valid GeoHash - Precision Global",
			hash:      "9",
			want:      []string{"c", "f", "d", "6", "3", "2", "8", "b"},
			wantError: assert.NoError,
		},
		{
			name:      "Valid GeoHash - Precision City",
			hash:      "9q8yy",
			want:      []string{"9q8zn", "9q8zp", "9q8yz", "9q8yx", "9q8yw", "9q8yt", "9q8yv", "9q8zj"},
			wantError: assert.NoError,
		},
		{
			name:      "Valid GeoHash - Precision Block",
			hash:      "9q8yyk8y",
			want:      []string{"9q8yyk8z", "9q8yyk9p", "9q8yyk9n", "9q8yyk9j", "9q8yyk8v", "9q8yyk8t", "9q8yyk8w", "9q8yyk8x"},
			wantError: assert.NoError,
		},
		{
			name:      "Invalid GeoHash - Too Short",
			hash:      "",
			want:      nil,
			wantError: assert.Error,
		},
		{
			name:      "Invalid GeoHash - Too Long",
			hash:      "9q8yyk8ytpxrs",
			want:      nil,
			wantError: assert.Error,
		},
		{
			name:      "Invalid GeoHash - Contains Invalid Character",
			hash:      "9q8yy!",
			want:      nil,
			wantError: assert.Error,
		},
		{
			name:      "Extremes zzzzz - At limit of latitude and longitude",
			hash:      "zzzzz",
			want:      []string{"pbpbp", "00000", "bpbpb", "bpbp8", "zzzzx", "zzzzw", "zzzzy", "pbpbn"},
			wantError: assert.NoError,
		},
		{
			name:      "Extremes 0 - At limit of latitude and longitude",
			hash:      "0",
			want:      []string{"2", "3", "1", "c", "b", "z", "p", "r"},
			wantError: assert.NoError,
		},
		{
			name:      "Extremes z - At limit of latitude and longitude",
			hash:      "z",
			want:      []string{"p", "0", "b", "8", "x", "w", "y", "n"},
			wantError: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Neighbors(tt.hash)
			if !tt.wantError(t, err, fmt.Sprintf("Neighbors(%v)", tt.hash)) {
				return
			}
			if tt.want != nil {
				assert.ElementsMatch(t, tt.want, got, "Neighbors(%v)", tt.hash)
			}
		})
	}
}

func TestMustNeighbors(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name        string
		args        args
		want        []string
		expectPanic bool
	}{
		{
			name:        "Valid GeoHash - Precision Global",
			args:        args{hash: "9"},
			want:        []string{"c", "f", "d", "6", "3", "2", "8", "b"},
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - Precision City",
			args:        args{hash: "9q8yy"},
			want:        []string{"9q8zn", "9q8zp", "9q8yz", "9q8yx", "9q8yw", "9q8yt", "9q8yv", "9q8zj"},
			expectPanic: false,
		},
		{
			name:        "Invalid GeoHash - Empty",
			args:        args{hash: ""},
			want:        nil,
			expectPanic: true,
		},
		{
			name:        "Invalid GeoHash - Too Long",
			args:        args{hash: "9q8yyk8ytpxrs"},
			want:        nil,
			expectPanic: true,
		},
		{
			name:        "Invalid GeoHash - Invalid Characters",
			args:        args{hash: "9q8yy!"},
			want:        nil,
			expectPanic: true,
		},
		{
			name:        "Extremes z - At limit of latitude and longitude",
			args:        args{hash: "z"},
			want:        []string{"p", "0", "b", "8", "x", "w", "y", "n"},
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				assert.Panics(t, func() {
					_ = MustNeighbors(tt.args.hash)
				}, "Expected panic for MustNeighbors(%v)", tt.args.hash)
			} else {
				assert.NotPanics(t, func() {
					got := MustNeighbors(tt.args.hash)
					assert.Equalf(t, tt.want, got, "MustNeighbors(%v)", tt.args.hash)
				})
			}
		})
	}
}

func TestNeighbor(t *testing.T) {
	type args struct {
		hash      string
		direction Direction
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Valid GeoHash - North Neighbor",
			args:    args{hash: "9q8yy", direction: N},
			want:    "9q8zn",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - North-East Neighbor",
			args:    args{hash: "9q8yy", direction: NE},
			want:    "9q8zp",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - East Neighbor",
			args:    args{hash: "9q8yy", direction: E},
			want:    "9q8yz",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - South-East Neighbor",
			args:    args{hash: "9q8yy", direction: SE},
			want:    "9q8yx",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - South Neighbor",
			args:    args{hash: "9q8yy", direction: S},
			want:    "9q8yw",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - South-West Neighbor",
			args:    args{hash: "9q8yy", direction: SW},
			want:    "9q8yt",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - West Neighbor",
			args:    args{hash: "9q8yy", direction: W},
			want:    "9q8yv",
			wantErr: assert.NoError,
		},
		{
			name:    "Valid GeoHash - North-West Neighbor",
			args:    args{hash: "9q8yy", direction: NW},
			want:    "9q8zj",
			wantErr: assert.NoError,
		},
		{
			name:    "Invalid GeoHash - Empty",
			args:    args{hash: "", direction: N},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name:    "Invalid GeoHash - Too Long",
			args:    args{hash: "9q8yyk8ytpxrs", direction: N},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name:    "Invalid GeoHash - Invalid Characters",
			args:    args{hash: "9q8yy!", direction: N},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name:    "Invalid Direction - Out of Range",
			args:    args{hash: "9q8yy", direction: Direction(8)},
			want:    "",
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Neighbor(tt.args.hash, tt.args.direction)
			if !tt.wantErr(t, err, fmt.Sprintf("Neighbor(%v, %v)", tt.args.hash, tt.args.direction)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Neighbor(%v, %v)", tt.args.hash, tt.args.direction)
		})
	}
}

func TestMustNeighbor(t *testing.T) {
	type args struct {
		hash      string
		direction Direction
	}
	tests := []struct {
		name        string
		args        args
		want        string
		expectPanic bool
	}{
		{
			name:        "Valid GeoHash - North Neighbor",
			args:        args{hash: "9q8yy", direction: N},
			want:        "9q8zn",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - North-East Neighbor",
			args:        args{hash: "9q8yy", direction: NE},
			want:        "9q8zp",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - East Neighbor",
			args:        args{hash: "9q8yy", direction: E},
			want:        "9q8yz",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - South-East Neighbor",
			args:        args{hash: "9q8yy", direction: SE},
			want:        "9q8yx",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - South Neighbor",
			args:        args{hash: "9q8yy", direction: S},
			want:        "9q8yw",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - South-West Neighbor",
			args:        args{hash: "9q8yy", direction: SW},
			want:        "9q8yt",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - West Neighbor",
			args:        args{hash: "9q8yy", direction: W},
			want:        "9q8yv",
			expectPanic: false,
		},
		{
			name:        "Valid GeoHash - North-West Neighbor",
			args:        args{hash: "9q8yy", direction: NW},
			want:        "9q8zj",
			expectPanic: false,
		},
		{
			name:        "Invalid GeoHash - Empty",
			args:        args{hash: "", direction: N},
			want:        "",
			expectPanic: true,
		},
		{
			name:        "Invalid GeoHash - Too Long",
			args:        args{hash: "9q8yyk8ytpxrs", direction: N},
			want:        "",
			expectPanic: true,
		},
		{
			name:        "Invalid GeoHash - Invalid Characters",
			args:        args{hash: "9q8yy!", direction: N},
			want:        "",
			expectPanic: true,
		},
		{
			name:        "Invalid Direction - Out of Range",
			args:        args{hash: "9q8yy", direction: Direction(8)},
			want:        "",
			expectPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectPanic {
				assert.Panics(t, func() {
					_ = MustNeighbor(tt.args.hash, tt.args.direction)
				}, "Expected panic for MustNeighbor(%v, %v)", tt.args.hash, tt.args.direction)
			} else {
				assert.NotPanics(t, func() {
					got := MustNeighbor(tt.args.hash, tt.args.direction)
					assert.Equalf(t, tt.want, got, "MustNeighbor(%v, %v)", tt.args.hash, tt.args.direction)
				})
			}
		})
	}
}
