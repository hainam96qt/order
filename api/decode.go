package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const MaxContentLength = 10000000

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, MaxContentLength)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	d := struct {
		Data interface{} `json:"data"`
	}{Data: dst}
	if err := json.Unmarshal(body, &d); err != nil {
		return err
	}
	return nil
}
