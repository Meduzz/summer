package framework

import "github.com/Meduzz/summer/api"

type (
	dummyIdentity struct{}
)

var _ Identity = (*dummyIdentity)(nil)

func NewDummyIdentity() Identity {
	return &dummyIdentity{}
}

func (d *dummyIdentity) Method(it string) bool {
	return true
}

func (d *dummyIdentity) Request(it *api.Request) bool {
	return true
}
