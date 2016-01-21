package gtester

import (
	"encoding/json"
	"github.com/Focinfi/gtester/httpmock"
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, result, exp interface{}) {
	if !Equal(result, exp) {
		t.Errorf("Expected %#v, Got %#v", exp, result)
	}
}

func Equal(result, exp interface{}) bool {
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

func AssertResponseEqual(t *testing.T, reponse *httpmock.Recorder, exp interface{}) {
	if !ResponseEqual(reponse, exp) {
		t.Errorf("Expected %#v, Got %#v", exp, reponse.Body.String())
	}
}

func stringfy(exp interface{}) string {
	expectedStr, ok := exp.(string)
	if !ok {
		b, err := json.Marshal(exp)
		if err != nil {
			return ""
		} else {
			expectedStr = string(b)
		}
	}
	return expectedStr
}

func ResponseEqual(response *httpmock.Recorder, exp interface{}) bool {
	return Equal(strings.TrimRight(response.Body.String(), "\n"), stringfy(exp))
}
