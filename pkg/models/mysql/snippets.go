package mysql

import (
	"database/sql"
	"github.com/rehydrate1/sniply/pkg/models"
)

// тип-обёртка пула подключения sql.DB
type SnippetModel struct {
	DB *sql.DB
}

// метод для создания заметки в бд
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
