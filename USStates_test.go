package StateAbbreviation

import "testing"

func test(t *testing.T, state string, expected string) {
	actual := GetAbbreviation(state)

	if actual == expected {
		t.Logf("%s ~ %s\t", state, expected)
	} else {
		t.Fatalf("%s != %s\t", state, expected)
	}
}

func TestGetAbbreviation(t *testing.T) {
	test(t, "Ala.", "AL")
	test(t, "Wyo.", "WY")
	test(t, "Wis..", "WI")
	test(t, "Tenn..", "TN")
	test(t, "Penna.", "PA")
	test(t, "Pa.", "PA")
	test(t, "PA", "PA")
	test(t, "Pennsylvania", "PA")
	test(t, "Ore.", "OR")
	test(t, "Ore.", "OR")
	test(t, "N.H.", "NH")

}
