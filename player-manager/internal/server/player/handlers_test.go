package player

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"player-manager/internal/mocks"
	"player-manager/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func setupRouter(method, path string, handler gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Handle(method, path, handler)
	return r
}

func TestController_GetPlayer(t *testing.T) {
	playerRepositoryMock := mocks.NewPlayerRepository(t)
	tests := []struct {
		name        string
		username    string
		prepareMock func()
		wantCode    int
		wantBody    string
	}{
		{
			name:     "not existing player error",
			username: "not-existing-username",
			prepareMock: func() {
				playerRepositoryMock.On("GetPlayer", mock.Anything, "not-existing-username").
					Return(model.Player{}, fmt.Errorf("error"))
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":400,\"message\":\"failed to get player\",\"error\":\"error\"}",
		},
		{
			name:     "existing player",
			username: "existing-username",
			prepareMock: func() {
				playerRepositoryMock.On("GetPlayer", mock.Anything, "existing-username").
					Return(model.Player{Username: "existing-username"}, nil)
			},
			wantCode: http.StatusOK,
			wantBody: "{\"username\":\"existing-username\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			ctrl := &Controller{playerRepositoryMock}
			router := setupRouter(http.MethodGet, "/players/:username", ctrl.GetPlayer)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/players/"+tt.username, nil)
			require.Nil(t, err)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestController_ListItems(t *testing.T) {
	playerRepositoryMock := mocks.NewPlayerRepository(t)
	tests := []struct {
		name        string
		username    string
		prepareMock func()
		wantCode    int
		wantBody    string
	}{
		{
			name:     "not existing player error",
			username: "not-existing-username",
			prepareMock: func() {
				playerRepositoryMock.On("ListItems", mock.Anything, "not-existing-username").
					Return(model.Items{}, fmt.Errorf("error"))
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":400,\"message\":\"failed to get items\",\"error\":\"error\"}",
		},
		{
			name:     "existing player items",
			username: "existing-username",
			prepareMock: func() {
				playerRepositoryMock.On("ListItems", mock.Anything, "existing-username").
					Return(model.Items{{Name: "wood", Quantity: 10}}, nil)
			},
			wantCode: http.StatusOK,
			wantBody: "{\"items\":[{\"name\":\"wood\",\"quantity\":10}]}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			ctrl := &Controller{playerRepositoryMock}
			router := setupRouter(http.MethodGet, "/players/:username/items", ctrl.ListItems)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/players/"+tt.username+"/items", nil)
			require.Nil(t, err)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestController_AddItem(t *testing.T) {
	playerRepositoryMock := mocks.NewPlayerRepository(t)
	tests := []struct {
		name        string
		username    string
		quantity    string
		prepareMock func()
		wantCode    int
		wantBody    string
	}{
		{
			name:        "quantity param error",
			username:    "username",
			quantity:    "invalid",
			prepareMock: func() {},
			wantCode:    http.StatusBadRequest,
			wantBody:    "{\"status\":400,\"message\":\"invalid quantity\",\"error\":\"strconv.Atoi: parsing \\\"invalid\\\": invalid syntax\"}",
		},
		{
			name:     "add item error",
			username: "error",
			quantity: "1",
			prepareMock: func() {
				playerRepositoryMock.On("AddItem", mock.Anything, "error", mock.Anything, mock.Anything).
					Return(fmt.Errorf("error"))
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":400,\"message\":\"failed to add item\",\"error\":\"error\"}",
		},
		{
			name:     "add item success",
			username: "success",
			quantity: "1",
			prepareMock: func() {
				playerRepositoryMock.On("AddItem", mock.Anything, "success", mock.Anything, mock.Anything).
					Return(nil)
			},
			wantCode: http.StatusNoContent,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			ctrl := &Controller{playerRepositoryMock}
			router := setupRouter(http.MethodPut, "/players/:username/items/:item", ctrl.AddItem)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPut, "/players/"+tt.username+"/items/wood?quantity="+tt.quantity, nil)
			require.Nil(t, err)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestController_DeleteItem(t *testing.T) {
	playerRepositoryMock := mocks.NewPlayerRepository(t)
	tests := []struct {
		name        string
		username    string
		quantity    string
		prepareMock func()
		wantCode    int
		wantBody    string
	}{
		{
			name:        "quantity param error",
			username:    "username",
			quantity:    "invalid",
			prepareMock: func() {},
			wantCode:    http.StatusBadRequest,
			wantBody:    "{\"status\":400,\"message\":\"invalid quantity\",\"error\":\"strconv.Atoi: parsing \\\"invalid\\\": invalid syntax\"}",
		},
		{
			name:     "add item error",
			username: "error",
			quantity: "1",
			prepareMock: func() {
				playerRepositoryMock.On("DeleteItem", mock.Anything, "error", mock.Anything, mock.Anything).
					Return(fmt.Errorf("error"))
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":400,\"message\":\"failed to delete item\",\"error\":\"error\"}",
		},
		{
			name:     "add item success",
			username: "success",
			quantity: "1",
			prepareMock: func() {
				playerRepositoryMock.On("DeleteItem", mock.Anything, "success", mock.Anything, mock.Anything).
					Return(nil)
			},
			wantCode: http.StatusNoContent,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			ctrl := &Controller{playerRepositoryMock}
			router := setupRouter(http.MethodDelete, "/players/:username/items/:item", ctrl.DeleteItem)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodDelete, "/players/"+tt.username+"/items/wood?quantity="+tt.quantity, nil)
			require.Nil(t, err)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}

func TestController_Build(t *testing.T) {
	playerRepositoryMock := mocks.NewPlayerRepository(t)
	tests := []struct {
		name        string
		username    string
		prepareMock func()
		wantCode    int
		wantBody    string
	}{
		{
			name:     "error",
			username: "not-existing-username",
			prepareMock: func() {
				playerRepositoryMock.On("Build", mock.Anything, "not-existing-username").
					Return(fmt.Errorf("error"))
			},
			wantCode: http.StatusBadRequest,
			wantBody: "{\"status\":400,\"message\":\"failed to build\",\"error\":\"error\"}",
		},
		{
			name:     "success",
			username: "existing-username",
			prepareMock: func() {
				playerRepositoryMock.On("Build", mock.Anything, "existing-username").
					Return(nil)
			},
			wantCode: http.StatusNoContent,
			wantBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMock()
			ctrl := &Controller{playerRepositoryMock}
			router := setupRouter(http.MethodPost, "/players/:username/action/build", ctrl.Build)

			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/players/"+tt.username+"/action/build", nil)
			require.Nil(t, err)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Result().StatusCode)
			assert.Equal(t, tt.wantBody, w.Body.String())
		})
	}
}