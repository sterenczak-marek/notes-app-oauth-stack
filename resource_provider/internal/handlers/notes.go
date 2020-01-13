package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"

	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/internal"
	"github.com/sterenczak-marek/notes-app-oauth-stack/resource_provider/internal/models"
)

func NoteListHandler(rw http.ResponseWriter, req *http.Request) {
	notes := models.GetNotes(getUserEmail(req))

	internal.ResponseJSON(rw, notes, http.StatusOK)
}

func NoteCreateHandler(rw http.ResponseWriter, req *http.Request) {
	note := &models.Note{
		UserEmail: getUserEmail(req),
	}
	err := json.NewDecoder(req.Body).Decode(note)
	if err != nil {
		internal.ResponseError(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err = note.Create(); err != nil {
		log.Panic(err)
	}

	internal.ResponseJSON(rw, note, http.StatusOK)
}

func NoteDetailHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uuid := vars["noteUUID"]

	note := models.GetNote(uuid, getUserEmail(req))
	if note == nil {
		internal.ResponseError(rw, fmt.Sprintf("Note '%s' not found", uuid), http.StatusNotFound)
		return
	}

	internal.ResponseJSON(rw, note, http.StatusOK)
}

func NoteUpdateHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uuid := vars["noteUUID"]

	note := models.GetNote(uuid, getUserEmail(req))
	if note == nil {
		internal.ResponseError(rw, fmt.Sprintf("Note '%s' not found", uuid), http.StatusNotFound)
		return
	}

	changedNote := &models.Note{}
	copier.Copy(&changedNote, &note)
	defer req.Body.Close()
	err := json.NewDecoder(req.Body).Decode(&changedNote)
	if err != nil {
		internal.ResponseError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err = note.Update(changedNote); err != nil {
		log.Panic(err)
	}
	internal.ResponseJSON(rw, note, http.StatusOK)
}

func NoteDeleteHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uuid := vars["noteUUID"]

	note := models.GetNote(uuid, getUserEmail(req))
	if note == nil {
		internal.ResponseError(rw, fmt.Sprintf("Note '%s' not found", uuid), http.StatusNotFound)
		return
	}

	if err := note.Delete(); err != nil {
		log.Panic(err)
	}
	internal.ResponseJSON(rw, note, http.StatusNoContent)
}

func getUserEmail(req *http.Request) string {
	return req.Header.Get("X-User")
}
