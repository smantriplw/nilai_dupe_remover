package main

// sql syntaxes
var (
	GET_ALL_SQL   = "SELECT * FROM nilai_semester"
	GET_NILAI_SQL = "SELECT * FROM nilai_semester WHERE archive_id = ?"
	DEL_NILAI_SQL = "DELETE FROM nilai_semester WHERE id = ? "
)
