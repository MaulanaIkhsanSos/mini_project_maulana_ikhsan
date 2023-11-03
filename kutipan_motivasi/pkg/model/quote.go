package model

// Quote adalah entitas yang merepresentasikan kutipan motivasi dalam aplikasi.
type Quote struct {
	KutipanID  int    `json:"kutipanID"`
	IsiKutipan string `json:"isiKutipan"`
	Tanggal    string `json:"tanggal"`
}

// NewQuote adalah fungsi untuk membuat kutipan motivasi baru.
func NewQuote(kutipanID int, isiKutipan, tanggal string) *Quote {
	return &Quote{
		KutipanID:  kutipanID,
		IsiKutipan: isiKutipan,
		Tanggal:    tanggal,
	}
}
