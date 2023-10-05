package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDataAccess(t *testing.T) {
	t.Run("Insert data", func(t *testing.T) {
		result, err := db.Exec("insert into banners values ($1, $2), ($3, $4)", 99, "banner 99", 100, "banner 100")
		require.NoError(t, err)
		rowsCount, err := result.RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(2), rowsCount)
	})

	t.Run("Select data", func(t *testing.T) {
		expected := []Entity{
			{id: 99, description: "banner 99"},
			{id: 100, description: "banner 100"},
		}
		rows, err := db.Query("select * from banners where id in ($1, $2)", 99, 100)
		require.NoError(t, err)
		for _, entity := range expected {
			var actual Entity
			rows.Next()
			err := rows.Scan(&actual.id, &actual.description)
			require.NoError(t, err)
			require.EqualValues(t, entity, actual)
		}
	})

	t.Run("Delete data", func(t *testing.T) {
		result, err := db.Exec("delete from banners where id in ($1, $2)", 99, 100)
		require.NoError(t, err)
		rowsCount, err := result.RowsAffected()
		require.NoError(t, err)
		require.Equal(t, int64(2), rowsCount)
	})
}
