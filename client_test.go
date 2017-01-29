package tgo

import (
	"errors"
	"reflect"
	"testing"
)

// TODO: test NewClient

// GetUpdates

func TestGetUpdates(t *testing.T) {
	requestMaker := new(requestMakerMock)

	c := NewClient("mockToken")
	c.requestMaker = requestMaker

	params := GetUpdatesParams{
		Timeout: 60,
	}

	updates, err := c.GetUpdates(params)
	if err != nil {
		t.Fatal(err)
	}

	if !requestMaker.makeRequestCalled {
		t.Fatal("makeRequest method not called as expected")
	}

	if requestMaker.passedMethod != "getUpdates" {
		t.Fatal("passed method to requestMaker is different than expected")
	}

	if !reflect.DeepEqual(requestMaker.passedParams.(GetUpdatesParams), params) {
		t.Fatal("passed params to requestMaker are different than expected")
	}

	if !reflect.DeepEqual(requestMaker.passedV.(*[]Update), updates) {
		t.Fatal("passed v to requestMaker is different than expected")
	}
}

func TestGetUpdatesForwardsRequestMakerErr(t *testing.T) {
	requestMaker := &requestMakerMock{
		errorToReturn: errors.New("mock error"),
	}

	c := NewClient("mockToken")
	c.requestMaker = requestMaker

	_, err := c.GetMe()
	if err == nil || err.Error() != "mock error" {
		t.Fatal("returned error is different than expected")
	}
}

// GetMe

func TestGetMe(t *testing.T) {
	requestMaker := new(requestMakerMock)

	c := NewClient("mockToken")
	c.requestMaker = requestMaker

	user, err := c.GetMe()
	if err != nil {
		t.Fatal(err)
	}

	expectedUser := User{}
	if *user != expectedUser {
		t.Fatal("returned user is different than expected")
	}

	if !requestMaker.makeRequestCalled {
		t.Fatal("makeRequest method not called as expected")
	}

	if requestMaker.passedMethod != "getMe" {
		t.Fatal("passed method to requestMaker is different than expected")
	}
}

func TestGetMeForwardsRequestMakerErr(t *testing.T) {
	requestMaker := &requestMakerMock{
		errorToReturn: errors.New("mock error"),
	}

	c := NewClient("mockToken")
	c.requestMaker = requestMaker

	params := GetUpdatesParams{
		Timeout: 60,
	}

	_, err := c.GetUpdates(params)
	if err == nil || err.Error() != "mock error" {
		t.Fatal("returned error is different than expected")
	}
}

// SendMessage

func TestSendMessage(t *testing.T) {
	requestMaker := new(requestMakerMock)

	c := NewClient("mockToken")
	c.requestMaker = requestMaker

	params := SendMessageParams{
		ChatID: "mockChatID",
		Text:   "mock text",
	}

	message, err := c.SendMessage(params)
	if err != nil {
		t.Fatal(err)
	}

	if !requestMaker.makeRequestCalled {
		t.Fatal("makeRequest method not called as expected")
	}

	if requestMaker.passedMethod != "sendMessage" {
		t.Fatal("passed method to requestMaker is different than expected")
	}

	if !reflect.DeepEqual(requestMaker.passedParams.(SendMessageParams), params) {
		t.Fatal("passed params to requestMaker are different than expected")
	}

	if !reflect.DeepEqual(requestMaker.passedV.(*Message), message) {
		t.Fatal("passed v to requestMaker is different than expected")
	}
}

func TestSendMessageForwardsRequestMakerErr(t *testing.T) {
	requestMaker := &requestMakerMock{
		errorToReturn: errors.New("mock error"),
	}

	c := NewClient("mockToken")
	c.requestMaker = requestMaker

	params := SendMessageParams{
		ChatID: "mockChatID",
		Text:   "mock text",
	}

	_, err := c.SendMessage(params)
	if err == nil || err.Error() != "mock error" {
		t.Fatal("returned error is different than expected")
	}
}

// TODO: test all client methods

// Mocks

type requestMakerMock struct {
	makeRequestCalled          bool
	makeMultipartRequestCalled bool

	passedMethod string
	passedParams interface{}
	passedV      interface{}

	errorToReturn error
}

func (rm *requestMakerMock) makeRequest(method string, params interface{}, v interface{}) error {
	rm.makeRequestCalled = true

	rm.passedMethod = method
	rm.passedParams = params
	rm.passedV = v

	return rm.errorToReturn
}

func (rm *requestMakerMock) makeMultipartRequest(method string, params interface{}, v interface{}) error {
	rm.makeMultipartRequestCalled = true

	rm.passedMethod = method
	rm.passedParams = params
	rm.passedV = v

	return rm.errorToReturn
}
