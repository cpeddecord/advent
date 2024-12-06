package advent2024

import "testing"

func Test_FindXMASCount(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "04",
		fileName: "sample_1",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := FindXMASCount(input)

	if result != 18 {
		t.Errorf("expected 18, got %d", result)
	}
}

func Test_Ceres_FullInput(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "04",
		fileName: "full_input",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := FindXMASCount(input)

	if result != 2583 {
		t.Errorf("expected 2583, got %d", result)
	}
}
