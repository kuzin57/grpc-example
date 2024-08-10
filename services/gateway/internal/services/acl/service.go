package acl

import "context"

var (
	_ = ACLService(&service{})
)

type ACLService interface {
	CheckAccess(ctx context.Context, token string) (bool, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}
