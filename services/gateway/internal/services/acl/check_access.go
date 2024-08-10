package acl

import "context"

func (s *service) CheckAccess(ctx context.Context, token string) (bool, error) {
	return false, nil
}
