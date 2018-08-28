package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	jsonq "github.com/jmoiron/jsonq"
	gorequest "github.com/parnurzeal/gorequest"
)

var wg = sync.WaitGroup{}

func main() {
	csv_files := []string{
		"stations.csv", "gas_stations.csv", "prices.csv", "gas_prices.csv", "", "all.csv",
	}
	fmt.Print("일 목록 선택 (0: 주유소입력, 1:충전소입력, 2:주유가격갱신, 3:충전가격갱신, 4: 미기재 주소좌표변환)\n키보드로 입력, 5: 무작위 전체 업데이트 ")

	var filenum int
	fmt.Scanf("%d", &filenum)
	log.Print("Start")

	db, err := sql.Open("mysql", "root:333157@tcp(saeparan.com:3306)/oil")
	if err != nil {
		log.Println(err)
	}

	if filenum == 4 {
		var id int
		var address string
		rows, err := db.Query("select id, address from gas_station where x is null and y is null limit 20")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		sql := "UPDATE gas_station SET old_address=?, x=?, y=?, daum_name=? WHERE id=?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			fmt.Println("Query Eerror")
			fmt.Println(err)
			os.Exit(1)
		}

		for rows.Next() {
			err := rows.Scan(&id, &address)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, address)

			wg.Add(1)
			go addressTrans(stmt, id, address)
		}

		wg.Wait()
		os.Exit(1)
	}

	csvfile, err := os.Open(csv_files[filenum])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sql := ""
	if filenum == 0 || filenum == 1 {
		sql = "INSERT gas_station SET station_key=?, address_head=?, name=?, vendor=?, address=?, phone=?, self=?;"
	} else if filenum == 2 || filenum == 3 {
		sql = "UPDATE gas_station SET price1=?, price2=?, price3=?, price4=? WHERE station_key=?"
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("Query Eerror")
		fmt.Println(err)
		os.Exit(1)
	}

	for i, data := range rawCSVdata {
		if i < 3 {
			continue
		}

		if filenum == 0 || filenum == 1 {
			is_self := 0
			if data[6] == "셀프" {
				is_self = 1
			}

			_, err = stmt.Exec(data[0], data[1], data[2], data[3], data[4], data[5], is_self)
		} else if filenum == 2 {
			_, err = stmt.Exec(data[6], data[7], data[8], data[9], data[0])
		} else if filenum == 3 {
			_, err = stmt.Exec(data[6], nil, nil, nil, data[0])
		} else if filenum == 5 {
			log.Println(data)
			if i == 20 {
				break
			}
		}
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func addressTrans(stmt *sql.Stmt, id int, address string) {
	goreq := gorequest.New()
	resp, body, _ := goreq.Get("https://apis.daum.net/local/geo/addr2coord?apikey=6eb28058ac2e05704a7566eace2c4e36&q=" + address + "&output=json").End()
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	newAddress, _ := jq.String("channel", "item", "0", "newAddress")
	lat, _ := jq.Float("channel", "item", "0", "lat")
	lng, _ := jq.Float("channel", "item", "0", "lng")
	buildingAddress, _ := jq.String("channel", "item", "0", "buildingAddress")
	if resp.StatusCode == 200 {
		_, err := stmt.Exec(newAddress, lat, lng, buildingAddress, id)
		if err != nil {
			log.Println(err)
		}
		log.Println(lat, lat, buildingAddress, newAddress, resp.StatusCode)
	} else {
		log.Println(resp.StatusCode)
		log.Println(resp)
	}

	defer wg.Done()
	//db.Query("UPDATE gas_station SET ")
}
