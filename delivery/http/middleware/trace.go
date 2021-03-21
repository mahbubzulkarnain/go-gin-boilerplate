package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// TraceJSON godoc.
type TraceJSON struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp string `json:"timestamp"`
	// StatusCode is HTTP response code.
	StatusCode int `json:"statuscode"`
	// Path is a path the client requests.
	Path string `json:"path"`
	// Method is the HTTP method given to the request.
	Method string `json:"method"`
	// Latency is how much time the server cost to process a certain request.
	Latency float64 `json:"latency"`
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"clientip"`
	// TraceJSON godoc.
	ClientAgent string `json:"clientagent"`
	// ErrorMessage is set if error has occurred in processing the request.
	// ErrorMessage string `json:"errormessage"`
	// BodySize is the size of the Response Body
	BodySize int `json:"bodysize"`
	// Request body
	Request map[string]interface{} `json:"request"`
	// Response body
	Response map[string]interface{} `json:"response"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Trace godoc.
func Trace(c *gin.Context) {
	// Start Timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	// Read the Body content
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	traceJSON := TraceJSON{
		ClientIP:    c.ClientIP(),
		ClientAgent: c.Request.UserAgent(),
		Method:      c.Request.Method,
		Request:     parseBody(bodyBytes),
	}

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	c.Next()

	traceJSON.StatusCode = c.Writer.Status()
	traceJSON.BodySize = c.Writer.Size()
	traceJSON.Path = path

	traceJSON.Response = parseBody(blw.body.Bytes())
	traceJSON.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	traceJSON.Latency = GetDurationInMillseconds(start)

	traceByte, _ := json.MarshalIndent(traceJSON, "", " ")
	traceString := string(traceByte)

	log.SetFlags(0)
	log.Print(traceString)
}

func parseBody(body []byte) (res map[string]interface{}) {
	_ = json.Unmarshal(body, &res)
	return
}

// GetDurationInMillseconds takes a start time and returns a duration in milliseconds
func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
