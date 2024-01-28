//#go:build integration

package service

import (
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"interview/pkg/core/port"
	"interview/pkg/core/service"
	"interview/pkg/infrastructure/db"
	"interview/pkg/infrastructure/repository"
	"interview/pkg/logger"

	"testing"
)

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}

type CartTestSuite struct {
	suite.Suite
	db          *gorm.DB
	cartService port.CartService
}

func (s *CartTestSuite) SetupSuite() {

	logger.Init("SetupSuite")
	lg := logger.General.Component("main")
	lg.Println("test started")
	db := db.GetDatabase(true, "")
	dbRepo := repository.NewRepository(db)
	cartService := service.NewCartService(service.CartServiceParam{
		CartItemRepo:   dbRepo,
		CartEntityRepo: dbRepo,
	})

	s.cartService = cartService
	s.db = db
}

func (s *CartTestSuite) TearDownSuite() {
	s.db.Exec("delete * from cart_items where 1)")
	s.db.Exec("delete * from carts where 1)")
}
