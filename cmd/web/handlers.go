package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/rehydrate1/sniply/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: url.Values{},
		Errors: make(map[string]string),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errorsMap := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errorsMap["title"] = "Это поле не может быть пустым"
	} else if utf8.RuneCountInString(title) > 100 {
		errorsMap["title"] = "Это поле слишком длинное (максимум 100 символов)"
	}

	if strings.TrimSpace(content) == "" {
		errorsMap["content"] = "Это поле не может быть пустым"
	}

	allowedExpires := map[string]bool{"1": true, "7": true, "365": true}
	if !allowedExpires[expires] {
		errorsMap["expires"] = "Недопустимое значение для срока хранения"
	}

	if len(errorsMap) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form: r.PostForm,
			Errors: errorsMap,
		})
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
