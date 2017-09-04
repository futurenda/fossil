package fossil

import (
	"io/ioutil"
	"testing"
)

func TestProcess(t *testing.T) {
	FossilDir(FossilParas{"../examples/input/", "../examples/output", false, 16, ""})
	outputs := ls("../examples/output", func(string) bool { return true }, false)

	var out []string
	for _, o := range outputs {
		content, err := ioutil.ReadFile(o.Path + o.Name)
		if err != nil {
			panic(err)
		}
		out = append(out, string(content))
	}

}
