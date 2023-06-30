package websiteaccess

import (
	"time"
)

type WebsiteAccess interface {
	CheckAvailability(url string) (bool, time.Time)
}

type WebsiteAccessExec struct{}
