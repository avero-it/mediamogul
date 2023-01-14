//go:generate mockgen -destination=../,,/mock/httpserver/httphelper.go . HttpHelper

package httpserver

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type HttpHelper[T any] interface {
	newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error)
	unMarshalReqBodyArgsJSONToStruct(r *http.Request, t *T) error
}

type httpHelper struct {
}

var _ httpHelper = httpHelper{}

func NewHttpHelper() httpHelper {
	return httpHelper{}
}

func (helper httpHelper) newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func (helper httpHelper) unMarshalReqBodyArgsJSONToStruct(r *http.Request, reqArgs *s3fileURI) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &reqArgs)
	if err != nil {
		return err
	}

	return nil
}
