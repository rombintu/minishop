package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/rombintu/minishop/config"
	"github.com/rombintu/minishop/internal/app"
	"github.com/rombintu/minishop/internal/store"
)

func TestPing(t *testing.T) {
	config := config.GetConfig("../../config/config.toml")

	s := app.NewApp(config)

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	s.ConfigureRouter()
	s.Router.ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "{\"message\":\"pong\"}")
}

func TestCreateUser(t *testing.T) {
	config := config.GetConfig("../../config/config.toml")

	s := app.NewApp(config)
	body, err := json.Marshal(store.User{
		Account:  "user1",
		Password: "123",
		Role:     "user",
	})

	if err != nil {
		t.Fatal(err)
	}
	r := strings.NewReader(string(body))
	req, err := http.NewRequest("POST", "/user", r)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	s.ConfigureRouter()
	s.Router.ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "{\"message\":\"user created\"}")
}

func TestGetUser(t *testing.T) {
	config := config.GetConfig("../../config/config.toml")

	s := app.NewApp(config)

	body, err := json.Marshal(store.User{
		Account:  "user1",
		Password: "123",
		Role:     "user",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	s.ConfigureRouter()
	s.Router.ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.Bytes(), body)
}
