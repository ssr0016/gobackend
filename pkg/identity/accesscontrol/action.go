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

	//Category
	ActionReadCategory   Action = "read_category"
	ActionCreateCategory Action = "create_category"
	ActionLookupCategory Action = "lookup_category"
	ActionUpdateCategory Action = "update_category"
)

const (
	UserGroup    string = "User"
	CategoryUser string = "User Actions"
	CategoryRole string = "Role Actions"

	CategoryGroup   string = "Product"
	CategoryProduct string = "Product Actions"
)

type PermissionCategory struct {
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

type AccessRole struct {
	Name                 string               `json:"name"`
	PermissionCategories []PermissionCategory `json:"permission_categories"`
}

type AccessControlList struct {
	Roles []AccessRole `json:"access_roles"`
}

var (
	AccessControls = AccessControlList{
		Roles: []AccessRole{
			{
				Name: UserGroup,
				PermissionCategories: []PermissionCategory{
					{
						Name: CategoryUser,
						Permissions: []Permission{
							{Action: ActionCreateUser},
							{Action: ActionLookupUser},
							{Action: ActionReadUser},
							{Action: ActionUpdateUser},
						},
					},
					{
						Name: CategoryRole,
						Permissions: []Permission{
							{Action: ActionCreateRole},
							{Action: ActionLookupRole},
							{Action: ActionReadRole},
							{Action: ActionUpdateRole},
						},
					},
				},
			},
			{
				Name: CategoryGroup,
				PermissionCategories: []PermissionCategory{
					{
						Name: CategoryProduct,
						Permissions: []Permission{
							{Action: ActionCreateCategory},
							{Action: ActionLookupCategory},
							{Action: ActionReadCategory},
							{Action: ActionUpdateCategory},
						},
					},
				},
			},
		},
	}
)
