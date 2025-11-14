package helper

import (
	"time"

	"github.com/BevisDev/godev/utils/datetime"
)

func GetNowDateTime() string {
	return datetime.ToString(time.Now(), datetime.DateTime)
}
