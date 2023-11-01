package controller

import (
	"fmt"
	"motivation-app/kutipan_motivasi/pkg/model"
)

// QuoteController adalah kontroler untuk mengelola kutipan motivasi.
type QuoteController struct {
	Quotes []model.Quote
}

// CreateQuote digunakan untuk membuat kutipan motivasi baru.
func (qc *QuoteController) CreateQuote(isiKutipan, tanggal string) {
	newQuote := model.NewQuote(len(qc.Quotes)+1, isiKutipan, tanggal)
	qc.Quotes = append(qc.Quotes, *newQuote)
	fmt.Println("Kutipan motivasi berhasil dibuat.")
}

// GetQuotes digunakan untuk mengambil semua kutipan motivasi.
func (qc *QuoteController) GetQuotes() []model.Quote {
	return qc.Quotes
}
