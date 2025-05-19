package handler_test

/*
import (
	"bytes"
	"encoding/json"
	"meeting-scheduler/internal/handler"
	"meeting-scheduler/internal/model"
	"meeting-scheduler/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mock service
type mockService struct {
	service.SchedulerService
	users map[string]*model.User
}

func (m *mockService) GetUser(id string) (*model.User, error) {
	if u, ok := m.users[id]; ok {
		return u, nil
	}
	return nil, model.ErrUserNotFound
}

func (m *mockService) CreateUser(u *model.User) error {
	m.users[u.ID] = u
	return nil
}

func (m *mockService) GetAllUsers() (map[string]*model.User, error) {
	return m.users, nil
}

func setupRouter() (*gin.Engine, *mockService) {
	gin.SetMode(gin.TestMode)
	mockSvc := &mockService{users: make(map[string]*model.User)}
	h := handler.NewHandler((*service.SchedulerService)(mockSvc))
	r := gin.Default()
	h.RegisterRoutes(r)
	return r, mockSvc
}

func TestCreateUser(t *testing.T) {
	r, _ := setupRouter()

	user := model.User{ID: "u1", Name: "Alice"}
	jsonBody, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUser(t *testing.T) {
	r, svc := setupRouter()
	svc.users["u2"] = &model.User{ID: "u2", Name: "Bob"}

	req, _ := http.NewRequest("GET", "/user/u2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var user model.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "Bob", user.Name)
}

func TestGetAllUsers(t *testing.T) {
	r, svc := setupRouter()
	svc.users["u1"] = &model.User{ID: "u1", Name: "Alice"}
	svc.users["u2"] = &model.User{ID: "u2", Name: "Bob"}

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var users map[string]model.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}
*/
