package dao

import (
	"time"

	"database/sql"

	"github.com/akatsukisun2020/fancy_painter/common/mysql"
	"github.com/akatsukisun2020/fancy_painter/server/fancy_painter/config"
)

// Users 用户资料
type Users struct {
	ID            int
	Confirmed     int
	Name          string
	AvatarUrl     string
	Phone         string
	Email         string
	Password_hash string
	Role_id       int
	UserId        string
	CreateTime    time.Time
	UpdateTime    time.Time
}

var gFancyPainterDB *sql.DB

// InitFancyPainterDB 初始化DB
func InitFancyPainterDB() error {
	if gFancyPainterDB == nil {
		var err error
		gFancyPainterDB, err = mysql.NewMYSQLCli(config.GetFancyPainterDBConfig().Dsn,
			config.GetFancyPainterDBConfig().MaxIdleConns,
			config.GetFancyPainterDBConfig().MaxOpenConns,
			time.Duration(config.GetFancyPainterDBConfig().MaxLifetime)*time.Second).OpenDB()
		return err
	}
	return nil
}
