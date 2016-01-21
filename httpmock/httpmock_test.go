package httpmock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Focinfi/gtester"
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
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	}
	ginHandler2 := func(ctx *gin.Context) {
		ctx.String(http.StatusOK, string("xx"))
	}
	router.GET("/hello/:id", ginHandler)
	router.GET("/hellos", ginHandler2)
	ListenAndServe("hello.com", router)

	response := NewRecorder()
	GET("/hello/1", response)
	// gtester.gtesterAssertJsonEqual(t, response, respMap)
	// gtester.gtesterAssertJsonEqual(t, response, `{"hello":"world"}`)
	// gtester.AssertEqual(t, response.Code, http.StatusOK)
}
