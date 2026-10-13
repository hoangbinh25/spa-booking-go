package branch

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Handler translates branch use cases into HTTP responses.
type Handler struct {
	lister Lister
	finder Finder
}

type ListResponse struct {
	Data []Branch `json:"data"`
}

type DetailResponse struct {
	Data Branch `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHandler(lister Lister, finder Finder) *Handler {
	return &Handler{
		lister: lister,
		finder: finder,
	}
}

// List godoc
// @Summary List active branches
// @Description Returns all active spa branches.
// @Tags branches
// @Produce json
// @Success 200 {object} ListResponse
// @Failure 500 {object} ErrorResponse
// @Router /branches [get]
func (h *Handler) List(c *gin.Context) {
	branches, err := h.lister.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "could not load branches",
		})
		return
	}

	c.JSON(http.StatusOK, ListResponse{Data: branches})
}

// GetByID godoc
// @Summary Get an active branch
// @Description Returns an active spa branch by its ID.
// @Tags branches
// @Produce json
// @Param id path string true "Branch UUID" format(uuid)
// @Success 200 {object} DetailResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /branches/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid branch id"})
		return
	}

	branch, err := h.finder.FindByID(c.Request.Context(), id.String())
	if errors.Is(err, ErrNotFound) {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "branch not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not load branch"})
		return
	}

	c.JSON(http.StatusOK, DetailResponse{Data: branch})
}
