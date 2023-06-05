package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type PDFOption func(*gofpdf.Fpdf)

func FillColor(c color.RGBA) PDFOption {
	return func(pdf *gofpdf.Fpdf) {
		r, g, b := rgb(c)
		pdf.SetFillColor(r, g, b)
	}
}

func rgb(c color.RGBA) (int, int, int) {
	alpha := float64(c.A) / 255.0
	alphaWhite := int(255 * (1.0 - alpha))
	r := int(float64(c.R)*alpha) + alphaWhite
	g := int(float64(c.G)*alpha) + alphaWhite
	b := int(float64(c.B)*alpha) + alphaWhite
	return r, g, b
}

type PDF struct {
	fpdf *gofpdf.Fpdf
	x, y float64
}

func (p *PDF) Move(x, y float64) {
	p.x, p.y = p.x+x, p.y+y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) MoveAbs(x, y float64) {
	p.x, p.y = x, y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) Text(text string) {
	p.fpdf.Text(p.x, p.y, text)
}

func (p *PDF) Polygon(pts []gofpdf.PointType, ops ...PDFOption) {
	for _, opt := range ops {
		opt(p.fpdf)
	}
	p.fpdf.Polygon(pts, "F")
}

func main() {
	// BuildInvoice(false)
	fpdf := gofpdf.New(gofpdf.OrientationLandscape,
		gofpdf.UnitPoint,
		gofpdf.PageSizeLetter, "")

	w, h := fpdf.GetPageSize()
	fpdf.AddPage()

	pdf := PDF{fpdf: fpdf}
	primary := color.RGBA{103, 60, 79, 255}
	secondary := color.RGBA{125, 80, 99, 220}

	// Set the top maroon bar
	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: 0},
		{X: 0, Y: h / 9.0},
		{X: w, Y: 0},
	}, FillColor(secondary))

	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: 0},
		{X: w, Y: 0},
		{X: w, Y: h / 9.0},
	}, FillColor(primary))

	// Set the bottom maroon bar
	pdf.Polygon([]gofpdf.PointType{
		{X: w, Y: h},
		{X: w, Y: h - h/8.0},
		{X: w / 6, Y: h},
	}, FillColor(secondary))

	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: h},
		{X: 0, Y: h - h/8.0},
		{X: w - (w / 6), Y: h},
	}, FillColor(primary))

	// add text
	fpdf.SetFont("times", "B", 50)
	fpdf.SetTextColor(50, 50, 50)
	pdf.MoveAbs(0, 100)
	_, lineHt := fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt,
		"Certificate of Completion",
		gofpdf.AlignCenter)
	pdf.Move(0, lineHt*2.0)

	fpdf.SetFont("arial", "", 20)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*2)

	fpdf.SetFont("times", "", 42)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "Asuzu Kosisochukwu", gofpdf.AlignCenter)
	pdf.Move(0, lineHt*1.5)

	fpdf.SetFont("arial", "", 20)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt,
		"For successfully completing all twenty programming exercices in the \n Gopherisizes programming course for budding Gophers (Go developers)",
		gofpdf.AlignCenter)
	pdf.Move(0, lineHt*3.0)

	fpdf.ImageOptions("images/safari_gopher2.png", w/2-50, pdf.y, 100, 0,
		false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	pdf.Move(0, 65.0)
	fpdf.SetFillColor(100, 100, 100)
	fpdf.Rect(60, pdf.y, 240, 1, "F")
	fpdf.Rect(490, pdf.y, 240, 1, "F")

	fpdf.SetFont("arial", "", 12)
	pdf.Move(0, lineHt/1.5)
	fpdf.SetTextColor(100, 100, 100)
	pdf.MoveAbs(60+100, pdf.y)
	pdf.Text("Date")
	pdf.MoveAbs(490+100, pdf.y)
	pdf.Text("Instructor")
	pdf.MoveAbs(60, pdf.y-lineHt/1.5)
	fpdf.SetFont("times", "", 22)

	pdf.Move(0, -lineHt)
	fpdf.SetTextColor(50, 50, 50)
	_, lineHt = fpdf.GetFontSize()
	date := time.Now()
	dateStr := fmt.Sprintf("%d/%d/%d", date.Day(), date.Month(), date.Year())
	fpdf.CellFormat(240, lineHt, dateStr,
		gofpdf.BorderNone, gofpdf.LineBreakNone,
		gofpdf.AlignCenter, false, 0, "")

	pdf.MoveAbs(490.0, pdf.y)
	svg, err := gofpdf.SVGBasicFileParse("images/sig.svg")
	if err != nil {
		panic(err)
	}
	fpdf.SVGBasicWrite(&svg, 0.12)
	fpdf.CellFormat(240, lineHt-6, "Jonathan Calhoun",
		gofpdf.BorderNone, gofpdf.LineBreakNone,
		gofpdf.AlignCenter, false, 0, "")

	DrawGrid(pdf.fpdf)
	err = pdf.fpdf.OutputFileAndClose("certificate.pdf")
	if err != nil {
		panic(err)
	}
}
