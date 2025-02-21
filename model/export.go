package model

type Style struct {
	FontFamily      string `json:"fontFamily"`
	FontSize        string `json:"fontSize"`
	FontWeight      string `json:"fontWeight"`
	FontStyle       string `json:"fontStyle"`
	TextDecoration  string `json:"textDecoration"`
	TextAlign       string `json:"textAlign"`
	Color           string `json:"color"`
	BackgroundColor string `json:"backgroundColor"`
	Padding         string `json:"padding"`
	Margin          string `json:"margin"`
}

type TextBox struct {
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Style  Style   `json:"style"`
}

type EditorState struct {
	PageSize  string    `json:"pageSize"`
	Width     float64   `json:"width"`
	Height    float64   `json:"height"`
	TextBoxes []TextBox `json:"textBoxes"`
}
