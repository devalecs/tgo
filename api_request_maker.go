package tgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type apiRequestMaker struct {
	token string
}

func (arm *apiRequestMaker) makeRequest(method string, params interface{}, v interface{}) error {
	url := arm.buildURL(method)

	bytesParams, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesParams))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	cli := new(http.Client)

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response := new(Response)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if !response.Ok {
		return errors.New(response.Description)
	}

	err = response.encodeResult(v)
	if err != nil {
		return err
	}

	return nil
}

func (arm *apiRequestMaker) makeMultipartRequest(method string, params interface{}, val interface{}) error {

	var b bytes.Buffer

	w := multipart.NewWriter(&b)

	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("form")
		value := v.Field(i)

		tagParts := strings.Split(tag, ",")
		field := tagParts[0]

		if field == "-" || arm.stringInSlice("omitempty", tagParts) && arm.isEmptyValue(value) {
			continue
		}

		if arm.stringInSlice("inputFile", tagParts) {
			path := fmt.Sprintf("%v", value.Interface())

			fw, err := w.CreateFormFile(field, path)
			if err != nil {
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			_, err = io.Copy(fw, file)
			if err != nil {
				return err
			}

			continue
		}

		encodedValue, err := json.Marshal(value.Interface())
		if err != nil {
			return err
		}

		fw, err := w.CreateFormField(field)
		if err != nil {
			return err
		}

		_, err = fw.Write([]byte(encodedValue))
		if err != nil {
			return err
		}
	}

	w.Close()

	url := arm.buildURL(method)

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	cli := new(http.Client)

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response := new(Response)

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if !response.Ok {
		return errors.New(response.Description)
	}

	err = response.encodeResult(val)
	if err != nil {
		return err
	}

	return nil
}

func (arm *apiRequestMaker) buildURL(method string) string {
	return fmt.Sprintf("%s/bot%s/%s", TelegramAPIEndpoint, arm.token, method)
}

func (arm *apiRequestMaker) stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}

	return false
}

func (arm *apiRequestMaker) isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
