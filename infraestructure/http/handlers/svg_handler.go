package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/confericis-backend/model"
)

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
			strings.ReplaceAll(box.Text, "\n", "<br/>")))
	}

	sb.WriteString("</svg>")
	return sb.String()
}

func HandleExportSVG(w http.ResponseWriter, r *http.Request) {
	// Configurar CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

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
