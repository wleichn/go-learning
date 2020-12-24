package _90

import "testing"

func Test990(t *testing.T) {
	samples := [][]string{{"a==b", "b!=a"}, {"b==a", "a==b"}, {"b==q", "g==l", "r==a", "t==o", "t==v", "i==b", "a==u", "e!=j", "d!=n", "y!=f", "h!=b", "u!=t", "o!=r", "k!=m"}}
	results := []bool{false, true, true}

	for index, sample := range samples {
		result := equationsPossible(sample)
		if result != results[index] {
			t.Errorf("Error for sample %d", index)
		}
	}

}
