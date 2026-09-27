package github

import (
	"context"
	"fmt"
)

const (
	PermissionAdmin = "admin"
	PermissionWrite = "write"
	PermissionRead  = "read"
	PermissionNone  = "none"
)

func (c *Client) ActorPermission(ctx context.Context, login string) (string, error) {
	if login == "" {
		return PermissionNone, nil
	}
	level, _, err := c.Raw.Repositories.GetPermissionLevel(ctx, c.Owner, c.Repo, login)
	if err != nil {
		return "", fmt.Errorf("github: get permission for %s: %w", login, err)
	}
	perm := level.GetPermission()
	switch perm {
	case PermissionAdmin, PermissionWrite, PermissionRead:
		return perm, nil
	default:
		return PermissionNone, nil
	}
}

func HasPermission(level, minimum string) bool {
	rank := map[string]int{PermissionNone: 0, PermissionRead: 1, PermissionWrite: 2, PermissionAdmin: 3}
	return rank[level] >= rank[minimum]
}
