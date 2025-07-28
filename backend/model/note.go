package model

type Note struct {
	ID        int    // Unique identifier for the note
	Title     string // Title of the note
	Content   string // Content of the note
	CreatedAt string // Creation timestamp of the note
	OwnerID   int    // ID of the user who owns the note
}

func NewNote(id int, title, content, createdAt string, ownerID int) *Note {
	return &Note{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
		OwnerID:   ownerID,
	}
}

func (n *Note) GetID() int {
	return n.ID
}

func (n *Note) GetTitle() string {
	return n.Title
}

func (n *Note) GetContent() string {
	return n.Content
}

func (n *Note) GetCreatedAt() string {
	return n.CreatedAt
}

func (n *Note) GetOwnerID() int {
	return n.OwnerID
}

func (n *Note) SetTitle(title string) {
	n.Title = title
}

func (n *Note) SetContent(content string) {
	n.Content = content
}
