package controllers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"server/src/configs"
	"server/src/models"
	"server/src/responses"
	"time"
)

var categoriesCollection *mongo.Collection = configs.GetCollection(configs.DB, "Categories")

func GetAllCategories(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var categories []models.Categories
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"index", 1}})

	results, err := categoriesCollection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.CategoriesResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleCategory models.Categories
		if err = results.Decode(&singleCategory); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.CategoriesResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		categories = append(categories, singleCategory)
	}

	return c.Status(http.StatusOK).JSON(
		responses.CategoriesResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": categories}},
	)
}
