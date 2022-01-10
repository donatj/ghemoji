package ghemoji

type EmojiResponse []EmojiData

type EmojiData struct {
	Emoji          string   `json:"emoji,omitempty"`
	Description    string   `json:"description,omitempty"`
	Category       string   `json:"category,omitempty"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version,omitempty"`
	IosVersion     string   `json:"ios_version,omitempty"`
}
