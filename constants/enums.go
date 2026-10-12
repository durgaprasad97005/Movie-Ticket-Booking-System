package constants

type UserRoleEnum string

const (
	SuperAdmin UserRoleEnum = "SUPER_ADMIN"
	Admin      UserRoleEnum = "ADMIN"
	User       UserRoleEnum = "USER"
)
