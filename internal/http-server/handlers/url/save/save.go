package save

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"github.com/Smbrer1/go-short/internal/helpers/api/response"
	"github.com/Smbrer1/go-short/internal/helpers/logger/sl"
	"github.com/Smbrer1/go-short/internal/helpers/urlcoder"
)

type Request struct {
	URL string `json:"url"             validate:"requiered,url"`
}

type Response struct {
	response.Response
	Error string `json:"error,omitempty"`
	Alias string `json:"alias,omitempty"`
}

// TODO: move aliasLength to config
const aliasLength = 6

type URLSaver interface {
	SaveURL(urlToSave string) (int64, error)
}

func New(log *slog.Logger, URLSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, response.Error("Failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err = validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, response.ValidationError(validateErr))

			return
		}

		id, err := URLSaver.SaveURL(req.URL)
		if err != nil {
			log.Error("failed to add url", sl.Err(err))

			render.JSON(w, r, response.Error("failed to add url"))
		}

		log.Info("url added", slog.Int64("id", id))

		alias, err := urlcoder.Encode(id)
		if err != nil {
			log.Error("failed to encode id", sl.Err(err))
			render.JSON(w, r, response.Error("failed to encode url"))
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}
