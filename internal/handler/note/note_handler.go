package note

import (
	"github.com/firdausalif/go-fiber/database"
	"github.com/firdausalif/go-fiber/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(ctx *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

func CreateNotes(ctx *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := ctx.BodyParser(note)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.Id = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func GetNote(ctx *fiber.Ctx) error {
	db := database.DB
	var note model.Note
	id := ctx.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.Id == uuid.Nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func UpdateNote(ctx *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := database.DB
	var note model.Note

	id := ctx.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.Id == uuid.Nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updateNoteData updateNote
	err := ctx.BodyParser(&updateNoteData)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.Subtitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text
	db.Save(&note)

	return ctx.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func DeleteNote(ctx *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := ctx.Params("noteId")
	db.Find(&note, "id = ?", id)
	if note.Id == uuid.Nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
