package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello WORLD")

	app := fiber.New()

	taskList := []Task{}

	app.Get("/api/tasks", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "success", "data": taskList})
	})

	app.Post("/api/tasks", func(c *fiber.Ctx) error {
		task := &Task{}

		if err := c.BodyParser(task); err != nil {
			return c.Status(422).JSON(fiber.Map{"message": "Invalid request body", "error": err.Error()})
		}

		if task.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		task.ID = len(taskList) + 1
		taskList = append(taskList, *task)

		return c.Status(201).JSON(fiber.Map{"message": "success", "data": task})
	})

	app.Patch("/api/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, task := range taskList {
			if fmt.Sprint(task.ID) == id {
				taskList[i].Completed = !taskList[i].Completed

				return c.Status(200).JSON(fiber.Map{"message": "success", "data": taskList[i]})
			}
		}

		return c.Status(404).JSON(fiber.Map{"message": "task not found"})
	})

	app.Delete("/api/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		taskID, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
		}

		for i, task := range taskList {
			if task.ID == taskID {
				taskList = append(taskList[:i], taskList[i+1:]...)

				return c.Status(200).JSON(fiber.Map{"message": "success"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"message": "task not found"})
	})

	log.Fatal(app.Listen(":5000"))
}
