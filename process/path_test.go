package process

import (
	"path/filepath"
	"testing"
)

func TestLs(t *testing.T) {
	infos := ls("../examples/input", func(string) bool { return true }, true)
	expected := 7
	if !(len(infos) == expected) {
		t.Errorf(`ls("../examples/input") get %d files, expected %d.`, len(infos), expected)
	}
}

func TestLsWithFilter(t *testing.T) {
	sqlFileFilter := func(s string) bool {
		return filepath.Ext(s) == ".sql"
	}
	infos := ls("../examples/input", sqlFileFilter, true)
	expected := 7
	if !(len(infos) == expected) {
		t.Errorf(`ls("../examples/input") get %d files, expected %d.`, len(infos), expected)
	}
}
