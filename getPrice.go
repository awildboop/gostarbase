package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hbollon/go-edlib"
)

func HandleGetPrice(client *http.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		oreName, err := edlib.FuzzySearchThreshold(c.Params("ore", ""), oreNames[:], .85, edlib.JaroWinkler)
		if err != nil || oreName == "" {
			oreName2, _ := edlib.FuzzySearchThreshold(c.Params("ore"), oreNames[:], .10, edlib.JaroWinkler)
			return c.Render("error", fiber.Map{
				"title":       "BOOP's Price Locator",
				"description": "An ore with the given name was not found, closest match: " + oreName2,
			})
		}

		ore, err := getOre(oreName)
		if err != nil {
			oreName2, _ := edlib.FuzzySearchThreshold(c.Params("ore"), oreNames[:], .10, edlib.JaroWinkler)
			return c.Render("error", fiber.Map{
				"title":       "BOOP's Price Locator",
				"description": "An ore with the given name was not found, closest match: " + oreName2,
			})
		}

		res, err := client.Get(fmt.Sprintf("https://api.sadtech.io/api/item/history/%d", ore.ID))
		if err != nil {
			c.Response().Header.SetStatusCode(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{"message": err.Error()})
		}
		defer res.Body.Close()

		var data OreData
		json.NewDecoder(res.Body).Decode(&data)

		return c.Render("price", fiber.Map{
			"title":       "Nhurgite Ore Price | Provided by Sadtech.io",
			"description": fmt.Sprintf("%d Cr", data[len(data)-1].Unitprice),
			"url":         fmt.Sprintf("https://sadtech.io/item/%d", ore.ID),
		})
	}
}
