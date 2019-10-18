package handler

import (
	"reflect"
	"testing"
)

func TestLocalRequest(t *testing.T) {
	handler := mockMosaicHandler{}
	res, _ := handler.LocalRequest("golang")
	want := "path/to/bucket/key"

	if res == "" {
		t.Error("Response is empty")
	}

	if !reflect.DeepEqual(res, want) {
		t.Errorf("Response not equat, got %v want %v", res, want)
	}
}

type mockMosaicHandler struct {
	mockInsta interface{}
}

type mockMosaicResponse struct {
	key string
}

func (m *mockMosaicHandler) HandleRequest(fakeTag string) (*mockMosaicResponse, error) {
	return &mockMosaicResponse{
		key: "path/to/bucket/key",
	}, nil
}

func (m *mockMosaicHandler) LocalRequest(fakeTag string) (string, error) {
	return "path/to/bucket/key", nil
}
