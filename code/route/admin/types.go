package admin

import (
	"time"
)

type Times struct {
	Name      string    `json:"name" form:"name"`
	City      string    `json:"city" form:"city"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
}

type TimeResp struct {
	Name      string `json:"name" form:"name"`
	City      string `json:"city" form:"city"`
	CreatedAt string `json:"createdAt" form:"createdAt"`
}
