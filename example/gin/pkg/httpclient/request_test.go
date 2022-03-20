package httpclient_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"modul/example/gin/internal/delivery/http/handler"
	"modul/example/gin/pkg/httpclient"
	"modul/example/gin/pkg/response"

	netHTTP "net/http"
)

func TestRequest(t *testing.T) {
	t.Skip()

	var res response.JSON
	payload := http.PlaceHolderPayload{
		StatusCode: netHTTP.StatusGatewayTimeout,
	}

	err := httpclient.Request(netHTTP.MethodPost, "http://localhost:9000/", payload, &res)
	if err != nil {
		t.Fatal(err)
	}

	out, _ := json.MarshalIndent(res, "", "")
	fmt.Println(out)
}
