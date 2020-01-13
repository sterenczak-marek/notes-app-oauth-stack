package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/config"
)

const noteTableName = "notes"

type Note struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UUID      uuid.UUID `gorm:"unique_index;not null" json:"uuid"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Text      string    `json:"text"`
	UserEmail string    `json:"-"`
}

func init() {
	db := config.DB
	db.Debug().AutoMigrate(&Note{})
}

func (note *Note) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

func (note *Note) Create() error {
	config.DB.Create(note)
	if note.ID <= 0 {
		return errors.New("Failed to create note, connection error.")
	}
	return nil
}

func (note *Note) Update(changedNote *Note) error {
	return config.DB.Model(&note).
		// allow update only on those fields
		Select("title", "text").
		Updates(
			changedNote,
		).
		Error
}
func (note *Note) Delete() error {
	return config.DB.
		Delete(&note).
		Error
}

func GetNote(uuid string, userEmail string) *Note {
	note := Note{}
	config.DB.Table(noteTableName).First(&note, "uuid = ? AND user_email = ?", uuid, userEmail)
	if note.ID == 0 {
		return nil
	}
	return &note
}

func GetNotes(userEmail string) []*Note {
	var notes []*Note
	config.DB.Table(noteTableName).Find(&notes, "user_email = ?", userEmail)
	return notes
}
