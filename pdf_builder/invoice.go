package main

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type InvoiceSummary struct {
	ClientName             string
	ClientAddress          string
	ClientCityStateCountry string
	ClientPostalCode       string
	InvoiceNumber          string
	InvoiceIssueDate       string
	InvoiceTotal           string
}

type InvoiceItem struct {
	Description  string
	PricePerUnit string
	Quantity     string
	Amount       string
}

func BuildInvoice(withGrid bool) {
	pdf := gofpdf.New(gofpdf.OrientationPortrait,
		gofpdf.UnitPoint,
		gofpdf.PageSizeLetter, "")

	const (
		bannerHt = 118.0
		xIndent  = 20.0
	)
	pdf.AddPage()
	pdf.SetFont("arial", "B", 38)
	pdf.SetTextColor(255, 0, 0)

	// Draw top banner
	DrawTopBanner(pdf, bannerHt, xIndent)
	// Draw bottom stip
	DrawBottomStrip(pdf, bannerHt)

	// Draw Invoice summary
	invoiceSummary := InvoiceSummary{
		"Client Name", "123 Client Address",
		"City, State, Country", "Postal Code",
		"000000123", "05/29/2023", "$1838.53"}
	DrawInvoiceSummary(pdf, &invoiceSummary, xIndent)

	// Draw Table content
	items := []InvoiceItem{
		{"2x6 Lumber -8'", "$3.75", "220", "$825.00"},
		{"2x6 Lumber -10'", "$5.55", "18", "$99.90"},
		{"2x6 Lumber -8'", "$2.99", "50", "$411.00"},
		{"DryWall Street", "$8.22", "3", "$43.65"},
		{"Paint", "14.55", "22", "$219.78"},
		{"2x6 Lumber -10'", "$5.55", "18", "$99.90"},
		{"2x6 Lumber -8'", "$2.99", "50", "$411.00"},
		{"DryWall Street", "$8.22", "3", "$43.65"},
		{"Paint", "14.55", "22", "$219.78"},
		{"Paint", "14.55", "22", "$219.78"},
		{"2x6 Lumber -10'", "$5.55", "18", "$99.90"},
		{"2x6 Lumber -8'", "$3.75", "220", "$825.00"},
		{"2x6 Lumber -10'", "$5.55", "18", "$99.90"},
		{"2x6 Lumber -8'", "$2.99", "50", "$411.00"},
		{"DryWall Street", "$8.22", "3", "$43.65"},
		{"Paint", "14.55", "22", "$219.78"},
		{"2x6 Lumber -10'", "$5.55", "18", "$99.90"},
		{"2x6 Lumber -8'", "$2.99", "50", "$411.00"},
		{"DryWall Street", "$8.22", "3", "$43.65"},
		{"Paint", "14.55", "22", "$219.78"},
		{"Paint", "14.55", "22", "$219.78"},
		{"2x6 Lumber -10'", "$5.55", "18", "$99.90"},
	}
	DrawInvoiceTable(pdf, items, &invoiceSummary, "$1838.53", xIndent, 260, bannerHt, withGrid)

	// Draw grid on image
	if withGrid {
		DrawGrid(pdf)
	}
	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		panic(err)
	}
}

func DrawTopBanner(pdf *gofpdf.Fpdf, bannerHt float64, xIndent float64) {
	w, _ := pdf.GetPageSize()
	pdf.SetFillColor(103, 60, 79)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: 0},
		{X: w, Y: 0},
		{X: w, Y: bannerHt},
		{X: 0, Y: bannerHt * 0.8},
	}, "F")

	// Write invoice text
	pdf.SetFont("Arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHeight := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+(lineHeight/3.1), "INVOICE")

	// Gopher image
	pdf.ImageOptions("images/safari_gopher2.png",
		220, lineHeight*0.3, 80, 80,
		false, gofpdf.ImageOptions{ReadDpi: true, AllowNegativePosition: true}, 0, "")

	// company information
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	pdf.Text(330.0, lineHeight*1.3, "(123) 456-7890")
	pdf.Text(330.0, lineHeight*1.7, "keloasuzu@yahoo.com")
	pdf.Text(330.0, lineHeight*2.1, "kosiasuzu.com")

	// address
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	pdf.Text(480.0, lineHeight*1.3, "Flat 4")
	pdf.Text(480.0, lineHeight*1.7, "25 Elvatham Road")
	pdf.Text(480.0, lineHeight*2.1, "Birmingham, UK")

}

func DrawBottomStrip(pdf *gofpdf.Fpdf, bannerHt float64) {
	w, h := pdf.GetPageSize()
	pdf.SetFillColor(103, 60, 79)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Polygon([]gofpdf.PointType{
		{X: 0, Y: h},
		{X: 0, Y: h - (bannerHt * 0.2)},
		{X: w, Y: h - (bannerHt * 0.1)},
		{X: w, Y: h},
	}, "F")
}

