package Controllers

import (
	"context"
	"encoding/json"

	"github.com/MinhajSadik/Go-MongoDB/DB"
	"github.com/MinhajSadik/Go-MongoDB/Models"
	"github.com/gofiber/fiber"
)

func Register(c *fiber.Ctx) error {
	collection := DB.SystemCollections.AuthPracticeDB
	var self Models.AuthRegister
	c.BodyParser(&self)

	if err := self.Validate(); err != nil {
		return err
	}

	res, err := collection.InsertOne(context.Background(), self)

	if err != nil {
		return nil
	}

	response, _ := json.Marshal(res)
	c.Status(200).Send(response)

	return nil

}
