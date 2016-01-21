package gtester

import (
	"encoding/json"
	"github.com/Focinfi/gtester/httpmock"
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, result, exp interface{}) {
	if !isEqual(result, exp) {
		t.Errorf("Expected %#v, Got %#v", exp, result)
	}
}

func isEqual(result, exp interface{}) bool {
	var equal bool
	switch exp.(type) {
	case string:
		if res, ok := result.(string); ok {
			equal = strings.EqualFold(res, exp.(string))
		}
	default:
		equal = result == exp
	}
	return equal
}

func AssertJsonEqual(t *testing.T, reponse *httpmock.Recoder, exp interface{}) {
	expectedStr, ok := exp.(string)
	if !ok {
		b, err := json.Marshal(exp)
		if err != nil {
			t.Error(err.Error())
			return
		} else {
			expectedStr = string(b)
		}
	}

	AssertEqual(t, strings.TrimRight(reponse.Body.String(), "\n"), expectedStr)
}
