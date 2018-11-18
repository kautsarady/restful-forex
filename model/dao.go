package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// DAO data access object
type DAO struct {
	DB *sql.DB
}

// Make sql connection
func Make(dbDriver, connStr string, withDummy bool) (*DAO, error) {
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		return nil, err
	}

	if err := repopulate(db, withDummy); err != nil {
		return nil, err
	}

	return &DAO{db}, nil
}

// Track track DAO
func (d *DAO) Track(date time.Time) ([]Exchange, error) {
	rows, err := d.DB.Query("SELECT DISTINCT \"from\", \"to\" FROM history")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exs []Exchange
	for rows.Next() {
		var ex Exchange
		if err := rows.Scan(&ex.From, &ex.To); err != nil {
			return nil, err
		}
		exs = append(exs, ex)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	dt := getDateRange(date)

	for i := 0; i < len(exs); i++ {
		// modify Exchange
		modEx(d.DB, dt, &exs[i])
	}

	return exs, nil
}

func modEx(db *sql.DB, dateRng []string, ex *Exchange) {
	rows, err := db.Query("SELECT rate, \"date\" FROM history WHERE \"from\"=$1 AND \"to\"=$2 AND \"date\"=ANY($3) ORDER BY \"date\" DESC",
		ex.From, ex.To, pq.Array(dateRng),
	)
	if err != nil {
		fmt.Errorf("Error tracking: %v", err)
		return
	}
	defer rows.Close()

	var rate7 []float64
	for rows.Next() {
		var rate sql.NullFloat64
		var date pq.NullTime
		if err := rows.Scan(&rate, &date); err != nil {
			fmt.Errorf("Error tracking: %v", err)
			return
		}
		if !rate.Valid || !date.Valid {
			continue
		}
		if ex.Rate == float64(0) {
			ex.Rate = rate.Float64
			ex.Date = date.Time
		}
		rate7 = append(rate7, rate.Float64)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("Error tracking: %v", err)
		return
	}

	if len(rate7) == 7 {
		ex.Rate7 = getAvg(rate7)
	}
}

// Input Exchange
func (d *DAO) Input(ex Exchange) error {
	_, err := d.DB.Exec("INSERT INTO history VALUES ($1, $2, $3, $4)",
		ex.From, ex.To, ex.Rate, ex.Date)
	if err != nil {
		return err
	}
	return nil
}

// Trend show trend
func (d *DAO) Trend(from, to string, rng float64) ([]Exchange, error) {
	rows, err := d.DB.Query("SELECT \"date\", rate FROM history WHERE \"from\"=$1 AND \"to\"=$2 AND rate >= $3 ORDER BY \"date\" DESC LIMIT 7", from, to, rng)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exs []Exchange
	for rows.Next() {
		var ex Exchange
		if err := rows.Scan(&ex.Date, &ex.Rate); err != nil {
			return nil, err
		}
		ex.From = from
		ex.To = to
		exs = append(exs, ex)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exs, nil
}

// Add Exchange
func (d *DAO) Add(ex Exchange) error {
	_, err := d.DB.Exec("INSERT INTO history VALUES ($1, $2)", ex.From, ex.To)
	if err != nil {
		return err
	}
	return nil
}

// Remove Exchange
func (d *DAO) Remove(ex Exchange) error {
	_, err := d.DB.Exec("DELETE FROM history WHERE \"from\"=$1 AND \"to\"=$2", ex.From, ex.To)
	if err != nil {
		return err
	}
	return nil
}
