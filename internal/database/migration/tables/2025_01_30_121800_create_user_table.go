package tables

import (
	. "github.com/donbarrigon/new-project/internal/database/migration"
)

func create_user_table() *Table {

	table := NewTable(
		"user",
		BigIncrements(),
		String("name"),
		String("email"),
		String("password"),
		CreatedAt(),
		UpdatedAt(),
		DeletedAt(),
	)
	return table
}
