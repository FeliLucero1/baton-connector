package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/stretchr/testify/assert"
	""your_project/client""
)

func TestListUsers(t *testing.T) {
	expectedUsers := []client.User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedUsers)
	}))
	defer server.Close()

	clientInstance := client.NewClient("testuser", "testpass", &uhttp.BaseHttpClient{Client: server.Client()})
	clientInstance.baseURL = server.URL

	users, err := clientInstance.ListUsers(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestGetUserByID(t *testing.T) {
	expectedUser := &client.User{ID: 1, Username: "user1", Email: "user1@example.com"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/1", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedUser)
	}))
	defer server.Close()

	clientInstance := client.NewClient("testuser", "testpass", &uhttp.BaseHttpClient{Client: server.Client()})
	clientInstance.baseURL = server.URL

	user, err := clientInstance.GetUserByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
