//go:build e2e

package gql

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"interview/pkg/core/common"
	"interview/pkg/core/service"
	"interview/pkg/infrastructure/db"
	"interview/pkg/infrastructure/repository"
	"interview/pkg/interface/api"
	"interview/pkg/logger"
	"net/http"
	"net/http/httptest"
	"net/url"

	"testing"
)

func TestGqlTestSuite(t *testing.T) {
	suite.Run(t, new(GinTestSuite))
}

type GinTestSuite struct {
	suite.Suite
	api    api.API
	dbRepo *gorm.Model
}

func (s *GinTestSuite) SetupSuite() {
	logger.Init("SetupSuite")
	lg := logger.General.Component("GinTestSuite")
	lg.Println("test started")

	dbRepo := repository.NewRepository(db.GetDatabase(true, ""))
	cartService := service.NewCartService(service.CartServiceParam{
		CartItemRepo:   dbRepo,
		CartEntityRepo: dbRepo,
	})

	s.api = api.NewServer(&api.ServerConfig{
		Port:    8088,
		GinMode: "test",
		Services: common.Services{
			Cart: cartService,
		},
	})

}

func (s *GinTestSuite) TearDownSuite() {

}

func (s *GinTestSuite) TestBaseGql() {
	req, err := http.NewRequest("GET", "/", nil)

	s.Require().NoError(err)

	req.Header.Set("Content-Type", "text/html")

	w := httptest.NewRecorder()
	//s.api.SetSessionKye()
	s.api.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)
}

func (s *GinTestSuite) TestAddItem() {
	formData := url.Values{
		"product":  {"show"},
		"quantity": {"1"},
	}

	req, err := http.NewRequest("POST", "/add-item", bytes.NewBufferString(formData.Encode()))

	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	s.api.ServeHTTP(w, req)

	s.Require().Equal(http.StatusFound, w.Code)
}

func (s *GinTestSuite) TestFailedAddItem() {

	formData := url.Values{
		"product":  {"value1"},
		"quantity": {"1"},
	}

	req, err := http.NewRequest("POST", "/add-item", bytes.NewBufferString(formData.Encode()))

	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.api.ServeHTTP(w, req)
	query := w.Result().Header.Get("Location")
	m, _ := url.ParseQuery(query)
	s.Require().Equal(m.Get("/?error"), "invalid item name")
	s.Require().Equal(http.StatusFound, w.Code)

}

func (s *GinTestSuite) TestDeleteItem() {

	req, err := http.NewRequest("GET", "/remove-cart-item?cart_item_id=", nil)

	s.Require().NoError(err)
	w := httptest.NewRecorder()
	s.api.ServeHTTP(w, req)
	query := w.Result().Header.Get("Location")
	m, _ := url.ParseQuery(query)
	s.Require().Equal(m.Get("/?error"), "invalid cart item id")
	s.Require().Equal(http.StatusFound, w.Code)

}
