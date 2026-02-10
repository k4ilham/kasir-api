package models

type DailyReport struct {
	TotalRevenue   int               `json:"total_revenue"`
	TotalTransaksi int               `json:"total_transaksi"`
	ProdukTerlaris TopSellingProduct `json:"produk_terlaris"`
}

type TopSellingProduct struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
