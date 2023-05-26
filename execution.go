package main

import (
	"database/sql"
	"fmt"
)

type DupeRow struct {
	RowId int
	Row   *NilaiRow
}

func isExistOnArray(arr []string, id string) bool {
	for _, v := range arr {
		if id == v {
			return true
		}
	}

	return false
}

func DeleteRow(db *sql.DB, dupe *DupeRow) bool {
	_, err := db.Exec(DEL_NILAI_SQL, dupe.RowId)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func ExecutionDupes(db *sql.DB) []DupeRow {
	lists := []string{}
	dupes := []DupeRow{}
	tmpCounter := 0

	rows, err := db.Query(GET_ALL_SQL)
	if err != nil {
		fmt.Println(err.Error())
		return []DupeRow{}
	}

	defer rows.Close()
	for rows.Next() {
		row := &NilaiRow{}

		if err = rows.Scan(&row.Id, &row.Lesson, &row.S1, &row.S2, &row.S3, &row.S4, &row.S5, &row.ArchiveID); err != nil {
			fmt.Println("Error while scan: ", err.Error())
			break
		}

		if !isExistOnArray(lists, row.ArchiveID) {
			tmpCounter = 0
			lists = append(lists, row.ArchiveID)
			tmpCounter++
		} else {
			tmpCounter++
			if tmpCounter > 5 {
				dupes = append(dupes, DupeRow{
					RowId: row.Id,
					Row:   row,
				})
			}
		}
	}

	return dupes
}
