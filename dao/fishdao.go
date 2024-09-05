package dao

import (
	"database/sql"
	"fmt"
	"web/global"
)

type Fish struct {
	Name           string `json:"name"`
	Price_official string `json:"price_official"`
	Price_original string `json:"price_original"`
	Minimum        string `json:"minimum"`
}

type FishImage struct {
	Name       string `json:"name"`
	Link       string `json:"link"`
	Type_Id    string `json:"type_id"`
	Picture_Id string `json:"Picture_id"`
}

func GetFishPrice() []Fish {
	rows, err := global.Mysql.Query("SELECT name, price_official, price_original, minimum from fish_price")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	if err != nil {
		fmt.Printf("Query failed to get fish_prices: %v\n", err)
		return nil
	} else {
		fmt.Println("query fish_price successfully")
	}
	var fishes = make([]Fish, 0)
	for rows.Next() {
		var fish Fish
		err := rows.Scan(&fish.Name, &fish.Price_official, &fish.Price_original, &fish.Minimum)
		if err != nil {
			return nil
		}
		fishes = append(fishes, fish)
	}
	return fishes
}

func GetFishImage() []FishImage {
	rows, err := global.Mysql.Query("SELECT name, link, type_id, picture_id from fish_images")

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	if err != nil {
		fmt.Printf("Query failed to get fish images: %v\n", err)
		return nil
	} else {
		fmt.Println("query data for fish image successfully")
	}
	var fishes = make([]FishImage, 0)
	for rows.Next() {
		var fish FishImage
		err := rows.Scan(&fish.Name, &fish.Type_Id, &fish.Picture_Id, &fish.Link)
		if err != nil {
			return nil
		}
		fishes = append(fishes, fish)
	}
	return fishes
}
