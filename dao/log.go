package dao

import (
	"github.com/stiwedung/libgo/log"
	xlog "xorm.io/xorm/log"
)

var _ xlog.Logger = &logger{}

type logger struct {
	*log.Logger
	isShowSQL bool
}

func (logger *logger) Level() xlog.LogLevel {
	return xlog.LOG_UNKNOWN
}

func (logger *logger) SetLevel(l xlog.LogLevel) {}

func (logger *logger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		return
	}
	logger.isShowSQL = show[0]
}

func (logger *logger) IsShowSQL() bool {
	return logger.isShowSQL
}
