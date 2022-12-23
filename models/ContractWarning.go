package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"
)

type ContractWarning struct {
	Period    int64  `json:"period,omitempty"`
	Day_begin string `json:"day_begin,omitempty"`
	Day_end   string `json:"day_end,omitempty"`
	Colour    string `json:"colour,omitempty"`
}

func TambahWarning(column ContractWarning) ContractWarning {
	// buka koneksi
	db := config.CreateConnection()
	//tutup koneksi di akhir proses
	defer db.Close()

	//query insert
	sqlStatement := "INSERT INTO contract_warning (period, day_begin, day_end, colour, updated_at, created_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING period, day_begin, colour"

	var period, day_begin int64
	var colour string

	err := db.QueryRow(sqlStatement, column.Period, column.Day_begin, column.Day_end, column.Colour).Scan(&period, &day_begin, &colour)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data singgle record %v, %v, %v", period, day_begin, colour)

	return column
}

func AmbilSemuaWarning() ([]ContractWarning, error) {
	db := config.CreateConnection()

	defer db.Close()

	var ContractWarnings []ContractWarning

	sqlStatement := `SELECT period, day_begin, day_end, colour FROM contract_warning`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var column ContractWarning
		err = rows.Scan(&column.Period, &column.Day_begin, &column.Day_end, &column.Colour)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		ContractWarnings = append(ContractWarnings, column)
	}
	return ContractWarnings, err
}

func AmbilSatuWarning(period int64, day_begin int64, colour string) (ContractWarning, error) {
	db := config.CreateConnection()
	defer db.Close()

	var column ContractWarning
	row := db.QueryRow("SELECT period, day_begin, day_end, colour FROM contract_warning where period=$1 and day_begin=$2 and colour=$3", period, day_begin, `#`+colour)
	err := row.Scan(&column.Period, &column.Day_begin, &column.Day_end, &column.Colour)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari")
		return column, nil
	case nil:
		return column, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return column, nil
}

func UpdateWarning(period int64, day_begin int64, colour string, column ContractWarning) (int64, ContractWarning) {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE contract_warning SET period=$1, day_begin=$2, day_end=$3, colour=$4, updated_at=now() WHERE period=$5 and day_begin=$6 and colour=$7`
	res, err := db.Exec(sqlStatement, column.Period, column.Day_begin, column.Day_end, column.Colour, period, day_begin, `#`+colour)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)
	return rowsAffected, column
}

func DeleteWarning(period int64, day_begin int64, colour string) int64 {
	db := config.CreateConnection()
	defer db.Close()

	res, err := db.Exec(`DELETE FROM contract_warning where period=$1 and day_begin=$2 and colour=$3`, period, day_begin, `#`+colour)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)
	return rowsAffected
}
