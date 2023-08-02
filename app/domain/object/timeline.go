package object

const (
	DefaultLimit = 40
	MaxLimit     = 80
)

type Timeline struct {
	//only_media bool

	MaxID *int64 `json:"max_id,omitempty"`

	SinceID *int64 `json:"since_id,omitempty"`

	Limit *int64 `json:"limit,omitempty"`
}
