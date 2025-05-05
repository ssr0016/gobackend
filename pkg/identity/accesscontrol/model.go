package accesscontrol

type Permission struct {
	ID          int64   `db:"id" json:"id,omitempty"`
	Name        string  `db:"name" json:"name,omitempty"`
	Action      Action  `db:"action" json:"action,omitempty"`
	Description *string `db:"description" json:"description,omitempty"`
	CreatedAt   string  `db:"created_at" json:"created_at,omitempty"`
}
