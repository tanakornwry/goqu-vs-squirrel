package main

import (
	sq "github.com/Masterminds/squirrel"
)

func SQBuildSQLSuit() {
	SQSelectSimple()
	SQSelectWithPgPlaceholder()
	SQInsertWithPGPlaceholder()
	SQUpdateWithPGPlaceholder()
	SQDeleteWithPGPlaceholder()
	SQComplexSelectWithPGPlaceHolder()
}

// SELECT id FROM tickets WHERE subdomain_id = ? AND (state = ? OR state = ?)
// [1 open spam]
func SQSelectSimple() {
	sq.Select("id").
		From("tickets").
		Where(sq.Eq{"subdomain_id": 1}).
		Where(
			sq.Or{
				sq.Eq{"state": "open"},
				sq.Eq{"state": "spam"},
			},
		).ToSql()
}

// SELECT id FROM tickets WHERE subdomain_id = $1 AND (state = $2 OR state = $3)
// [1 open spam]
func SQSelectWithPgPlaceholder() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	psql.Select("id").
		From("tickets").
		Where(
			sq.Eq{"subdomain_id": 1},
		).
		Where(
			sq.Or{
				sq.Eq{"state": "open"},
				sq.Eq{"state": "spam"},
			},
		).ToSql()
}

// INSERT INTO users (first_name,last_name) VALUES ($1,$2),($3,$4),($5,$6)
// [Greg Farley Jimmy Stewart Jeff Jeffers]
func SQInsertWithPGPlaceholder() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	psql.Insert("users").Columns("first_name", "last_name").
		Values("Greg", "Farley").
		Values("Jimmy", "Stewart").
		Values("Jeff", "Jeffers").
		ToSql()
}

// UPDATE users SET first_name = $1, last_name = $2 WHERE first_name = $3
// [Greg Farley Gregory]
func SQUpdateWithPGPlaceholder() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	psql.Update("users").
		Set("first_name", "Greg").
		Set("last_name", "Farley").
		Where(sq.Eq{"first_name": "Gregory"}).
		ToSql()
}

// DELETE FROM users WHERE last_name IS NULL []
func SQDeleteWithPGPlaceholder() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	psql.Delete("users").Where(sq.Eq{"last_name": nil}).ToSql()
}

// SELECT * FROM users AS u INNER JOIN occupation AS occ ON u.id = occ.user_id LEFT JOIN addresses AS addr ON u.id = addr.user_id WHERE occ.occupation_name = $1 AND addr.type = $2 ORDER BY u.created_at ASC LIMIT 100
// [Software engineer work]
func SQComplexSelectWithPGPlaceHolder() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	ds := psql.Select("*").
		From("users AS u").
		InnerJoin("occupation AS occ ON u.id = occ.user_id").
		LeftJoin("addresses AS addr ON u.id = addr.user_id").
		OrderBy("u.created_at ASC").
		Limit(100)

	ds = ds.Where(sq.Eq{"occ.occupation_name": "Software engineer"})
	ds = ds.Where(sq.Eq{"addr.type": "work"})

	ds.ToSql()
}
