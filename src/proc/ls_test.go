package proc

import (
	"testing"
)

func TestLs(t *testing.T) {
	infos := Ls("../../examples/input", func(string) bool { return true }, true)
	expected := 7
	if !(len(infos) == expected) {
		t.Errorf(`Ls("../../examples/input") get %d files, expected %d.`, len(infos), expected)
	}
}
