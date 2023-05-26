package main

import (
	"fmt"
	"os"
)

func main() {
	// extract args from cmd-line
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Error: missing dsn")
		os.Exit(1)
	}

	dsn := args[0]

	// open connection
	db := InitializeDB(dsn)
	if db == nil {
		os.Exit(1)
	}

	defer db.Close()

	dupesArrays := ExecutionDupes(db)
	for _, dupe := range dupesArrays {
		successDelete := DeleteRow(db, &dupe)
		if successDelete {
			fmt.Printf("Sukses menghapus %d dari %s - %s\n", dupe.RowId, dupe.Row.ArchiveID, dupe.Row.Lesson)
		} else {
			fmt.Printf("Gagal menghapus %d dari %s - %s\n", dupe.RowId, dupe.Row.ArchiveID, dupe.Row.Lesson)
		}
	}
}
