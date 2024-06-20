package httpcontext

import (
	"cfasuite/internal/database"
	"context"
	"html/template"

	"time"
)

type Context struct {
	context.Context
	Templates *template.Template
	StartTime time.Time
	DB        database.DB
}
