package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GetNextLine(lineHeight float64, startingPoint float64) func() float64 {
	var nextPoint float64 = startingPoint

	return func() float64 {
		nextPoint += lineHeight
		return float64(nextPoint)
	}
}

func Experiment() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait,
		gofpdf.UnitPoint,
		gofpdf.PageSizeLetter, "")

	width, height := pdf.GetPageSize()

	fmt.Println("Width", width, "Height", height)

	pdf.AddPage()
	pdf.SetFont("arial", "B", 38)
	pdf.SetTextColor(255, 0, 0)
	_, lineHeight := pdf.GetFontSize()
	nextLine := GetNextLine(lineHeight, 0)
	pdf.Text(0, nextLine(), "Hello, world")
	pdf.SetFont("times", "B", 16)
	pdf.SetTextColor(0, 255, 0)
	prevLineHeight := lineHeight
	_, lineHeight = pdf.GetFontSize()
	nextLine = GetNextLine(lineHeight, nextLine()-prevLineHeight)
	pdf.Text(0, nextLine(), "Sample world")

	pdf.SetTextColor(100, 100, 100)
	pdf.MultiCell(0, lineHeight*1.5,
		"Blafe wge g woewo g jwof ew vwogrj w \n howewigowhrw  woewoihwogir g ohw ow j\n oyewwoygwogiwryghorwi",
		gofpdf.BorderNone, gofpdf.AlignRight, false)

	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")

	pdf.Polygon([]gofpdf.PointType{
		{X: 110, Y: 250},
		{X: 160, Y: 300},
		{X: 110, Y: 350},
		{X: 60, Y: 300},
	}, "FD")

	pdf.ImageOptions("images/header.jpg",
		244, 277, 300, 100, false, gofpdf.ImageOptions{ReadDpi: true, AllowNegativePosition: true}, 0, "")
	DrawGrid(pdf)

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}

func DrawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)

	for x := 0.0; x < w; x += (w / 20.0) {
		pdf.Line(x, 0, x, h)
		_, lineHeight := pdf.GetFontSize()
		pdf.Text(x, lineHeight, fmt.Sprintf("%d", int(x)))
	}

	for x := 0.0; x < h; x += (h / 20.0) {
		pdf.Line(0, x, w, x)
		pdf.Text(0, x, fmt.Sprintf("%d", int(x)))
	}
}
