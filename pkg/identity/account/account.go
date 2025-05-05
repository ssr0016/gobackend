package account

type Account int

const (
	Admin Account = iota + 1
	User
)
