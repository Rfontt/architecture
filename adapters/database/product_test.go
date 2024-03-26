package db

var DB *sqlDB

fun setup() {
	Db, _ = sql.Open("sqlite3", ":memory")
}