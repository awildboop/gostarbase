package main

import "time"

type Ore struct {
	Name string
	ID   int
}

type OreData []struct {
	ID        string        `json:"id"`
	Itemid    int           `json:"itemId"`
	Date      time.Time     `json:"date"`
	Scanid    string        `json:"scanID"`
	Unitprice int           `json:"unitPrice"`
	Available int           `json:"available"`
	Listings  []interface{} `json:"listings"`
}
