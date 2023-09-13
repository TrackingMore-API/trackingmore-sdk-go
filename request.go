package trackingmore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

var (
	baseUrl    string
	apiVersion string
)

func init() {

	baseUrl = "https://api.trackingmore.com/"

	apiVersion = "v4"

}

func (client *Client) sendApiRequest(ctx context.Context, method, path string, queryParams interface{}, inputData interface{}, resultData interface{}) (*Response, error) {
	var body io.Reader
	if inputData != nil {
		jsonData, err := json.Marshal(inputData)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	}

	requestUrl := baseUrl + apiVersion + path
	req, err := http.NewRequestWithContext(ctx, method, requestUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Tracking-Api-Key", client.apiKey)

	if queryParams != nil {
		queryString := url.Values{}
		if err := addStructParams(queryParams, &queryString); err != nil {
			return nil, err
		}
		req.URL.RawQuery = queryString.Encode()
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody := new(bytes.Buffer)
	respBody.ReadFrom(resp.Body)

	fmt.Println("Body:", respBody.String())

	result := &Response{
		Meta: Meta{},
		Data: resultData,
	}

	err = json.Unmarshal([]byte(respBody.String()), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func addStructParams(params interface{}, values *url.Values) error {
	v := url.Values{}
	val := reflect.ValueOf(params)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("params must be a struct or a pointer to struct")
	}
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" {
			tag = field.Name
		}
		value := val.Field(i).Interface()
		if value != reflect.Zero(val.Field(i).Type()).Interface() {
			v.Add(tag, fmt.Sprintf("%v", value))
		}
	}
	*values = v
	return nil
}
