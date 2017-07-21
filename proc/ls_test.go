package proc

import (
	"testing"
	"path/filepath"
)

func TestLs(t *testing.T) {
	infos := Ls("../examples/input", func(string) bool { return true }, true)
	expected := 7
	if !(len(infos) == expected) {
		t.Errorf(`Ls("../examples/input") get %d files, expected %d.`, len(infos), expected)
	}
}

func TestLsWithFilter(t *testing.T) {
	sqlFileFilter := func(s string) bool {
		return filepath.Ext(s) == ".sql"
	}
	infos := Ls("../examples/input", sqlFileFilter, true)
	expected := 7
	if !(len(infos) == expected) {
		t.Errorf(`Ls("../examples/input") get %d files, expected %d.`, len(infos), expected)
	}
}
