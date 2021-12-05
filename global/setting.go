package global

import (
	"gitee.com/gjianbo/web/pkg/logger"
	"gitee.com/gjianbo/web/pkg/setting"
)

var (
	ServetSettingS   *setting.ServetSettingS
	AppSettingS      *setting.AppSettingS
	DatabaseSettings *setting.DatabaseSettings
	Logger           *logger.Logger
)