func DrawInvoiceSummary(pdf *gofpdf.Fpdf, invoiceSummary *InvoiceSummary, xIndent float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	pdf.Text(xIndent, 130, "Billed To")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(xIndent, 160, invoiceSummary.ClientName)
	pdf.Text(xIndent, 180, invoiceSummary.ClientAddress)
	pdf.Text(xIndent, 200, invoiceSummary.ClientCityStateCountry)
	pdf.Text(xIndent, 220, invoiceSummary.ClientPostalCode)

	pdf.SetTextColor(180, 180, 180)
	pdf.Text(xIndent+200, 130, "Invoice Number")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(xIndent+200, 150, invoiceSummary.InvoiceNumber)
	pdf.SetTextColor(180, 180, 180)
	pdf.Text(xIndent+200, 180, "Date of Issue")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(xIndent+200, 200, invoiceSummary.InvoiceIssueDate)

	pdf.SetTextColor(180, 180, 180)
	pdf.Text(xIndent+400, 130, "Invoice Total")
	pdf.SetTextColor(80, 80, 80)
	pdf.SetFont("arial", "", 30)
	pdf.Text(xIndent+400, 160, invoiceSummary.InvoiceTotal)

	w, _ := pdf.GetPageSize()

	pdf.Rect(20, 230, w-40, 3.0, "F")
}

func DrawInvoiceTable(pdf *gofpdf.Fpdf, items []InvoiceItem, invoiceSummary *InvoiceSummary, totalAmount string, xIndent float64, yStart float64, bannerHt float64, drawGrid bool) {
	initYstart := yStart
	columnZero, columnOne, columnTwo, columnThree, columnFour := xIndent, xIndent+25, xIndent+170, xIndent+320, xIndent+470
	w, _ := pdf.GetPageSize()

	pdf.SetTextColor(180, 180, 180)
	pdf.SetFont("times", "", 14)
	pdf.Text(columnZero, yStart, "S/N")
	pdf.Text(columnOne, yStart, "Description")
	pdf.Text(columnTwo, yStart, "Price Per Unit")
	pdf.Text(columnThree, yStart, "Quantity")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(columnFour, yStart, "Amount")

	pdf.SetFillColor(120, 120, 120)
	for index, value := range items {
		yStart = yStart + 25
		sn := strconv.Itoa(index + 1)
		pdf.Text(columnZero, yStart, sn)
		pdf.Text(columnOne, yStart, value.Description)
		pdf.Text(columnTwo, yStart, value.PricePerUnit)
		pdf.Text(columnThree, yStart, value.PricePerUnit)
		pdf.Text(columnFour, yStart, value.Amount)
		yStart = yStart + 15
		pdf.Rect(xIndent, yStart, w-40, 1.0, "F")

		if (index+1)%12 == 0 {
			// draw the grid on current page before going to new page if draw grid is set to true
			if drawGrid {
				DrawGrid(pdf)
			}
			// create new page
			pdf.AddPage()
			pdf.SetFont("arial", "B", 38)
			pdf.SetTextColor(255, 0, 0)

			// Draw top banner on new page
			DrawTopBanner(pdf, bannerHt, xIndent)
			// Draw bottom stip on new page
			DrawBottomStrip(pdf, bannerHt)
			// Draw invoice summary
			DrawInvoiceSummary(pdf, invoiceSummary, xIndent)
			yStart = initYstart
			pdf.SetTextColor(180, 180, 180)
			pdf.SetFont("times", "", 14)
			pdf.Text(columnZero, yStart, "S/N")
			pdf.Text(columnOne, yStart, "Description")
			pdf.Text(columnTwo, yStart, "Price Per Unit")
			pdf.Text(columnThree, yStart, "Quantity")
			pdf.SetTextColor(80, 80, 80)
			pdf.Text(columnFour, yStart, "Amount")

			pdf.SetFillColor(120, 120, 120)
		}
	}
	yStart = yStart + 30
	pdf.SetTextColor(180, 180, 180)
	pdf.Text(columnThree, yStart, "Subtotal")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(columnFour, yStart, totalAmount)
	yStart = yStart + 15
	pdf.SetTextColor(180, 180, 180)
	pdf.Text(columnThree, yStart, "Tax")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(columnFour, yStart, "$0.00")
	yStart = yStart + 15
	pdf.SetTextColor(180, 180, 180)
	pdf.Text(columnThree, yStart, "Total")
	pdf.SetTextColor(80, 80, 80)
	pdf.Text(columnFour, yStart, totalAmount)
}
