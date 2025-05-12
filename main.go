package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"is_done"`
}

var DB *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=warrior port=5432 sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := DB.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

// GET

func GetTasks(c echo.Context) error {
	var tasks []Task
	isDoneQuery := c.QueryParam("is_done")

	if isDoneQuery != "" {
		isDone, err := strconv.ParseBool(isDoneQuery)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error":   "invalid_boolean",
				"message": fmt.Sprintf("Неверное значение параметра is_completed: %v", isDone),
			})
		}
		if err := DB.Find(&tasks, "is_done = ?", isDone).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Could not get tasks",
			})
		}

		return c.JSON(http.StatusOK, tasks)
	}
	if err := DB.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not get tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request id")
	}

	var task Task
	if err := DB.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not find task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

// POST
func CreateTask(c echo.Context) error {
	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
	}

	task := Task{
		ID:     uuid.New(),
		Title:  body.Task,
		IsDone: body.IsDone,
	}

	if err := DB.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not create task",
		})
	}
	return c.JSON(http.StatusCreated, task)
}

// PATCH
func UpdateTaskCompletedById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_id",
			"message": fmt.Sprintf("Неверный формат id: %v", err),
		})
	}

	var task Task
	if err := DB.Find(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Could not find task",
		})
	}
	task.IsDone = !task.IsDone
	if err := DB.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not update task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

// PUT

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func UpdateTitleTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_id",
			"message": fmt.Sprintf("Неверный формат id: %v", err),
		})
	}

	var body RequestBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "Invalid request data",
			"message": "Проверьте введенные данные",
		})
	}

	var task Task
	if err := DB.Find(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Could not find task",
		})
	}

	task.Title = body.Task

	if err := DB.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not save task",
		})
	}

	return c.JSON(http.StatusOK, task)
}

// DELETE

func DeleteTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   "invalid_id",
			"message": fmt.Sprintf("Неверный формат id: %v", err),
		})
	}

	if err := DB.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Could not delete task",
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// xh http://localhost:8080/tasks
	e.GET("/tasks", GetTasks)
	// xh http://localhost:8080/tasks/{id}
	e.GET("/tasks/:id", GetTaskById)

	// xh POST http://localhost:8080/tasks task="Новая задача"
	e.POST("/tasks", CreateTask)

	// xh patch http://localhost:8080/tasks/{id}
	e.PATCH("/tasks/:id", UpdateTaskCompletedById)
	// xh put http://localhost:8080/tasks/{id} task="Обновленная задача"
	e.PUT("/tasks/:id", UpdateTitleTaskById)

	// xh delete http://localhost:8080/tasks/{id}
	e.DELETE("/tasks/:id", DeleteTaskById)

	e.Use(middleware.Logger())
	initDB()
	e.Start(":8080")
}
