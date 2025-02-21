package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/luispfcanales/confericis-backend/model"
)

func escapeXML(s string) string {
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"'":  "&apos;",
		"\"": "&quot;",
	}

	result := s
	for k, v := range replacements {
		result = strings.ReplaceAll(result, k, v)
	}
	return result
}

func generateSVG(state model.EditorState) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="%f" height="%f" viewBox="0 0 %f %f">
    <defs>
        <style type="text/css">
            @import url('https://fonts.googleapis.com/css2?family=Arial&amp;family=Times+New+Roman&amp;display=swap');
            .text-box {
                box-sizing: border-box;
                white-space: pre-wrap;
                word-wrap: break-word;
                overflow: visible;
            }
        </style>
    </defs>
    <rect width="100%%" height="100%%" fill="white"/>
`, state.Width, state.Height, state.Width, state.Height))

	for _, box := range state.TextBoxes {
		// Escape text content properly for XML
		safeText := escapeXML(box.Text)
		// Replace newlines with <br/> after escaping the text
		safeTextWithBreaks := strings.ReplaceAll(safeText, "\n", "<br/>")

		sb.WriteString(fmt.Sprintf(`    <g transform="translate(%f, %f)">
        <foreignObject width="%f" height="%f">
            <div xmlns="http://www.w3.org/1999/xhtml" 
                 class="text-box"
                 style="
                    width: %fpx;
                    height: %fpx;
                    font-family: %s;
                    font-size: %s;
                    font-weight: %s;
                    font-style: %s;
                    text-decoration: %s;
                    text-align: %s;
                    color: %s;
                    background-color: %s;
                    padding: %s;
                    margin: %s;
                    border: none;
                    display: block;
                    position: absolute;
                    box-sizing: border-box;
                    overflow: visible;
                    line-height: normal;
                    -webkit-font-smoothing: antialiased;
                    -moz-osx-font-smoothing: grayscale;
            ">%s</div>
        </foreignObject>
    </g>
`,
			box.X, box.Y,
			box.Width, box.Height,
			box.Width, box.Height,
			box.Style.FontFamily,
			box.Style.FontSize,
			box.Style.FontWeight,
			box.Style.FontStyle,
			box.Style.TextDecoration,
			box.Style.TextAlign,
			box.Style.Color,
			box.Style.BackgroundColor,
			box.Style.Padding,
			box.Style.Margin,
			safeTextWithBreaks))
	}

	sb.WriteString("</svg>")
	return sb.String()
}

func HandleExportSVG(w http.ResponseWriter, r *http.Request) {
	var state model.EditorState
	if err := json.NewDecoder(r.Body).Decode(&state); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	svg := generateSVG(state)

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Content-Disposition", "attachment; filename=document.svg")
	w.Write([]byte(svg))
}
