package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	iconv "gopkg.in/iconv.v1"

	"github.com/dustin/go-humanize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron"
)

type Station struct {
	gorm.Model
	Id uint `gorm:"primary_key"`
	//StationKey  string
	AddressHead string
	Name        string
	Vendor      string
	Phone       string
	Self        int
	StationType int
	Price       int64 `gorm:"-"`
	Price1      int64
	Price2      int64
	Price3      int64
	Price4      int64
	PriceStr    string `gorm:"-"`
	Wash        int
	Garage      int
	Cvs         int
	X           float64
	Y           float64
	OldAddress  string
	Address     string
	DaumName    string
	Distance    float64 `gorm:"-"`
}

type Charger struct {
	gorm.Model
	Id            uint `gorm:"primary_key"`
	Name          string
	Location      string
	Sido          string
	Restday       string
	Starttime     string
	Endtime       string
	Slow          string
	Quick         string
	Quicktype     string
	SlowQuantity  int64
	QuickQuantity int64
	ParkFee       string
	Address       string
	OldAddress    string
	ManageCorp    string
	ManageTel     string
	X             float64
	Y             float64
	OfficialDate  string
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root@/oil?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000", "https://oilzip.saeparan.com"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	c := cron.New()
	c.AddFunc("1 2 * * * *", updateStations)
	c.AddFunc("1 14 * * * *", updateStations)
	c.Start()

	e.GET("/stations/:lat/:lng/:fuelType/:distance/:wash", getStations)
	e.GET("/chargers/:lat/:lng/:distance/:wash", getChargers)

	e.Logger.Fatal(e.Start(":1323"))
}

func getChargers(c echo.Context) error {
	lat := c.Param("lat")
	lng := c.Param("lng")
	distanceGet := c.Param("distance")
	distance, _ := strconv.Atoi(distanceGet)
	if distance > 11 {
		c.String(404, "Wrong Request.")
	}

	sql := db.Model(&Charger{}).Select("chargers.*, ( 6371 * acos( cos( radians(?) ) * cos( radians( x ) ) * cos( radians( y ) - radians(?) ) + sin( radians(?) ) * sin( radians( x ) ) ) ) AS distance", lat, lng, lat).Having("distance <= ?", distance)
	rows, _ := sql.Rows()
	defer rows.Close()

	var chargers []Charger
	for rows.Next() {
		var charger Charger
		db.ScanRows(rows, &charger)

		chargers = append(chargers, charger)
	}

	return c.JSON(http.StatusOK, chargers)
}

func getStations(c echo.Context) error {
	lat := c.Param("lat")
	lng := c.Param("lng")
	fuelTypeGet := c.Param("fuelType")
	fuelType, _ := strconv.Atoi(fuelTypeGet)
	distanceGet := c.Param("distance")
	distance, _ := strconv.Atoi(distanceGet)
	washGet := c.Param("wash")
	wash, _ := strconv.Atoi(washGet)
	stationType := 1
	//db.Where("name LIKE ?", "%SK%").Find(&station)
	if distance > 11 {
		c.String(404, "Wrong Request.")
	}

	if fuelType == 4 {
		stationType = 2
	}

	sql := db.Model(&Station{}).Select("stations.*, ( 6371 * acos( cos( radians(?) ) * cos( radians( x ) ) * cos( radians( y ) - radians(?) ) + sin( radians(?) ) * sin( radians( x ) ) ) ) AS distance", lat, lng, lat).Having("distance <= ?", distance).Where("station_type=?", stationType)
	if fuelType == 2 {
		sql = sql.Where("price1 > ?", 0)
	}
	if wash == 1 {
		sql = sql.Where("wash = ?", wash)
	}

	rows, _ := sql.Rows()
	defer rows.Close()

	var stations []Station
	for rows.Next() {
		var priceTmp int64
		var station Station
		db.ScanRows(rows, &station)

		switch fuelType {
		case 1:
			priceTmp = station.Price2
			break
		case 2:
			priceTmp = station.Price1
			break
		case 4:
			priceTmp = station.Price1
			break
		case 3:
			priceTmp = station.Price3
			break
		}

		if priceTmp == 0 {
			continue
		}

		station.PriceStr = humanize.Comma(priceTmp)
		station.Price = priceTmp
		station.Distance = Round(station.Distance, 0.5, 1)
		stations = append(stations, station)
	}

	return c.JSON(http.StatusOK, stations)
}

func updateStations() {
	downloadFromUrl("A")
	applyStations()

	downloadFromUrl("B")
	applyStations()
}

func applyStations() {
	log.Println("db apply job... start")
	f, _ := os.Open("./origin.csv")
	r := csv.NewReader(bufio.NewReader(f))
	cnt := 0
	station := Station{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if cnt > 2 {
			rl := len(record)
			if rl > 3 {
				price1, _ := strconv.ParseInt(record[6], 10, 64)
				price2 := int64(0)
				price3 := int64(0)
				price4 := int64(0)

				if rl > 7 {
					price2, _ = strconv.ParseInt(record[7], 10, 64)

					price3, _ = strconv.ParseInt(record[8], 10, 64)

					price4, _ = strconv.ParseInt(record[9], 10, 64)
				}
				db.Model(&station).Where("station_key = ?", record[0]).Updates(map[string]interface{}{"price1": price1, "price2": price2, "price3": price3, "price4": price4})
			}
		}
		cnt++
	}
	log.Println("db apply complete.")
}

func downloadFromUrl(LPG_CD string) {
	log.Println("Start download LPG Type : ", LPG_CD)
	urlStr := "https://www.opinet.co.kr/user/main/main_download_excel.do"

	fileName := "temp.csv"
	fmt.Println("Downloading", urlStr, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	t := time.Now()
	dates := t.Format("20060102")

	// LPG_CD A기름 BLPG
	a := url.Values{"LPG_CD": {LPG_CD}, "DATE_DIV_CD": {""}, "PAGE_DIV": {"PAGE_DIV_5"}, "SIDO_NM": {"%EC%8B%9C%2F%EB%8F%84"}, "SIGUN_NM": {"%EC%8B%9C%2F%EA%B5%B0%2F%EA%B5%AC"}, "rdo1": {"A"}, "rdo2": {"A"}, "rdo3": {"A"}, "rdo4": {"X"}, "START_DT": {dates}, "END_DT": {dates}}
	response, err := http.PostForm(urlStr, a)
	if err != nil {
		fmt.Println("Error while downloading", urlStr, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", urlStr, "-", err)
		return
	}

	cd, err := iconv.Open("utf-8", "euc-kr") // convert gbk to utf8
	if err != nil {
		fmt.Println("iconv.Open failed!")
		return
	}
	defer cd.Close()

	input, _ := os.Open(fileName) // eg. input := os.Stdin || input, err := os.Open(file)
	bufSize := 0                  // default if zero

	r := iconv.NewReader(cd, input, bufSize)

	output2, err := os.Create("origin.csv")
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output2.Close()

	_, err = io.Copy(output2, r)
	if err != nil {
		fmt.Println("\nio.Copy failed:", err)
		return
	}

	os.Remove(fileName)

	fmt.Println("download complete.", n)
}

func Round(val float64, roundOn float64, places int) float64 {

	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)

	var round float64
	if val > 0 {
		if div >= roundOn {
			round = math.Ceil(digit)
		} else {
			round = math.Floor(digit)
		}
	} else {
		if div >= roundOn {
			round = math.Floor(digit)
		} else {
			round = math.Ceil(digit)
		}
	}

	return round / pow
}
