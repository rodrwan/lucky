package db

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

// QueryData bla bla
type QueryData struct {
	Description string `db:"description"`
	Extended    string `db:"extended_description"`
	Category    int    `db:"sub_category_id"`
}

// ReadDb connect, read and extract data
func ReadDb(db *sqlx.DB) (string, error) {
	file, err := os.Create("training_data.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	query := `
		SELECT
			COALESCE(t.description, '') AS description,
			COALESCE(t.extended_description, '') AS extended_description,
			t.sub_category_id AS sub_category_id
		FROM
			transactions t
		JOIN sub_categories sc ON t.sub_category_id = sc.id
		WHERE
		(
			t.categorized_by = 'user'
		OR
			t.categorized_by = 'category_match'
		OR
			t.categorized_by = 'admin'
		)
		AND NOT
		(
			t.sub_category_id = 1
		OR
			t.sub_category_id = 76
		)
		AND sc.category_id != 20
		AND t.ignored = 'f'
		AND NOT
		(
			t.description = 'COMPRA EN EL COMERCIO'
		OR
			t.description ~ 'Movimiento fecha'
		OR
			t.description ~ 'Transf\.'
		);
	`

	rows, err := db.Queryx(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	w := bufio.NewWriter(file)
	for rows.Next() {
		e := QueryData{}

		if err := rows.StructScan(&e); err != nil {
			return "", err
		}

		d := fmt.Sprintf("%d#%s %s", e.Category, e.Description, e.Extended)
		// fmt.Printf("%s\n", d)
		fmt.Fprintln(w, d)
	}

	w.Flush()
	return "hola", nil
}
