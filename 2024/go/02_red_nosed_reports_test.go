package advent2024

import "testing"

func Test_SafetyReport(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "02",
		fileName: "sample_1",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := SafetyReport(input)

	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}

func Test_SafetyReport_FullInput(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "02",
		fileName: "full_input",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := SafetyReport(input)

	if result != 516 {
		t.Errorf("expected 516, got %d", result)
	}
}

func Test_SafetyReportWithDampener(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "02",
		fileName: "sample_1",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := SafetyReportWithDampener(input)

	if result != 1 {
		t.Errorf("expected 4, got %d", result)
	}
}
