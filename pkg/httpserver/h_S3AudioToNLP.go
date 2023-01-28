package httpserver

import (
	"net/http"
)

type s3fileURI struct {
	S3file string `json:"s3_file"`
}

// swagger:route POST /api/s3audiotonlp handleS3AudioToNLP
// fetches an S3 file and does STT and NLP steps
//
// Schemes: http
//
// Responses:
// default: errorResp
// 200: okResp
func (srv server) handleS3AudioToNLP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
		//// unmarshal body into
		//var fileURI s3fileURI
		//defer r.Body.Close()
		//
		//err := srv.helper.unMarshalReqBodyArgsJSONToStruct(r, &fileURI)
		//if err != nil {
		//	logger.Group("httpServer").
		//		WithFields(
		//			logrus.Fields{
		//				"received_body": r.Body,
		//			},
		//		).
		//		Errorf("unable to unMarshalReqBodyJSONToStruct, error: %v", err.Error())
		//	return
		//}
		//
		//// fetch S3 file
		//S3file, err := srv.s3client.DownloadFileFromS3Path(fileURI.S3file)
		//if err != nil {
		//	logger.Group("httpServer").
		//		WithFields(
		//			logrus.Fields{
		//				"requested file": fileURI.S3file,
		//			},
		//		).
		//		Errorf("unable to OpenFile, error: %v", err.Error())
		//	errorHandler(w, r, http.StatusInternalServerError)
		//	return
		//}
		//
		//// prepare STT request
		//extraParams := make(map[string]string)
		//extraParams["s3path"] = S3file.S3Path
		//request, err := srv.helper.newfileUploadRequest(srv.STTURI, extraParams, "audio_file", S3file.Path)
		//if err != nil {
		//	log.Fatal(err)
		//	errorHandler(w, r, http.StatusInternalServerError)
		//	return
		//}
		//
		//client := &http.Client{}
		//request.Header.Set("accept", "application/json")
		//resp, err := client.Do(request)
		//if err != nil {
		//	log.Fatal(err)
		//	errorHandler(w, r, http.StatusInternalServerError)
		//	return
		//}
		//
		//body := &bytes.Buffer{}
		//_, err = body.ReadFrom(resp.Body)
		//if err != nil {
		//	log.Fatal(err)
		//	errorHandler(w, r, http.StatusInternalServerError)
		//	return
		//}
		//
		//resp.Body.Close()
		//
		//// response
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
		//w.Write(body.Bytes())
		//
		//srv.newRelicApp.RecordCustomEvent("mediaMogulNlp: audio processed", map[string]interface{}{
		//	"bytesLenght": 1000,
		//})

		return
	}
}
