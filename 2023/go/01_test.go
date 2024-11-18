package advent2023

import (
	"testing"
)

func Test_Calibrate(t *testing.T) {
	params := GetInputParams{
		day:      "01",
		fileName: "sample_1",
	}
	_, input := GetInputByDay(params)
	calibrated := Calibrate(input)

	if calibrated != 142 {
		t.Errorf("Expected 142, got %d", calibrated)
	}
}

func Test_CalibrateWithWords(t *testing.T) {
	params := GetInputParams{
		day:      "01",
		fileName: "sample_2",
	}
	_, input := GetInputByDay(params)
	calibrated := CalibrateWithWords(input)

	if calibrated != 281 {
		t.Errorf("Expected 281, got %d", calibrated)
	}
}
