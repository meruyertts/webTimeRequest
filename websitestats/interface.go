package websitestats

import (
	"webTimeRequest/statrecords"
	"webTimeRequest/websiteaccess"
	"webTimeRequest/websiterecords"
)

type WebsiteStats struct {
	WebsiteAccess websiteaccess.WebsiteAccess
	WebsiteInfo   websiterecords.WebsiteInfo
	StatInfo      statrecords.StatInfo
}
