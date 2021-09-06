package main

import (
	"fmt"
	"strings"
)

var oreNames = [16]string{"aegisium", "ajatite", "arkanium", "bastium", "charodium", "corazium", "exorium", "ice", "karnite", "kutonium", "lukium", "nhurgite", "surtrite", "valkite", "vokarium", "ymrium"}

var ores = [16]Ore{
	{
		Name: "Aegisium",
		ID:   1,
	},
	{
		Name: "Ajatite",
		ID:   2,
	},
	{
		Name: "Arkanium",
		ID:   3,
	},
	{
		Name: "Bastium",
		ID:   4,
	},
	{
		Name: "Charodium",
		ID:   5,
	},
	{
		Name: "Corazium",
		ID:   6,
	},
	{
		Name: "Exorium",
		ID:   7,
	},
	{
		Name: "Ice",
		ID:   9,
	},
	{
		Name: "Karnite",
		ID:   1,
	},
	{
		Name: "Kutonium",
		ID:   12,
	},
	{
		Name: "Lukium",
		ID:   13,
	},
	{
		Name: "Nhurgite",
		ID:   15,
	},
	{
		Name: "Surtrite",
		ID:   17,
	},
	{
		Name: "Valkite",
		ID:   20,
	},
	{
		Name: "Vokarium",
		ID:   21,
	},
	{
		Name: "Ymrium",
		ID:   23,
	},
}

// Takes in the ore name and gives the ore, returns empty ore struct if none matching
func getOre(name string) (Ore, error) {
	for _, ore := range ores {
		if strings.EqualFold(name, ore.Name) {
			return ore, nil
		}
	}
	return Ore{}, fmt.Errorf("no matching ore with given name")
}
