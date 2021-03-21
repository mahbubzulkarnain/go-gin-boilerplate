package httpclient

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gomodul/fn"
	"github.com/mahbubzulkarnain/golang-gin-boilerplate/helper/response"
)

// Request godoc.
func Request(method, baseURL string, payload interface{}, result interface{}) *response.JSON {
	body, errMarshal := json.Marshal(payload)
	if errMarshal != nil {
		return response.Err(errMarshal)
	}

	req, errRequest := http.NewRequest(method, baseURL, bytes.NewBuffer(body))
	if errRequest != nil {
		return response.Err(errRequest)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "mozze_cart_order_engine/1.0")

	client := &http.Client{Timeout: time.Second * time.Duration(30)}
	resDo, errDo := client.Do(req)
	if errDo != nil {
		switch err := errDo.(type) {
		case net.Error:
			if err.Timeout() {
				return response.ErrGatewayTimeout.SetError(err)
			}
		}
		return response.Err(errDo)
	}
	defer fn.Check(resDo.Body.Close)

	if resDo.StatusCode < http.StatusOK || resDo.StatusCode >= http.StatusMultipleChoices {
		return response.ErrByStatusCode(resDo.StatusCode)
	}

	errParse := json.NewDecoder(resDo.Body).Decode(&result)
	if errParse != nil {
		return response.Err(errParse)
	}

	return nil
}
