//go:build integration

package sql

import (
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"interview/pkg/core/port"
	db2 "interview/pkg/infrastructure/db"
	"interview/pkg/infrastructure/repository"
	"interview/pkg/logger"

	"testing"
)

func TestMongoTestSuite(t *testing.T) {
	suite.Run(t, new(SqlTestSuite))
}

type SqlTestSuite struct {
	suite.Suite
	db             *gorm.DB
	cartItemRepo   port.CartItemRepository
	cartEntityRepo port.CartEntityRepository
}

func (s *SqlTestSuite) SetupSuite() {

	logger.Init("SetupSuite")
	lg := logger.General.Component("main")
	lg.Println("test started")
	dbRepo := db2.GetDatabase(true, "")

	repo := repository.NewRepository(dbRepo)

	s.cartItemRepo = repo
	s.cartEntityRepo = repo
	s.db = dbRepo
}

func (s *SqlTestSuite) TearDownSuite() {
	s.db.Exec("delete * from cart_items where 1)")
	s.db.Exec("delete * from carts where 1)")
}
