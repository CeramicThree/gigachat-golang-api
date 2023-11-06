package domain

type Message struct {
	Role    string `json:"role"`    // Контекст
	Content string `json:"content"` // Запрос пользователя
}
