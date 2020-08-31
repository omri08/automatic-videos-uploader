package models

// Privacy of the video
type Privacy int

// The privacy options : public , unlisted ,  private
const (
	Public Privacy = iota
	Unlisted
	Private
)

func (p Privacy) String() string {
	return [...]string{"public", "unlisted", "private"}[p]
}

// Video contains all the fields we need for uploading a video
type Video struct {
	Path, Title, Description, Category, Keywords, FileName string
	Privacy                                                Privacy
}
