package gtester

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test(t *testing.T) {
	respMap := map[string]string{"hello": "world"}
	respJson, _ := json.Marshal(respMap)
	handler := func(wr http.ResponseWriter, req *http.Request) {
		wr.WriteHeader(http.StatusOK)
		fmt.Fprint(wr, string(respJson))
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	router := gin.Default()
	ginHandler := func(ctx *gin.Context) {
		ctx.String(http.StatusOK, string(respJson))
	}
	router.GET("/hello", ginHandler)
	ListenAndServe("hello.com", router)

	response := NewRecorder()
	Get("/hello", response)
	AssertJsonEqaul(t, response, respMap)
	AssertEqaul(t, response.Code, http.StatusOK)
}
