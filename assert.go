package gtester

import (
	"encoding/json"
	"testing"
)

func AssertEqaul(t *testing.T, result, exp interface{}) {
	if result != exp {
		t.Errorf("Expected %v, Got %v", exp, result)
	}
}

func AssertJsonEqaul(t *testing.T, reponse *recoder, exp interface{}) {
	var expectedStr string
	if _, ok := exp.(string); !ok {
		b, err := json.Marshal(exp)
		if err != nil {
			t.Error(err.Error())
			return
		} else {
			expectedStr = string(b)
		}
	}

	AssertEqaul(t, reponse.Body.String(), expectedStr)
}
