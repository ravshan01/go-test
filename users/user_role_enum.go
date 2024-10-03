//go:generate go-enum -f=user_role_enum.go --marshal
package users

type UserRole int

const (
	UserRootRole UserRole = iota
	UserAdminRole
	UserUserRole
)

func (role UserRole) String() string {
	return [...]string{"root", "admin", "user"}[role]
}
