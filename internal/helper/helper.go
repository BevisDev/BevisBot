package helper

import (
	"github.com/BevisDev/godev/utils/datetime"
	"time"
)

func GetNowDateTime() string {
	return datetime.ToString(time.Now(), datetime.DateTime)
}
