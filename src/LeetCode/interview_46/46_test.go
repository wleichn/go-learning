package interview_46

import "testing"

func TestTranslateNum(t *testing.T) {
	samples := []int{12258, 506}
	results := []int{5, 1}

	for index, sample := range samples {
		if translateNum(sample) != results[index] {
			t.Errorf("Error result in sample %d", index)
		}
	}
}
