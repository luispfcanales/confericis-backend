package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/luispfcanales/confericis-backend/model"
)

func GeneratePDFHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}()

	var input model.EditorState
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Configurar PDF
	pdf := gofpdf.New("P", "pt", input.PageSize, "")
	// pdf := gofpdf.New("P", "pt", "", "")
	//pdf.SetPageSize(input.Width*0.75, input.Height*0.75)

	setupFonts(pdf) // Configuración de fuentes

	pdf.AddPage()

	for _, tb := range input.TextBoxes {
		processTextBox(pdf, tb)
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Escribir respuesta exitosa
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK) // Opcional (200 es el default)
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("Error escribiendo PDF: %v", err)
	}
}

func setupFonts(pdf *gofpdf.Fpdf) {
	pdf.AddUTF8Font("Calibri", "", "./fonts/calibri.ttf")
	pdf.AddUTF8Font("Calibri", "B", "./fonts/calibrib.ttf")
	pdf.AddUTF8Font("Calibri", "I", "./fonts/calibrii.ttf")
}

func processTextBox(pdf *gofpdf.Fpdf, tb model.TextBox) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Error en caja de texto: %v", err)
		}
	}()

	x := tb.X
	y := tb.Y
	width := tb.Width
	height := tb.Height

	// Aplicar padding
	pad := parsePadding(tb.Style.Padding)
	x += pad
	y += pad
	width -= pad * 2
	height -= pad * 2

	// Validaciones críticas
	width = math.Max(1, width)
	height = math.Max(1, height)
	if width <= 0 || height <= 0 {
		return
	}

	// Configurar estilo
	fontSize := parseFontSize(tb.Style.FontSize)
	pdf.SetFont("Calibri", buildFontStyle(tb.Style), fontSize)
	setTextColor(pdf, tb.Style.Color)
	setBackground(pdf, tb.Style.BackgroundColor, x, y, width, height)

	// Procesar texto
	text := strings.TrimSpace(tb.Text)
	if text == "" {
		return
	}

	// Calcular parámetros
	lineHeight := fontSize * 1.25
	if lineHeight <= 0 {
		return
	}

	// Dividir texto
	lines := pdf.SplitText(text, width)
	if len(lines) == 0 {
		return
	}

	// Calcular máximas líneas
	maxLines := int(math.Max(0, math.Floor(height/lineHeight)))
	if maxLines <= 0 {
		return
	}

	// Ajustar líneas visibles
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	// Escribir texto
	pdf.SetXY(x, y)
	currentY := y
	for _, line := range lines {
		if currentY+lineHeight > y+height {
			break
		}
		pdf.CellFormat(width, lineHeight, line, "", 0, mapTextAlign(tb.Style.TextAlign), false, 0, "")
		currentY += lineHeight
		pdf.SetXY(x, currentY)
	}
}

// Funciones auxiliares actualizadas
func buildFontStyle(style model.Style) string {
	var result string
	if style.FontWeight == "bold" {
		result += "B"
	}
	if style.FontStyle == "italic" {
		result += "I"
	}
	return result
}

func parseFontSize(fontSize string) float64 {
	sizeStr := strings.TrimSuffix(fontSize, "px")
	size, _ := strconv.ParseFloat(sizeStr, 64)
	// return size * 0.75 // px -> pt
	return size
}

func setTextColor(pdf *gofpdf.Fpdf, color string) {
	r, g, b := parseColor(color)
	pdf.SetTextColor(r, g, b)
}

func setBackground(pdf *gofpdf.Fpdf, color string, x, y, w, h float64) {
	r, g, b := parseColor(color)
	pdf.SetFillColor(r, g, b)
	pdf.Rect(x, y, w, h, "F")
}

func parseColor(color string) (int, int, int) {
	switch strings.ToLower(color) {
	case "black":
		return 0, 0, 0
	case "white":
		return 255, 255, 255
	case "red":
		return 255, 0, 0
	case "green":
		return 0, 255, 0
	case "blue":
		return 0, 0, 255
	default:
		if strings.HasPrefix(color, "rgba") {
			parts := strings.Split(strings.Trim(color, "rgba() "), ",")
			if len(parts) >= 3 {
				r, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
				g, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
				b, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
				return r, g, b
			}
		}
		return 0, 0, 0 // Negro por defecto
	}
}

func mapTextAlign(align string) string {
	switch align {
	case "left":
		return "L"
	case "right":
		return "R"
	case "center":
		return "C"
	case "justify":
		return "J"
	default:
		return "L"
	}
}

func parsePadding(padding string) float64 {
	if padding == "" {
		return 0
	}

	// Eliminar "px" o "pt" y convertir a float64
	padding = strings.TrimSpace(padding)
	padding = strings.TrimSuffix(padding, "px")
	padding = strings.TrimSuffix(padding, "pt")

	size, err := strconv.ParseFloat(padding, 64)
	if err != nil {
		return 0
	}

	// Si el valor estaba en píxeles, convertirlo a puntos (1px = 0.75pt)
	if strings.Contains(padding, "px") {
		size *= 0.75
	}

	return size
}

// func parsePadding(padding string) float64 {
// 	if padding == "" {
// 		return 0
// 	}
// 	parts := strings.Split(padding, " ")
// 	if len(parts) > 0 {
// 		// val := strings.TrimSuffix(parts[0], "px")
// 		val := strings.TrimSuffix(parts[0], "pt")
// 		size, err := strconv.ParseFloat(val, 64)
// 		if err != nil {
// 			return 0
// 		}
// 		// return size * 0.75
// 		return size
// 	}
// 	return 0
// }
