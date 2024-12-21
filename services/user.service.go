package services

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/poonyawat0511/go-fiber/models"
)

// database crud

func createUser(db *sql.DB, first_name string, last_name string) (int, error) {
	var id int
	err := db.QueryRow(`INSERT INTO users(first_name, last_name) VALUES($1, $2) RETURNING id;`, first_name, last_name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.First_name, &u.Last_name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func getUser(db *sql.DB, id int) (models.User, error) {
	var u models.User
	row := db.QueryRow(`SELECT id, first_name, last_name FROM users WHERE id = $1;`, id)
	err := row.Scan(&u.ID, &u.First_name, &u.Last_name)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func updateUser(db *sql.DB, id int, first_name string, last_name string) error {
	_, err := db.Exec(`UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;`, first_name, last_name, id)
	return err
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM users WHERE id = $1;`, id)
	return err
}

// fiber crud

func CreateUserHandle(c *fiber.Ctx, db *sql.DB) error {
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	id, err := createUser(db, u.First_name, u.Last_name)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	u.ID = id

	return c.JSON(u)
}

func GetUsersHandle(c *fiber.Ctx, db *sql.DB) error {
	users, err := getUsers(db)
	if err != nil {
		fmt.Println("Error querying users:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}

func GetUserHandle(c *fiber.Ctx, db *sql.DB) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	user, err := getUser(db, userId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(user)
}

func UpdateUserHandle(c *fiber.Ctx, db *sql.DB) error {
	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	err = updateUser(db, userId, u.First_name, u.Last_name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Query ผู้ใช้ที่อัปเดตแล้วกลับมา
	user, err := getUser(db, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func DeleteUserHandle(c *fiber.Ctx, db *sql.DB) error {

	userId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteUser(db, userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":      userId,
		"message": fmt.Sprintf("User with ID %d has been deleted", userId),
	})
}
