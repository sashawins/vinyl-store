package main

import (
	// "context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"log"
	"vinyl-store/internal/db"
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

	// ─── Genres ──────────────────────────────────────────────

	api.Get("/genres", func(c *fiber.Ctx) error {
		genres, err := client.Genre.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(genres)
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

	// ─── Users (временно, для отладки) ──────────────────────

	api.Get("/users", func(c *fiber.Ctx) error {
		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(users)
	})

	log.Fatal(app.Listen(":8080"))
}
