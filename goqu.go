package main

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func GoquBuildSQLSuit() {
	GoquSelectSimple()
	GoquSelectWithPGPlaceholder()
	GoquInsertWithPGPlaceholder()
	GoquUpdateWithPGPlaceholder()
	GoquDeleteWithPGPlaceholder()
	GoquComplexSelectWithPGPlaceHolder()
}

// SELECT "id" FROM "tickets" WHERE (("subdomain_id" = 1) AND (("state" = 'open') OR ("state" = 'spam')))
// []
func GoquSelectSimple() {
	goqu.Select("id").
		From("tickets").
		Where(
			goqu.I("subdomain_id").Eq(1),
			goqu.Or(
				goqu.I("state").Eq("open"),
				goqu.I("state").Eq("spam"),
			),
		).ToSQL()
}

// SELECT "id" FROM "tickets" WHERE (("subdomain_id" = $1) AND (("state" = $2) OR ("state" = $3)))
// [1 open spam]
func GoquSelectWithPGPlaceholder() {
	d := goqu.Dialect("postgres")
	d.Select("id").
		From("tickets").
		Where(
			goqu.I("subdomain_id").Eq(1),
			goqu.Or(
				goqu.I("state").Eq("open"),
				goqu.I("state").Eq("spam"),
			),
		).Prepared(true).ToSQL()
}

// INSERT INTO "user" ("first_name", "last_name") VALUES ($1, $2), ($3, $4), ($5, $6)
// [Greg Farley Jimmy Stewart Jeff Jeffers]
func GoquInsertWithPGPlaceholder() {
	d := goqu.Dialect("postgres")

	d.Insert("user").
		Cols("first_name", "last_name").
		Vals(
			goqu.Vals{"Greg", "Farley"},
			goqu.Vals{"Jimmy", "Stewart"},
			goqu.Vals{"Jeff", "Jeffers"},
		).Prepared(true).ToSQL()
}

// UPDATE "users" SET "first_name"=$1,"last_name"=$2 WHERE ("first_name" = $3)
// [Greg Farley Gregory]
func GoquUpdateWithPGPlaceholder() {
	d := goqu.Dialect("postgres")

	d.From("users").Update().Set(
		goqu.Record{"first_name": "Greg", "last_name": "Farley"},
	).Where(
		goqu.C("first_name").Eq("Gregory"),
	).Prepared(true).ToSQL()
}

// DELETE FROM "users" WHERE ("last_name" IS NULL) []
func GoquDeleteWithPGPlaceholder() {
	d := goqu.Dialect("postgres")

	d.Delete("users").
		Where(goqu.Ex{"last_name": nil}).ToSQL()
}

// SELECT * FROM "users" AS "u" INNER JOIN "occupation" AS "occ" ON ("u"."id" = "occ"."user_id") LEFT JOIN "addresses" AS "addr" ON ("u"."id" = "addr"."user_id") WHERE (("occ"."occupation_name" = $1) AND ("addr"."type" = $2)) ORDER BY "u"."created_at" ASC LIMIT $3
// [Software engineer work 100]
func GoquComplexSelectWithPGPlaceHolder() {
	dPg := goqu.Dialect("postgres")

	ds := dPg.From(goqu.T("users").As("u")).
		InnerJoin(
			goqu.T("occupation").As("occ"),
			goqu.On(goqu.Ex{"u.id": goqu.I("occ.user_id")}),
		).
		LeftJoin(
			goqu.T("addresses").As("addr"),
			goqu.On(goqu.Ex{"u.id": goqu.I("addr.user_id")}),
		).
		Order(goqu.I("u.created_at").Asc()).
		Limit(100)

	ds = ds.Where(goqu.I("occ.occupation_name").Eq("Software engineer"))
	ds = ds.Where(goqu.Ex{"addr.type": "work"})

	ds.Prepared(true).ToSQL()
}
