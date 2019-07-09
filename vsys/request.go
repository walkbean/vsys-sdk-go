package vsys

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeOut = 15 * time.Second

func Post(url string, data interface{}) (body []byte, err error) {
	return PostWithTimeOut(url, data, defaultTimeOut)
}

func PostWithTimeOut(url string, data interface{}, duration time.Duration) (body []byte, err error) {
	client := http.Client{
		Timeout: duration,
	}
	d, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	if err := getErrResp(resp,body);err != nil{
		return []byte{}, err
	}

	return body, nil
}

func UrlGetContent(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	if err := getErrResp(resp,body);err != nil{
		return []byte{}, err
	}
	return body, nil
}

func getErrResp(resp *http.Response, body []byte)(err error){
	if resp.StatusCode != 200 {
		errResp := CommonResp{}
		if err := json.Unmarshal(body, &errResp); err != nil {
			return errors.New("StatusCodeError")
		} else {
			return errors.New(errResp.Message)
		}
	}
	return nil
}