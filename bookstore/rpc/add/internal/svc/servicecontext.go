package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gozero/bookstore/rpc/add/internal/config"
	"gozero/bookstore/rpc/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.BookModel   // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
