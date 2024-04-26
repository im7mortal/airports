// Created by Petr Lozhkin

package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/im7mortal/airports/pkg/airports/calc/test_cases"
	"github.com/im7mortal/airports/pkg/airports/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCalcHTTP_Get(t *testing.T) {

	sdkEngine := server.New()

	// run test server on 55555 port for more comfortable testing
	l, err := net.Listen("tcp", "127.0.0.1:55555")
	if err != nil {
		log.Fatal(err)
	}
	ts := httptest.NewUnstartedServer(sdkEngine.GetMainEngine())
	// NewUnstartedServer creates a listener. Close that listener and replace
	// with the one we created.
	ts.Listener.Close()
	ts.Listener = l
	// Start the server.
	ts.Start()
	// Stop the server on return from the function.
	defer ts.Close()

	tstUrl, err := url.Parse(ts.URL)
	if err != nil {
		glog.Error(err)
		return
	}
	tstUrl.Path = "/calculate"
	for i, testCase := range append(test_cases.Standard, test_cases.MaxSequenceRequest()) {
		requestBody, _ := json.Marshal(testCase.Input)
		resp, err := http.Post(tstUrl.String(), "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		println(resp.StatusCode)
		// You can check if the response status is as expected
		if resp.StatusCode != testCase.ExpectedStatusCode {

			b, _ := io.ReadAll(resp.Body)
			println(string(b))

			t.Errorf("Test case %d: got status = %d, want %d", i, resp.StatusCode, testCase.ExpectedStatusCode)
		}

		resp.Body.Close()
	}
}
