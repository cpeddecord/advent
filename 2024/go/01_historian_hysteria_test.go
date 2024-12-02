package advent2024

import "testing"

func Test_TotalDistance(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "01",
		fileName: "sample_1",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := TotalDistance(input)

	if result != 11 {
		t.Errorf("expected 11, got %d", result)
	}
}

func Test_TotalDistance_FullInput(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "01",
		fileName: "full_input",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := TotalDistance(input)

	if result != 3508942 {
		t.Errorf("expected 3508942, got %d", result)
	}
}

func Test_ScoreSimilarity(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "01",
		fileName: "sample_1",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := ScoreSimilarity(input)

	if result != 31 {
		t.Errorf("expected 31, got %d", result)
	}
}

func Test_ScoreSimilarity_FullInput(t *testing.T) {
	err, input := GetInputByDay(GetInputParams{
		day:      "01",
		fileName: "full_input",
	})

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	result := ScoreSimilarity(input)

	if result != 26593248 {
		t.Errorf("expected 26593248, got %d", result)
	}
}
