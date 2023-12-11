package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const (
	titleTypeMovie    = "movie"
	titleTypeTVShow   = "tvSeries"
	topMoviesAPIPath  = "/top/movies"
	topTVShowsAPIPath = "/top/tv-shows"
	keyLimit          = "limit"
	keyNumVotes       = "numVotes"
	topTitlesKey      = "topTitles"
)

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/", s.handleHealthCheck)
	s.Router.HandleFunc(topMoviesAPIPath, s.handleTopTitles)
	s.Router.HandleFunc(topTVShowsAPIPath, s.handleTopTitles)
}

var (
	errInternalServerError      = errors.New("internal server error")
	internalServerErrorResponse = jsonResponse{Message: errInternalServerError.Error()}
	errInvalidTitleType         = errors.New("invalid title type")
	invalidTitleTypeResponse    = jsonResponse{Message: errInvalidTitleType.Error()}
)

type jsonResponse struct {
	Message string `json:"message"`
}

// respondJson sets the status code and response to be sent from the server for the http request and logs each request and status code
func respondJson(w http.ResponseWriter, r *http.Request, response interface{}, statusCode int, logger *slog.Logger) {
	logger.Info(fmt.Sprintf("Request URL: %+v, status code=%+v\n", r.URL, statusCode))
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Info("error while marshaling json ", err)
		return
	}
	logger.Debug(fmt.Sprintf("Response body : %+v", string(responseBytes)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(responseBytes)
	if err != nil {
		logger.Info("error writing into responseWriter %+v\n", err)
		return
	}
}

// handleHealthCheck returns an ok response to signal that the server is alive
func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJson(w, r, map[string]string{"status": "ok"}, http.StatusOK, s.Logger)
}

func (s *Server) handleTopTitles(w http.ResponseWriter, r *http.Request) {
	logger := s.Logger.With("request_id", uuid.NewString())
	var titleType string
	if strings.Contains(r.URL.Path, topMoviesAPIPath) {
		titleType = titleTypeMovie
	} else if strings.Contains(r.URL.Path, topTVShowsAPIPath) {
		titleType = titleTypeTVShow
	} else {
		respondJson(w, r, invalidTitleTypeResponse, http.StatusBadRequest, logger)
		return
	}

	limit, numVotes, err := s.parseTopTitlesParams(r)
	if err != nil {
		slog.Debug("error while parsing top titles params %+v\n", err)
		respondJson(w, r, internalServerErrorResponse, http.StatusInternalServerError, logger)
		return
	}
	topTitles, err := s.getTopTitles(limit, numVotes, titleType)
	if err != nil {
		slog.Debug("error while getting top titles %+v\n", err)
		respondJson(w, r, internalServerErrorResponse, http.StatusInternalServerError, logger)
		return
	}
	respondJson(w, r, map[string]interface{}{topTitlesKey: topTitles}, http.StatusOK, logger)
}

func (s *Server) parseTopTitlesParams(r *http.Request) (int, int, error) {
	limit := 10
	numVotes := 1000
	var err error
	if r.URL.Query().Get(keyLimit) != "" {
		limit, err = strconv.Atoi(r.URL.Query().Get(keyLimit))
		if err != nil {
			return 0, 0, err
		}
		if limit > 250 {
			limit = 250
		}
	}
	if r.URL.Query().Get(keyNumVotes) != "" {
		numVotes, err = strconv.Atoi(r.URL.Query().Get(keyNumVotes))
		if err != nil {
			return 0, 0, err
		}
	}
	return limit, numVotes, nil
}

func (s *Server) getTopTitles(limit, numVotes int, titleType string) ([]map[string]interface{}, error) {
	rows, err := s.DB.Queryx(fmt.Sprintf(topTitlesQuery, titleType, numVotes, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies = make([]map[string]interface{}, 0)
	for rows.Next() {
		var movie Title
		err = rows.StructScan(&movie)
		if err != nil {
			return nil, err
		}
		mapVal := StructToMap(movie)
		movies = append(movies, mapVal)
	}
	return movies, nil
}

func StructToMap(input interface{}) map[string]interface{} {
	val := reflect.ValueOf(input)
	out := make(map[string]interface{})

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		switch field.Kind() {
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(sql.NullString{}) {
				if field.Interface().(sql.NullString).Valid {
					out[fieldType.Name] = field.Interface().(sql.NullString).String
				}
			} else if field.Type() == reflect.TypeOf(sql.NullInt32{}) {
				if field.Interface().(sql.NullInt32).Valid {
					out[fieldType.Name] = field.Interface().(sql.NullInt32).Int32
				}
			} else if field.Type() == reflect.TypeOf(sql.NullBool{}) {
				if field.Interface().(sql.NullBool).Valid {
					out[fieldType.Name] = field.Interface().(sql.NullBool).Bool
				}
			} else if field.Type() == reflect.TypeOf(sql.NullFloat64{}) {
				if field.Interface().(sql.NullFloat64).Valid {
					out[fieldType.Name] = field.Interface().(sql.NullFloat64).Float64
				}
			} else {
				mapVal := StructToMap(field.Interface())
				if len(mapVal) > 0 {
					out[fieldType.Name] = StructToMap(field.Interface())
				}
			}
		}
	}
	return out
}
