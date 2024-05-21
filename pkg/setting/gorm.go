package setting

import (
	"eat_box/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDBEngine(databasesetting *MysqlSettings) (*gorm.DB, error) {

	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databasesetting.Username,
		databasesetting.Password,
		databasesetting.Host,
		databasesetting.Port,
		databasesetting.Dbname)), &gorm.Config{Logger: newlogger})

	if err != nil {
		return nil, err
	}

	//自动迁移
	err = db.AutoMigrate(&model.User{},
		&model.Business{},
		&model.BusinessScore{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
