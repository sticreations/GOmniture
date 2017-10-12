package GOmniture

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Create new Gomniture by calling New()
type Gomniture struct {
	username, sharedSecred string
}

func New(username, sharedSecred string) *Gomniture {
	gomni := Gomniture{username, sharedSecred}
	return &gomni
}

func buildMd5(s string) string {
	md5 := md5.New()
	io.WriteString(md5, s)
	return fmt.Sprintf("%x", md5.Sum(nil))
}

func sha_64(s string) string {
	h := sha1.New()
	enc := base64.StdEncoding
	io.WriteString(h, s)
	bytes := h.Sum(nil)
	return enc.EncodeToString(bytes)
}

func (gomni *Gomniture) buildXWSSE() string {
	enc := base64.StdEncoding
	time := time.Now().In(time.UTC)
	createdDate := time.Format("2006-01-02T15:04:05Z")
	nonce := buildMd5(time.String())
	base64Nonce := enc.EncodeToString([]byte(nonce))
	passwordDigest := sha_64(fmt.Sprintf("%s%s%s", nonce, createdDate, gomni.sharedSecred))
	return fmt.Sprintf("UsernameToken Username=\"%s\", PasswordDigest=\"%s\", Nonce=\"%s\", Created=\"%s\"", gomni.username, passwordDigest, base64Nonce, createdDate)
}

func (gomni *Gomniture) sendRequest(method string, request interface{}) ([]byte, error) {
	client := &http.Client{}

	rq, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest("POST", "https://api.omniture.com/admin/1.4/rest/?method="+method, bytes.NewBuffer(rq))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Add("X-WSSE", gomni.buildXWSSE())
	res, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

/*
QueueReport dispatches a ReportQuery Request and returns a ReportID on success
*/
func (gomni *Gomniture) QueueReport(request ReportQuery) (int, error) {
	resp, err := gomni.sendRequest("Report.Queue", request)
	if err != nil {
		return -1, err
	}
	var response queueReportResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return -1, err
	}
	return response.ReportID, nil
}

/*
Gets Report by reportID
*/
func (gomni *Gomniture) GetReport(reportID int) (ReportResponse, error) {
	var response ReportResponse
	resp, err := gomni.sendRequest("Report.Get", getReport{ReportID: reportID})
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
