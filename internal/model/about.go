package model

type AboutItem struct {
	CommentEnabled string
	Content        string
	MusicId        string
	Title          string
}

type GetAboutOutput struct {
	AboutItem
}

type UpdateAboutInput struct {
	AboutItem
}
