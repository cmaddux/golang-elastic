package encoding

import "testing"

func TestEncodingEmpty(t *testing.T) {
	str := ""

	encoded := RunLength(str)
	if encoded != nil {
		t.Error("Expected [] got", encoded)
	}
}

func TestEncodingGeneral(t *testing.T) {
	str := "aaaaabbbbc"
	expect := [][2]interface{}{
		[2]interface{}{ 'a', 5 },
		[2]interface{}{ 'b', 4 },
		[2]interface{}{ 'c', 1 },
	}

	encoded := RunLength(str)
	for idx, item := range expect {
		if encoded[idx] != item {
			t.Error("Expected", item, "got", encoded[idx])
		}
	}

}
