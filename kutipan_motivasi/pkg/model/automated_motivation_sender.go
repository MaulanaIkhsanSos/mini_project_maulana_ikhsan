package model

// AutomatedMotivationSender adalah entitas yang digunakan untuk mengirim kutipan motivasi otomatis.
type AutomatedMotivationSender struct {
	PengirimID        int
	IsiMotivasi       string
	JamPengiriman     string
	TanggalPengiriman string
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
