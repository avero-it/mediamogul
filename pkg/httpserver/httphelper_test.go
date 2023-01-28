package httpserver

import (
	"bytes"
	"github.com/Flaque/filet"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_newfileUploadRequest(t *testing.T) {
	tmpdir := filet.TmpDir(t, "") // Creates a temporary dir with no parent directory
	file := filet.TmpFile(t, tmpdir, "some content")

	type args struct {
		uri       string
		params    map[string]string
		paramName string
		path      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"all ok",
			args{
				uri: "",
				params: map[string]string{
					"ok": "ok",
				},
				paramName: "file",
				path:      file.Name(),
			},
			false,
		},
		{
			"fails to open file",
			args{
				uri: "",
				params: map[string]string{
					"ok": "ok",
				},
				paramName: "file",
				path:      "",
			},
			true,
		},
		{
			"missing param name",
			args{
				uri: "",
				params: map[string]string{
					"ok": "ok",
				},
				paramName: "",
				path:      "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHttpHelper()
			_, err := h.newfileUploadRequest(tt.args.uri, tt.args.params, tt.args.paramName, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("newfileUploadRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_unMarshalReqBodyArgsJSONToStruct(t *testing.T) {
	type args struct {
		r     *http.Request
		rArgs *s3fileURI
	}

	type testCase[T any] struct {
		name    string
		args    args
		wantErr bool
	}
	tests := []testCase[s3fileURI]{
		{
			"all ok",
			args{
				r: &http.Request{
					Body: ioutil.NopCloser(bytes.NewBufferString(
						`
{
  "s3_file": "audiodigitale_128.mp3"
}`,
					))},
				rArgs: &s3fileURI{
					S3file: "filename",
				},
			},
			false,
		},
		{
			"wrong body",
			args{
				r: &http.Request{
					Body: ioutil.NopCloser(bytes.NewBufferString("no JSON here"))},
				rArgs: &s3fileURI{
					S3file: "filename",
				},
			},
			true,
		},
		{
			"unreadable body",
			args{
				r: &http.Request{
					Body: ioutil.NopCloser(bytes.NewBufferString(""))},
				rArgs: &s3fileURI{
					S3file: "filename",
				},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := httpHelper{}
			err := h.unMarshalReqBodyArgsJSONToStruct(tt.args.r, tt.args.rArgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("unMarshalReqBodyArgsJSONToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
