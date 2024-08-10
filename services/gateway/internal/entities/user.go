package entities

type User struct {
	ID string `db:"user_id"`

	Login          string `db:"login"`
	HashedPassword string `db:"hashed_password"`
}
