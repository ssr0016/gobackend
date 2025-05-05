package accesscontrol

type Action string

const (
	//User
	ActionReadUser   Action = "read_user"
	ActionCreateUser Action = "create_user"
	ActionLookupUser Action = "lookup_user"
	ActionUpdateUser Action = "update_user"

	//Role
	ActionReadRole   Action = "read_role"
	ActionCreateRole Action = "create_role"
	ActionLookupRole Action = "lookup_role"
	ActionUpdateRole Action = "update_role"
)

const (
	UserGroup string = "User"
)
