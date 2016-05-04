package db

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Data bla bla
type Data struct {
	Max int `db:"max"`
}

// CategoriesData bla bl
type CategoriesData struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
}

// Categories from db
func Categories(db *sqlx.DB) (map[uint]string, error) {
	m := make(map[uint]string)
	file, err := os.Create("labels.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	query := `
		SELECT id, name FROM sub_categories WHERE NOT sub_categories.category_id = 20;
	`

	d := CategoriesData{}
	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	w := bufio.NewWriter(file)
	for rows.Next() {
		err := rows.StructScan(&d)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		m[d.ID] = d.Name
		sc := fmt.Sprintf("%d#%s", d.ID, d.Name)
		// fmt.Printf("%s\n", d)
		fmt.Fprintln(w, sc)
	}

	w.Flush()
	return m, nil
}

// Load ...
func Load(subCatPath string) map[uint]string {
	f, err := os.Open(subCatPath)
	if err != nil {
		panic("cant open sub category file")
	}
	defer f.Close()

	m := make(map[uint]string)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		splitedStr := strings.Split(line, "#")
		id, name := splitedStr[0], splitedStr[1]
		i, err := strconv.Atoi(id)
		if err != nil {
			panic("cannot convert string to int")
		}
		m[uint(i)] = name
	}

	return m
}
