package save

import (
	"log/slog"
	"net/http"

	"github.com/Smbrer1/go-short/internal/helpers/api/response"
	"github.com/gin-contrib/requestid"
)

type Request struct {
	URL   string `json:"url"             validate:"requiered,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Error string `json:"error,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, URLSaver URLSaver) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"
		log = log.With(slog.String("op", op),
			// TODO add request id middleware
		// slog.String("request_id", middle )
	}
}
