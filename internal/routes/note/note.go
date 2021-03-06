package note

import (
	noteHandler "github.com/firdausalif/go-fiber/internal/handler/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", noteHandler.CreateNotes)
	note.Get("/", noteHandler.GetNotes)
	note.Get("/:noteId", noteHandler.GetNote)
	note.Put("/:noteId", noteHandler.UpdateNote)
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
