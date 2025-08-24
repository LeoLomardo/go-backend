package user

import (
	"go-backend/internal/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {

	payload := new(LoginPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ERROR": "leitura do JSON falhou"})
	}

	// Busca o usuario no banco
	user := new(User)
	err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", payload.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"ERROR": "credenciais invalidas"})
	}

	// Compara a senha do body com o hash salvo no banco (DIFERENCIAL)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"ERROR": "credenciais invalidas"})
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expira em 72 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("mysecret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ERROR": "could not login"})
	}

	return c.JSON(fiber.Map{"token": t})
}
