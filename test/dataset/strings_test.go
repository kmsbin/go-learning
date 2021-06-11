package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - expeted index (%d) found (%d) "

func TestIndex(t *testing.T) {
	dataSet := []struct {
		text     string
		part     string
		expected int
	}{
		{"kaulindo é vida", "kaulindo", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"kaulindo", "n", 5},
	}
	for _, textTest := range dataSet {
		t.Logf("Massa: %v", textTest)
		now := strings.Index(textTest.text, textTest.part)

		if now != textTest.expected {
			t.Errorf(msgIndex, textTest.text, textTest.part, textTest.expected, now)
		}
	}
}
