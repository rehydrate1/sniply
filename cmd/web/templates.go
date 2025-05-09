package main

import "github.com/rehydrate1/sniply/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
