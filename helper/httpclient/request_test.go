package httpclient_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mahbubzulkarnain/golang-gin-boilerplate/delivery/http/handler"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/helper/httpclient"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/helper/response"
)

func TestRequest(t *testing.T) {
	var res response.JSON
	payload := handler.PlaceHolderPayload{
		StatusCode: http.StatusGatewayTimeout,
	}

	err := httpclient.Request(http.MethodPost, "http://localhost:9000/", payload, &res)
	if err != nil {
		t.Fatal(err)
	}

	out, _ := json.MarshalIndent(res, "", "")
	fmt.Println(out)
}
