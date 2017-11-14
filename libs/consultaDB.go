package consultaDespachos

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

type Year struct {
	Year int `json:"Year"`
}

type DigitYear struct {
	YearMonth string `json:"YearMonth"`
	DigitOk   int    `json:"DigitOk"`
}

func GetYears() (yearList []Year) {
	condb, errdb := sql.Open("mssql", "server=192.168.169.208\\sqlexpress;user id=sa;password=ALpha2000;database=Bank;connection timeout=30")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	rows, err := condb.Query("select distinct year(FechaWsDok) as year from despachos where FechaWsDok is not null order by year(FechaWsDok) desc")
	if err != nil {
		log.Fatal(err)
	}

	var yearValue int
	for rows.Next() {
		err := rows.Scan(&yearValue)
		if err != nil {
			log.Fatal(err)
		}

		yearList = append(yearList, Year{Year: yearValue})

	}

	defer condb.Close()

	return
}

func GetYear(year int) (digitYearLst []DigitYear) {

	condb, errdb := sql.Open("mssql", "server=192.168.169.208\\sqlexpress;user id=sa;password=ALpha2000;database=Bank;connection timeout=30")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	rows, err := condb.Query("exec sp_DespachosMensual '" + strconv.Itoa(year) + "'")
	if err != nil {
		log.Fatal(err)
	}

	var yearValue string
	var digitOk int

	for rows.Next() {
		err := rows.Scan(&yearValue, &digitOk)
		if err != nil {
			log.Fatal(err)
		}

		digitYearLst = append(digitYearLst, DigitYear{YearMonth: yearValue, DigitOk: digitOk})

	}

	defer condb.Close()

	return

}

func GetYearMonth(yearMonth string) (digitYearLst []DigitYear) {

	condb, errdb := sql.Open("mssql", "server=192.168.169.208\\sqlexpress;user id=sa;password=ALpha2000;database=Bank;connection timeout=30")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	rows, err := condb.Query("exec sp_DespachosDiarios '" + yearMonth + "'")
	if err != nil {
		log.Fatal(err)
	}

	var yearValue string
	var digitOk int

	for rows.Next() {
		err := rows.Scan(&yearValue, &digitOk)
		if err != nil {
			log.Fatal(err)
		}

		digitYearLst = append(digitYearLst, DigitYear{YearMonth: yearValue, DigitOk: digitOk})

	}
	defer condb.Close()
	return

}
