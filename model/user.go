package model

import "github.com/gocraft/dbr/v2"

type (
	// User struct
	User struct {
		ID        int64        `db:"id"`
		IsAuthed  bool         `db:"is_authed"`
		Name      string       `db:"name"`
		DeletedAT dbr.NullTime `db:"deleted_at"`
	}
)
