package server

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"events-api/pkg/config"
	"events-api/pkg/errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RestClient interface {
	Get(uriKey string, uriParams map[string]string, body interface{}, result interface{}, headers map[string]string) error
}

func NewRestClient(config config.RestConfig) RestClient {
	return &restClient{
		httpClient: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       config.HTTPClient.Timeout,
		},
		apiDomain:     config.ApiDomain,
		externalCalls: config.ExternalCalls,
	}
}

type restClient struct {
	httpClient    http.Client
	apiDomain     string
	externalCalls map[string]config.ExternalCall
}

func (r *restClient) Get(uriKey string, uriParams map[string]string, body interface{}, result interface{}, headers map[string]string) error {
	req, url, err := r.createRequest(uriKey, "GET", uriParams, body, headers)
	if err != nil {
		return err
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		if resp != nil {
			if resp.StatusCode == http.StatusBadRequest {
				return errors.NewValidationError("bad_request", fmt.Sprintf("failed to make GET request to %s: %s", url, err))
			}
		}
		return errors.NewCommunicationError(fmt.Sprintf("failed to make GET request to %s: %s", url, err))
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		println(context.Background(), fmt.Sprintf("fatal error reading response body from GET to %s: %+v", url, err))
		return errors.NewUnknownError("unexpected_status_code", fmt.Sprintf("fatal error reading response body from GET to %s", url))
	}

	bodyString := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		return errors.NewUnknownError("unexpected_status_code", fmt.Sprintf("unexpected status code in GET to %s: %d (%s)", url, resp.StatusCode, bodyString))
	}

	if resp.Header.Get("Content-Type") == "application/json" {
		err = json.Unmarshal(bodyBytes, result)
	}
	if resp.Header.Get("Content-Type") == "text/xml" {
		err = xml.Unmarshal(bodyBytes, result)
	}
	if err != nil {
		return errors.NewCommunicationError(fmt.Sprintf("rest: unmarshall of response failed: uriKey: %s - error: %s", uriKey, err))
	}

	return nil
}

func (r *restClient) createRequest(uriKey string, method string, uriParams map[string]string, body interface{}, headers map[string]string) (*http.Request, string, error) {
	b := new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, "", errors.NewValidationError("error_encoding_body", fmt.Sprintf("error encoding body in %s: %s", method, err))
		}
	}

	url, err := r.formatExternalCall(uriKey, uriParams)
	if err != nil {
		return nil, "", errors.NewValidationError("error_retrieving_url", fmt.Sprintf("error retrieving url in %s: %s", method, err))
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, "", errors.NewUnknownError("error_creating_request", fmt.Sprintf("error creating request in %s to %s", method, url))
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, url, nil
}

func (r *restClient) formatExternalCall(uriKey string, uriParams map[string]string) (string, error) {
	endpoint := r.getExternalCall(uriKey)
	if endpoint == "" {
		return "", fmt.Errorf(fmt.Sprintf("invalid uriKey %s", uriKey))
	}

	formattedUrl := r.apiDomain + r.formatEndpoint(endpoint, uriParams)

	return formattedUrl, nil
}

func (r *restClient) getExternalCall(uriKey string) string {
	if url, ok := r.externalCalls[uriKey]; ok {
		return url.RequestUri
	}

	return ""
}

func (r *restClient) formatEndpoint(url string, uriParams map[string]string) string {
	args, i := make([]string, len(uriParams)*2), 0
	for k, v := range uriParams {
		args[i] = "{" + k + "}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}

	return strings.NewReplacer(args...).Replace(url)
}
