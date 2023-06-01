package controller_test

//import (
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/mock"
//	uDomain "github.com/superosystem/bantumanten-backend/src/businesses/users"
//	uMock "github.com/superosystem/bantumanten-backend/src/businesses/users/mocks"
//	"github.com/superosystem/bantumanten-backend/src/controllers/users"
//	"net/http"
//	"net/http/httptest"
//	"net/url"
//	"strings"
//	"testing"
//	"time"
//)
//
//var (
//	userDomain     uDomain.Domain
//	userUseCase    = new(uMock.UseCase)
//	userController users.Controller
//)
//
//func TestMain(m *testing.M) {
//
//	userDomain = uDomain.Domain{
//		ID:        1,
//		FullName:  "John Doe",
//		Email:     "johndoe@exampl.com",
//		Password:  "hashedpassword",
//		Photo:     "",
//		Roles:     "USER",
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//
//	m.Run()
//}
//
//func TestSignUp(t *testing.T) {
//	reqBody := make(url.Values)
//	reqBody.Set("full_name", "John Doe")
//	reqBody.Set("email", "johndoe@example.com")
//	reqBody.Set("password", "passwd123")
//	reqBody.Set("confirmation_password", "passwd123")
//
//	//_ := users.NewUserController(userUseCase)
//
//	t.Run("SIGN UP | SUCCESS", func(t *testing.T) {
//		e := echo.New()
//		req := httptest.NewRequest(http.MethodPost, "/request/v1/auth/register",
//			strings.NewReader(reqBody.Encode()))
//		rec := httptest.NewRecorder()
//		c := e.NewContext(req, rec)
//
//	//t.Run("SIGN UP | SUCCESS", func(t *testing.T) {
//	//
//	//	userUseCase.On("SignUp", mock.AnythingOfType("*users.Domain")).Return(nil).Once()
//	//
//	//	e := echo.New()
//	//	rec := httptest.NewRecorder()
//	//
//	//	req := httptest.NewRequest(http.MethodPost,
//	//		"/request/v1/auth/register", strings.NewReader(reqBody.Encode()))
//	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	//
//	//	e.ServeHTTP(rec, req)
//	//
//	//	assert.Equal(t, http.StatusCreated, rec.Code)
//	//	//controller := userController.SignUp(c)
//	//	//
//	//	//if assert.NoError(t, controller) {
//	//	//
//	//	//}
//	//})
//
//}
