package converter

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		from    TempUnit
		to      TempUnit
		want    float64
		wantErr bool
	}{
		{"celsius to fahrenheit", 100, Celsius, Fahrenheit, 212, false},
		{"fahrenheit to celsius", 32, Fahrenheit, Celsius, 0, false},
		{"celsius to kelvin", 0, Celsius, Kelvin, 273.15, false},
		{"kelvin to celsius", 273.15, Kelvin, Celsius, 0, false},
		{"same unit", 42, Celsius, Celsius, 42, false},
		{"unsupported from unit", 0, "X", Celsius, 0, true},
		{"unsupported to unit", 0, Celsius, "X", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertTemperature(tt.value, tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ConvertTemperature() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !almostEqual(got, tt.want) {
				t.Errorf("ConvertTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertLength(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		from    LengthUnit
		to      LengthUnit
		want    float64
		wantErr bool
	}{
		{"km to m", 1, Kilometer, Meter, 1000, false},
		{"m to km", 1000, Meter, Kilometer, 1, false},
		{"mile to m", 1, Mile, Meter, 1609.344, false},
		{"foot to m", 1, Foot, Meter, 0.3048, false},
		{"unsupported from unit", 0, "X", Meter, 0, true},
		{"unsupported to unit", 0, Meter, "X", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertLength(tt.value, tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ConvertLength() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !almostEqual(got, tt.want) {
				t.Errorf("ConvertLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertWeight(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		from    WeightUnit
		to      WeightUnit
		want    float64
		wantErr bool
	}{
		{"kg to g", 1, Kilogram, Gram, 1000, false},
		{"g to kg", 1000, Gram, Kilogram, 1, false},
		{"lb to kg", 1, Pound, Kilogram, 0.45359237, false},
		{"unsupported from unit", 0, "X", Kilogram, 0, true},
		{"unsupported to unit", 0, Kilogram, "X", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertWeight(tt.value, tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ConvertWeight() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !almostEqual(got, tt.want) {
				t.Errorf("ConvertWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
