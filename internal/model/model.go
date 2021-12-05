package model

import (
	"fmt"

	"gitee.com/gjianbo/web/global"
	"gitee.com/gjianbo/web/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint32 `json:"id"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modifie_by,omitempty"`
	CreatedOn  uint32 `json:"created_on,omitempty"`
	ModifiedOn uint32 `json:"modified_on,omitempty"`
	DeletedOn  uint32 `json:"deleted_on,omitempty"`
	IsDel      uint8  `json:"is_del,omitempty"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"

	//   "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"

	// 数据库链接格式 : root:1234567@tcp(127.0.0.1:3306)/blog_service?charset=utf8&parseTime=true&loc=Local

	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))

	if err != nil {
		return nil, err
	}
	if global.ServetSettingS.RunMode == "debug" {
		db.LogMode(true)
	}

	//db.SingularTable(true)
	//在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。即，如果有一个user表，那么如果你定义的结构体名为：User，gorm会默认表名为users而不是user。
	//db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)

	//
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil

}
