package strings

import "testing"

func TestCountSpecialEmpty(t *testing.T) {
	var encoded [][2]interface{}
	res := CountSpecial(encoded)
	if res != 0 {
		t.Error("Expected 0 got", res)
	}
}

func TestCountSpecialType1(t *testing.T) {
	var encoded [][2]interface{}
	encoded = append(encoded, [2]interface{}{ 'a', 4 })
	res := CountSpecial(encoded)
	if res != 10 {
		t.Error("Expected 10 got", res)
	}
}

func TestCountSpecialType2(t *testing.T) {
	var encoded [][2]interface{}
	encoded = append(encoded, [2]interface{}{ 'a', 4 })
	encoded = append(encoded, [2]interface{}{ 'b', 1 })
	encoded = append(encoded, [2]interface{}{ 'a', 3 })
	res := CountSpecial(encoded)
	if res != 20 {
		t.Error("Expected 20 got", res)
	}
}

func TestCountSpecialNoType2(t *testing.T) {
	var encoded [][2]interface{}
	encoded = append(encoded, [2]interface{}{ 'a', 4 })
	encoded = append(encoded, [2]interface{}{ 'b', 1 })
	encoded = append(encoded, [2]interface{}{ 'c', 4 })
	res := CountSpecial(encoded)
	if res != 21 {
		t.Error("Expected 21 got", res)
	}
}
