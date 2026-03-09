package models

import "context"

// QueryContext — контекст запроса, проходящий через pipeline
type QueryContext struct {
	Query           string         // исходный запрос
	NormalizedQuery string         // после нормализации
	Tokens          []string       // токены
	Suggestions     []Suggestion   // подсказки
	Metadata        map[string]any // метаданные
}

// Suggestion — подсказка автодополнения
type Suggestion struct {
	Text      string
	Frequency int64
}

// Processor — интерфейс модуля pipeline
type Processor interface {
	Process(ctx context.Context, qc *QueryContext) (*QueryContext, error)
}
