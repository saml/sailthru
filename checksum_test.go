package sailthru_test

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/saml/sailthru"
)

func TestExtractLeaf(t *testing.T) {
	var m map[string]interface{}
	json.Unmarshal([]byte(`{
		"a": null,
		"b": [null, 1, true, false, "a", {"a": "b", "c": ["b", "c"]}, null, 0]
	}`), &m)

	expected := []string{"1", "1", "0", "a", "b", "b", "c", "0"}
	sort.Strings(expected)

	output := sailthru.ExtractParams(m)
	sort.Strings(output)

	if len(expected) != len(output) {
		t.Errorf("Length differs: %v != %v", expected, output)
	}

	for i, x := range output {
		if expected[i] != x {
			t.Errorf("%vth element is different: %v != %v", i, expected, output)
		}
	}
}
