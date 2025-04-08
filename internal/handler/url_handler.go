package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harsh6373/go-url-shortner/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(s *service.URLService) *URLHandler {
	return &URLHandler{s}
}

func (h *URLHandler) Shorten(c *fiber.Ctx) error {
	type request struct {
		URL        string `json:"url"`
		CustomSlug string `json:"custom_slug"`
		ExpireAt   string `json:"expire_at"` // ISO format
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	var expires *time.Time
	if body.ExpireAt != "" {
		t, err := time.Parse(time.RFC3339, body.ExpireAt)
		if err == nil {
			expires = &t
		}
	}

	url, err := h.service.Shorten(body.URL, body.CustomSlug, expires)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"slug": url.Slug,
		"url":  url.Original,
	})
}

func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	slug := c.Params("slug")
	userAgent := c.Get("User-Agent")

	original, err := h.service.Resolve(slug, userAgent)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Redirect(original, fiber.StatusMovedPermanently)
}

func (h *URLHandler) GetAnalytics(c *fiber.Ctx) error {
	slug := c.Params("slug")

	clicks, err := h.service.GetClickAnalytics(slug)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(clicks)
}
