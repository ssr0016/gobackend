package migration

import "backend/pkg/infra/storage/migrator"

func addUserMigration(mg *migrator.Migrator) {
	userTable := migrator.Table{
		Name: "user",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "uuid", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "first_name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "middle_name", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "last_name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "login_name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "password", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "status", Type: migrator.DB_TinyInt, Nullable: false, Default: "1"},
			{Name: "email", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "salt", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "created_by", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "created_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "updated_by", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "updated_at", Type: migrator.DB_DateTime, Nullable: true},
		},
		Indices: []*migrator.Index{
			{Cols: []string{"login_name"}, Type: migrator.UniqueIndex},
		},
	}

	mg.AddMigration("create user table", migrator.NewAddTableMigration(userTable))
	mg.AddMigration("add index user.login_name", migrator.NewAddIndexMigration(userTable, userTable.Indices[0]))
}
