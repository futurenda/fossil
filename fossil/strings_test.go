package fossil

import "testing"

func stringToStringTest(apply func(string) string, test []string, result []string, t *testing.T) {
	for i, s := range test {
		r := apply(s)
		if r != result[i] {
			t.Errorf("Error: %s should be %s but got %s\n", s, result[i], r)
		}
	}
}

func TestRegularToSnake(t *testing.T) {
	test := []string{"aaa", "a.a", "AAA", "%%%A", "aaa_", "aaa_aaa", "_aaa", "aAa", "AaA", "000A", "0A0", "0aA0", "中文"}
	result := []string{"aaa", "a_a", "a_a_a", "a", "aaa_", "aaa_aaa", "_aaa", "a_aa", "aa_a", "a", "a0", "a_a0", "中文"}
	stringToStringTest(regularizeToSnakeCase, test, result, t)
}

func TestSnakeToCamel(t *testing.T) {
	test := []string{"aaa", "a_a", "a_a_a", "a", "aaa_", "aaa_aaa", "_aaa", "a_aa", "aa_a"}
	result := []string{"Aaa", "AA", "AAA", "A", "Aaa", "AaaAaa", "Aaa", "AAa", "AaA"}
	stringToStringTest(snakeToCamelCase, test, result, t)
}
