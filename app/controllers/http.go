package controllers

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"tms/app/models"
)

const timeout = time.Duration(10 * time.Second)

func HttpGetRequest(url string, monitor models.Monitor, target interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: timeout}

	var (
		req *http.Request
		res *http.Response
		err error
	)

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return err
	}
	req.SetBasicAuth(monitor.MonitorName, monitor.MonitorPass)
	if res, err = client.Do(req); err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func DellHttpGetRequest(url string, monitor models.Monitor, target interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: timeout}

	var (
		req *http.Request
		res *http.Response
		err error
	)

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return err
	}
	req.Header["Accept"] = []string{"*/*"}
	req.SetBasicAuth(monitor.MonitorName, monitor.MonitorPass)
	if res, err = client.Do(req); err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func HttpPostRequest(url string, monitor models.Monitor, body io.Reader) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: timeout}

	var (
		req *http.Request
		res *http.Response
		err error
	)

	if req, err = http.NewRequest("POST", url, body); err != nil {
		return err
	}
	if monitor.Manuf == models.DELL {
		req.Header["Accept"] = []string{"*/*"}
	} else {
		req.Header["Content-Type"] = []string{"application/json"}
	}

	req.SetBasicAuth(monitor.MonitorName, monitor.MonitorPass)
	if res, err = client.Do(req); err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
