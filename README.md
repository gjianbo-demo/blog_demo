# 三方包
## web
gin 框架
```
go get -u github.com/gin-gonic/gin@v1.7.7

```

## 配置
viper
```
go get -u github.com/spf13/viper@v1.9.0
```

## 数据库
gorm
```
go get -u github.com/jinzhu/gorm@v1.9.16
```
连接数据库

```
import (
	"fmt"

	"gitee.com/gjianbo/web/global"
	"gitee.com/gjianbo/web/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


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

    // 设置数据库连接数
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
```

## 日志
```
go get -u gopkg.in/natefinch/lumberjack.v2
```

## 文档
```
go get -u github.com/swaggo/swag/cmd/swag@v.1.7.6
go get -u github.com/swaggo/gin-swagger@1.3.3
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template

```


## 接口验证
```
go get -u github.com/go-playground/validator/v10
```

