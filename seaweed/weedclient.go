package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type Client struct {
	Address string
}

// {"count":1,"fid":"3,01637037d6","url":"127.0.0.1:8080","publicUrl":"localhost:8080"}
type AssignResult struct {
	Count     int
	Fid       string
	Url       string
	PublicUrl string
}

func (c Client) Assign() (AssignResult, error) {
	var result AssignResult

	url := fmt.Sprintf("http://%s/dir/assign", c.Address)
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(bs, &result); err != nil {
		return result, err
	}
	return result, nil
}

// {"name":"myphoto.jpg","size":43234,"eTag":"1cc0118e"}
type StoreResult struct {
	AssignResult
	Name string
	Size int64
	ETag string
}

func (c Client) Store(filename string, body io.Reader) (result StoreResult, err error) {
	assignResult, err := c.Assign()
	if err != nil {
		return result, err
	}

	url := assignResult.Url
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	form, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return result, err
	}

	if _, err := io.Copy(form, body); err != nil {
		return result, err
	}

	storeUrl := url + "/" + assignResult.Fid
	contentType := writer.FormDataContentType()

	resp, err := http.Post(storeUrl, contentType, &buffer)
	if err != nil {
		return result, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(bs, &result); err != nil {
		return result, err
	}
	return result, nil
}
