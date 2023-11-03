package controller

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// QuoteController adalah kontroler untuk mengelola kutipan motivasi.
type QuoteController struct {
	Quotes []model.Quote
}

// CreateQuote digunakan untuk membuat kutipan motivasi baru.
func (qc *QuoteController) CreateQuote(c echo.Context) error {
	isiKutipan := c.FormValue("isiKutipan")
	tanggal := c.FormValue("tanggal")

	newQuote := model.NewQuote(len(qc.Quotes)+1, isiKutipan, tanggal)
	qc.Quotes = append(qc.Quotes, *newQuote)
	fmt.Println("Kutipan motivasi berhasil dibuat.")
	return c.JSON(http.StatusOK, newQuote)
}

// GetQuotes digunakan untuk mengambil semua kutipan motivasi.
func (qc *QuoteController) GetQuotes(c echo.Context) error {
	return c.JSON(http.StatusOK, qc.Quotes)
}

// SetupQuoteRoutes digunakan untuk menambahkan rute-rute HTTP untuk kutipan motivasi.
func SetupQuoteRoutes(e *echo.Echo, qc *QuoteController) {
	e.POST("/quotes", qc.CreateQuote)
	e.GET("/quotes", qc.GetQuotes)
}
