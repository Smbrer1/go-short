package save_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/Smbrer1/go-short/internal/http-server/handlers/url/save/mocks"
)

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name      string
		url       string
		respError string
		mockError error
	}{
		{
			name: "Success",
			url:  "https://google.com",
		},
		{
			name:      "Empty URL",
			url:       "",
			respError: "field URL is a required field",
		},
		{
			name:      "Invalid URL",
			url:       "some invalid URL",
			respError: "field URL is not a valid URL",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			urlSaveMock := mocks.NewURLSaver(t)

			if tc.respError == "" || tc.mockError != nil {
				urlSaveMock.On("SaveURL", tc.url, mock.AnythingOfType("string")).Return(int64(1), tc.mockError).Once()
			}
		})
	}
}
