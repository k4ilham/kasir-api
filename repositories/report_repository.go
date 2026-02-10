package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetReport(startDate, endDate string) (*models.DailyReport, error) {
	var report models.DailyReport

	whereClause := "DATE(created_at) = CURRENT_DATE"
	whereClauseT := "DATE(t.created_at) = CURRENT_DATE"
	args := []interface{}{}

	if startDate != "" && endDate != "" {
		whereClause = "DATE(created_at) BETWEEN $1 AND $2"
		whereClauseT = "DATE(t.created_at) BETWEEN $1 AND $2"
		args = append(args, startDate, endDate)
	}

	// Total Revenue
	queryRevenue := fmt.Sprintf("SELECT COALESCE(SUM(total_amount), 0) FROM transactions WHERE %s", whereClause)
	err := repo.db.QueryRow(queryRevenue, args...).Scan(&report.TotalRevenue)
	if err != nil {
		return nil, err
	}

	// Total Transactions
	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM transactions WHERE %s", whereClause)
	err = repo.db.QueryRow(queryCount, args...).Scan(&report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Top Selling Product
	queryTop := fmt.Sprintf(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE %s
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`, whereClauseT)
	err = repo.db.QueryRow(queryTop, args...).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = models.TopSellingProduct{Nama: "-", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	return &report, nil
}
