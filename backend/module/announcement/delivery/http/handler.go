package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"siakup-fs-test-starter/module/announcement/domain"
)

type AnnouncementHandler struct {
	useCase domain.AnnouncementUseCase
}

func NewAnnouncementHandler(useCase domain.AnnouncementUseCase) *AnnouncementHandler {
	return &AnnouncementHandler{useCase: useCase}
}

// RegisterRoutes is already wired for you — you should not need to change this.
func (h *AnnouncementHandler) RegisterRoutes(app *fiber.App) {
	group := app.Group("/announcements")
	group.Get("/", h.GetList)
	group.Post("/", h.Create)
}

type apiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type listMeta struct {
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	Page       int   `json:"page"`
}

func parsePagination(c *fiber.Ctx) domain.Pagination {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = 1
	}
	return domain.Pagination{Page: page, Size: size}
}

// GET /announcements?search=&target_role=&page=&size=
// Fully wired — once you implement UseCase.GetList this endpoint works as-is.
func (h *AnnouncementHandler) GetList(c *fiber.Ctx) error {
	filter := domain.AnnouncementFilter{
		Pagination: parsePagination(c),
		Search:     c.Query("search"),
		TargetRole: c.Query("target_role"),
	}

	items, total, totalPages, err := h.useCase.GetList(c.UserContext(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(apiResponse{Success: false, Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(apiResponse{
		Success: true,
		Data:    items,
		Meta:    listMeta{Total: total, TotalPages: totalPages, Page: filter.Page},
	})
}

// POST /announcements
// Body: {"title": "...", "content": "...", "target_role": "student|lecturer|all"}
// Fully wired — once you implement UseCase.Create this endpoint works as-is.
func (h *AnnouncementHandler) Create(c *fiber.Ctx) error {
	var req domain.CreateAnnouncementRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(apiResponse{Success: false, Message: "invalid request body"}, 400)
	}

	
	created, err := h.useCase.Create(c.UserContext(), req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(apiResponse{Success: false, Message: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(apiResponse{Success: true, Data: created, Message: "announcement created"})
}
