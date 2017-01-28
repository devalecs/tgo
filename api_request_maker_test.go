package tgo

import "testing"

func TestBuildURL(t *testing.T) {
	arm := apiRequestMaker{
		token: "mockToken",
	}

	u := arm.buildURL("GetMe")

	if u != "https://api.telegram.org/botmockToken/GetMe" {
		t.Fatal("built URL is different than expected")
	}
}

// TODO: test makeRequest
// TODO: test makeMultipartRequest
