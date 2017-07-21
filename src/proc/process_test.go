package proc

import (
	"testing"
	"io/ioutil"
)

func TestProcess(t *testing.T) {
	FossilDir(Paras{"../../examples/input/", "../../examples/output", false, 16})
	outputs := Ls("../../examples/output", func(string) bool { return true }, false)

	var out []string
	for _, o := range outputs {
		content, err := ioutil.ReadFile(o.Path + o.Name)
		if err != nil {
			panic(err)
		}
		out = append(out, string(content))
	}

}

func stringToStringTest(apply func(string) string, test []string, result []string, t *testing.T) {
	for i, s := range test {
		r := apply(s)
		if r != result[i] {
			t.Errorf("Error: %s should be %s but got %s\n", s, result[i], r)
		}
	}
}

func TestRegularToSnake(t *testing.T) {
	test := []string{"aaa", "a.a", "AAA", "%%%A", "aaa_", "aaa_aaa", "_aaa", "aAa", "AaA"}
	result := []string{"aaa", "a_a", "a_a_a", "a", "aaa_", "aaa_aaa", "_aaa", "a_aa", "aa_a"}
	stringToStringTest(regularizeToSnakeCase, test, result, t)
}

func TestSnakeToCamel(t *testing.T) {
	test := []string{"aaa", "a_a", "a_a_a", "a", "aaa_", "aaa_aaa", "_aaa", "a_aa", "aa_a"}
	result := []string{"Aaa", "AA", "AAA", "A", "Aaa", "AaaAaa", "Aaa", "AAa", "AaA"}
	stringToStringTest(snakeToCamelCase, test, result, t)
}
