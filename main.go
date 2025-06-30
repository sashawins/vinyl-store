package main

import (
	// "context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
	"vinyl-store/internal/db"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	entuser "vinyl-store/ent/user"
	entadmin "vinyl-store/ent/admin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	client := db.Connect()
	defer client.Close()

	app := fiber.New()
	api := app.Group("/api/v1")

	// ─── Vinyls ──────────────────────────────────────────────

	api.Get("/vinyls", func(c *fiber.Ctx) error {
		items, err := client.Vinyl.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(items)
	})

	api.Get("/vinyls/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("invalid UUID")
		}
		item, err := client.Vinyl.Get(c.Context(), id)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.JSON(item)
	})

	api.Post("/vinyls", func(c *fiber.Ctx) error {
		var body struct {
			Title      string  `json:"title"`
			Price      float64 `json:"price"`
			StockCount int     `json:"stock_count"`
			ArtistID   string  `json:"artist_id"`
			GenreID    string  `json:"genre_id"`
			CoverURL   string  `json:"cover_url"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		artistID, err := uuid.Parse(body.ArtistID)
		if err != nil {
			return c.Status(400).SendString("invalid artist_id")
		}
		genreID, err := uuid.Parse(body.GenreID)
		if err != nil {
			return c.Status(400).SendString("invalid genre_id")
		}

		item, err := client.Vinyl.Create().
			SetTitle(body.Title).
			SetPrice(body.Price).
			SetStockCount(int32(body.StockCount)).
			SetArtistID(artistID).
			SetGenreID(genreID).
			SetCoverURL(body.CoverURL).
			Save(c.Context())

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(201).JSON(item)
	})

	api.Delete("/vinyls/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("invalid UUID")
		}
		err = client.Vinyl.DeleteOneID(id).Exec(c.Context())
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}
		return c.SendStatus(204)
	})

	// ─── Artists ─────────────────────────────────────────────

	api.Get("/artists", func(c *fiber.Ctx) error {
		artists, err := client.Artist.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(artists)
	})

	api.Post("/artists", func(c *fiber.Ctx) error {
		var body struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		if body.Name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Artist name is required")
		}

		artist, err := client.Artist.Create().
			SetName(body.Name).
			SetCountry(body.Country).
			Save(c.Context())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(artist)
	})

	// ─── Genres ──────────────────────────────────────────────

	api.Get("/genres", func(c *fiber.Ctx) error {
		genres, err := client.Genre.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(genres)
	})

	api.Post("/genres", func(c *fiber.Ctx) error {
		var body struct {
			Name string `json:"name"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		if body.Name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Genre name is required")
		}

		genre, err := client.Genre.Create().
			SetName(body.Name).
			Save(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(genre)
	})

	// ─── News ────────────────────────────────────────────────

	api.Get("/news", func(c *fiber.Ctx) error {
		items, err := client.News.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(items)
	})

	api.Post("/news", func(c *fiber.Ctx) error {
		var body struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		item, err := client.News.Create().
			SetTitle(body.Title).
			SetContent(body.Content).
			Save(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(201).JSON(item)
	})

	// ─── Orders ──────────────────────────────────────────────

	api.Get("/orders", func(c *fiber.Ctx) error {
		orders, err := client.Order.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(orders)
	})

	// Получение по ID
	api.Get("/orders/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid order ID format")
		}

		order, err := client.Order.Get(c.Context(), id)
		if err != nil {
			return c.Status(404).SendString("Order not found")
		}

		return c.JSON(order)
	})


	api.Post("/orders", func(c *fiber.Ctx) error {
		var body struct {
			UserID string `json:"user_id"`
			Items  []struct {
				VinylID  string `json:"vinyl_id"`
				Quantity int    `json:"quantity"`
			} `json:"items"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		if body.UserID == "" || len(body.Items) == 0 {
			return c.Status(400).SendString("User ID and at least one item are required")
		}

		userID, err := uuid.Parse(body.UserID)
		if err != nil {
			return c.Status(400).SendString("Invalid user ID format")
		}

		total := 0.0
		ctx := c.Context()

		// Подготовим order_items
		var orderItems []struct {
			VinylID uuid.UUID
			Qty     int
			Price   float64
		}
		for _, item := range body.Items {
			vinylID, err := uuid.Parse(item.VinylID)
			if err != nil {
				return c.Status(400).SendString("Invalid vinyl ID format")
			}
			vinyl, err := client.Vinyl.Get(ctx, vinylID)
			if err != nil {
				return c.Status(404).SendString("Vinyl not found: " + vinylID.String())
			}
			orderItems = append(orderItems, struct {
				VinylID uuid.UUID
				Qty     int
				Price   float64
			}{
				VinylID: vinylID,
				Qty:     item.Quantity,
				Price:   vinyl.Price,
			})
			total += vinyl.Price * float64(item.Quantity)
		}

		// Создание заказа
		order, err := client.Order.Create().
			SetUserID(userID).
			SetStatus("pending").
			SetTotalAmount(total).
			Save(ctx)
		if err != nil {
			return c.Status(500).SendString("Failed to create order: " + err.Error())
		}

		// Создание позиций заказа
		for _, item := range orderItems {
			err := client.OrderItem.Create().
				SetOrderID(order.ID).
				SetVinylID(item.VinylID).
				SetQuantity(item.Qty).
				SetUnitPrice(item.Price).
				Exec(ctx)
			if err != nil {
				return c.Status(500).SendString("Failed to create order item: " + err.Error())
			}
		}

		return c.Status(201).JSON(fiber.Map{
			"order_id": order.ID,
			"total":    total,
		})
	})


	// ─── Admin ───────────────────────────────────────────────

	api.Post("/admin/login", func(c *fiber.Ctx) error {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		admin, err := client.Admin.Query().
			Where(entadmin.EmailEQ(body.Email)).
			Only(c.Context())
		if err != nil {
			return c.Status(401).SendString("Invalid email or password")
		}

		err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(body.Password))
		if err != nil {
			return c.Status(401).SendString("Invalid email or password")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"admin_id": admin.ID.String(),
			"role":     "admin",
			"exp":      time.Now().Add(72 * time.Hour).Unix(),
		})

		secret := os.Getenv("JWT_SECRET")
		tokenStr, err := token.SignedString([]byte(secret))
		if err != nil {
			return c.Status(500).SendString("Failed to generate token")
		}

		return c.JSON(fiber.Map{
			"token": tokenStr,
		})
	})

	api.Post("/admin/register", func(c *fiber.Ctx) error {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}
		if body.Email == "" || body.Password == "" {
			return c.Status(400).SendString("Email and password are required")
		}
	
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).SendString("Failed to hash password")
		}
	
		admin, err := client.Admin.Create().
			SetEmail(body.Email).
			SetPasswordHash(string(hash)).
			SetName(body.Name).
			Save(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	
		return c.Status(201).JSON(fiber.Map{
			"message": "Admin registered successfully",
			"admin":   admin,
		})
	})


	// ─── Users ───────────────────────────────────────────────

	api.Get("/users", func(c *fiber.Ctx) error {
		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(users)
	})

	api.Post("/users/register", func(c *fiber.Ctx) error {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}
		if body.Email == "" || body.Password == "" {
			return c.Status(400).SendString("Email and password are required")
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).SendString("Failed to hash password")
		}

		user, err := client.User.Create().
			SetEmail(body.Email).
			SetPasswordHash(string(hash)).
			SetName(body.Name).
			Save(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(201).JSON(user)
	})

	api.Post("/users/login", func(c *fiber.Ctx) error {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		user, err := client.User.Query().
			Where(entuser.EmailEQ(body.Email)).
			Only(c.Context())
		if err != nil {
			return c.Status(401).SendString("Invalid email or password")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
		if err != nil {
			return c.Status(401).SendString("Invalid email or password")
		}

		// Генерация JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID.String(),
			"exp":     time.Now().Add(72 * time.Hour).Unix(),
		})

		secret := os.Getenv("JWT_SECRET")
		tokenStr, err := token.SignedString([]byte(secret))
		if err != nil {
			return c.Status(500).SendString("Failed to generate token")
		}

		return c.JSON(fiber.Map{
			"token": tokenStr,
		})
	})

	log.Fatal(app.Listen(":8080"))
}
