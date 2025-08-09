package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	svg "github.com/ajstarks/svgo"
)

var Db *sql.DB
var Conf Config

const dimension = 4
const max = dimension * dimension

type Config struct {
	Svg ConfigSvg
	Db  ConfigDb
}

type ConfigSvg struct {
	BaseWidth  int
	BaseHeight int
	X1         int
	Y1         int
}

type ConfigDb struct {
	User     string
	Password string
	DBName   string
}

func init() {
	Conf = Config{
		Svg: ConfigSvg{
			BaseWidth:  120,
			BaseHeight: 120,
			X1:         10,
			Y1:         90,
		},
		Db: ConfigDb{
			User:     "",
			Password: "",
			DBName:   "",
		},
	}
	Db, _ = sql.Open("mysql", Conf.Db.User+":"+Conf.Db.Password+"@/"+Conf.Db.DBName)
}

func main() {
	rows, err := Db.Query("SELECT * FROM 4x4 ORDER BY 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16")
	if err != nil {
		fmt.Println("Error sql select data: ", err)
	}
	var squares [][max]int
	for rows.Next() {
		var square [max]int
		rows.Scan(&square[0], &square[1], &square[2], &square[3], &square[4], &square[5], &square[6], &square[7], &square[8], &square[9], &square[10], &square[11], &square[12], &square[13], &square[14], &square[15])
		squares = append(squares, square)
	}
	rows.Close()

	for k, square := range squares {
		file, err := os.Create("4x4/" + strconv.Itoa(k+1) + ".svg")
		if err != nil {
			panic(err)
		}
		defer file.Close() // Ensure the file is closed when the function exits

		// Create a new SVG instance, directing output to the file
		canvas := svg.New(file)
		canvas.Start(Conf.Svg.BaseWidth*dimension, Conf.Svg.BaseHeight*dimension)

		var opacityStep float64 = 2 / (float64(max - 1))
		for i, v := range square {
			var opacity string
			if v <= max/2 {
				opacity = strconv.FormatFloat(opacityStep*float64(v), 'f', 2, 64)
			} else {
				opacity = strconv.FormatFloat(2-opacityStep*float64(v-2), 'f', 2, 64)
			}
			canvas.Rect(i%dimension*Conf.Svg.BaseWidth, i/dimension*Conf.Svg.BaseHeight, 120, 120, "fill=\"white\" fill-opacity=\""+opacity+"\"")
			canvas.Text(Conf.Svg.X1+i%dimension*Conf.Svg.BaseWidth, i/dimension*Conf.Svg.BaseHeight+Conf.Svg.Y1, strconv.Itoa(v), "stroke=\"#000\" font-family=\"'Domine'\" font-size=\"100\"")
		}

		for i := range dimension - 1 {
			canvas.Line((i+1)*Conf.Svg.BaseWidth, 0, (i+1)*Conf.Svg.BaseWidth, Conf.Svg.BaseWidth*dimension, "stroke-width=\"3\" stroke=\"#000\" fill=\"none\"")
		}
		for i := range dimension - 1 {
			canvas.Line(0, (i+1)*Conf.Svg.BaseHeight, Conf.Svg.BaseHeight*dimension, (i+1)*Conf.Svg.BaseHeight, "stroke-width=\"3\" stroke=\"#000\" fill=\"none\"")
		}
		canvas.End()
	}
}
