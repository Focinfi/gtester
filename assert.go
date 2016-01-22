package gtester

import (
	"encoding/json"
	"github.com/Focinfi/gtester/httpmock"
	"reflect"
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, result, exp interface{}) {
	if !Equal(result, exp) {
		t.Errorf("Expected %#v, Got %#v", exp, result)
	}
}

func Equal(result, exp interface{}) bool {
	return reflect.DeepEqual(result, exp)
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
