package domain

type Choices struct {
	Finish_reason string
	Index         int64
	Message       *Message
}

type Usage struct {
	Completion_tokens int64
	Prompt_tokens     int64
	Total_tokens      int64
}

type GigaResponse struct {
	Choices []*Choices
	Created int64
	Model   string
	Object  string
	Usage	*Usage
}
