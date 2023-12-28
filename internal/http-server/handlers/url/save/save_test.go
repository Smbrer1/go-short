package save_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/Smbrer1/go-short/internal/helpers/logger/handlers/slogdiscard"
	"github.com/Smbrer1/go-short/internal/http-server/handlers/url/save"
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
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			urlSaveMock := mocks.NewURLSaver(t)

			if tc.respError == "" || tc.mockError != nil {
				urlSaveMock.On("SaveURL", tc.url, mock.AnythingOfType("string")).Return(int64(1), tc.mockError).Once()
			}

			handler := save.New(slogdiscard.NewDiscardLogger(), urlSaveMock)
			input := fmt.Sprintf(`{"url": "%s"}`, tc.url)

			req, err := http.NewRequest(http.MethodPost, "/save", bytes.NewReader([]byte(input)))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, rr.Code, http.StatusOK)

			body := rr.Body.String()

			var resp save.Response

			require.NoError(t, json.Unmarshal([]byte(body), &resp))

			require.Equal(t, tc.respError, resp.Error)
		})
	}
}
