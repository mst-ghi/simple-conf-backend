package socket

import "video-conf/pkg/helpers"

type CuidGenerator struct {
	ID string
}

func (ig *CuidGenerator) NewID() string {
	return helpers.CUID()
}
