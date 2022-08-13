package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	. "github.com/Peyton232/HTTPServerGenerator/database"
	"github.com/Peyton232/HTTPServerGenerator/models"
)

type Handler struct {
	// add any needed fields here
	DB *DB
}

func NewHandler() (*Handler, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	return &Handler{
		DB: db,
	}, nil
}

// Here, we implement all of the handlers in the ServerInterface
func (p *Handler) FindPets(ctx echo.Context, params models.FindPetsParams) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) AddPet(ctx echo.Context) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) UpdatePet(ctx echo.Context) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) FindPetByID(ctx echo.Context, petId int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}

func (p *Handler) DeletePet(ctx echo.Context, id int64) error {
	// BUSINESS LOGIC
	return ctx.NoContent(http.StatusNotImplemented)
}
