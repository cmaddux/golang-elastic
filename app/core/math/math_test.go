package math

import "testing"

func TestMinIntV1Min(t *testing.T) {
	v1 := 10
	v2 := 20

	res := MinInt(v1, v2)
	if res != v1 {
		t.Error("Expected", v1, "got", res)
	}
}

func TestMinIntV2Min(t *testing.T) {
	v1 := 20
	v2 := 10

	res := MinInt(v1, v2)
	if res != v2 {
		t.Error("Expected", v2, "got", res)
	}
}

func TestMinIntEqual(t *testing.T) {
	v1 := 20
	v2 := 20

	res := MinInt(v1, v2)
	if res != v1 {
		t.Error("Expected", v1, "got", res)
	}
}
