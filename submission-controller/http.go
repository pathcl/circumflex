package submission_controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func get(url string) ([]byte, error) {
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer closeStream(r.Body)

	body, readError := ioutil.ReadAll(r.Body)
	if readError != nil {
		return nil, readError
	}

	return body, nil
}

func closeStream(body io.ReadCloser) {
	_ = body.Close()
}

func getSubmissions(url string) []Submission{
	JSON, _ := get(url)
	var submissions []Submission
	_ = json.Unmarshal(JSON, &submissions)
	return submissions
}
