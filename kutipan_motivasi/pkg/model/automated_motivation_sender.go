package model

// AutomatedMotivationSender adalah entitas yang digunakan untuk mengirim kutipan motivasi otomatis.
type AutomatedMotivationSender struct {
	PengirimID        int    `json:"pengirimID"`
	IsiMotivasi       string `json:"isiMotivasi"`
	JamPengiriman     string `json:"jamPengiriman"`
	TanggalPengiriman string `json:"tanggalPengiriman"`
}

// NewAutomatedMotivationSender adalah fungsi untuk membuat pengirim motivasi otomatis baru.
func NewAutomatedMotivationSender(pengirimID int, isiMotivasi, jamPengiriman, tanggalPengiriman string) *AutomatedMotivationSender {
	return &AutomatedMotivationSender{
		PengirimID:        pengirimID,
		IsiMotivasi:       isiMotivasi,
		JamPengiriman:     jamPengiriman,
		TanggalPengiriman: tanggalPengiriman,
	}
}
