package repository

import (
	"database/sql"
	"github.com/c0caina/inaWarp/repository/models"
	"github.com/go-gl/mathgl/mgl64"
	_ "github.com/mattn/go-sqlite3"
)

type WarpSqlite struct {
	db *sql.DB
}

func NewWarpSqlite(db *sql.DB) *WarpSqlite {
	return &WarpSqlite{db: db}
}

type Warp interface {
	Insert(warp models.Warp) error
	Delete(name string) error
	SelectName(name string) (mgl64.Vec3, error)
	SelectAll() ([]models.Warp, error)
}

func (r *WarpSqlite) Insert(warp models.Warp) error {
	_, err := r.db.Exec(`INSERT INTO warp VALUES (?1 ,?2, ?3 ,?4)`, &warp.Name, &warp.XYZ[0], &warp.XYZ[1], &warp.XYZ[2])
	if err != nil {
		return err
	}
	return nil
}

func (r *WarpSqlite) Delete(name string) error {
	_, err := r.db.Exec("DELETE FROM warp WHERE Name = ?1", name)
	if err != nil {
		return err
	}
	return nil
}

func (r *WarpSqlite) SelectName(name string) (mgl64.Vec3, error) {
	var warp models.Warp

	err := r.db.QueryRow("SELECT * FROM warp WHERE Name = ?1", name).Scan(&warp.Name, &warp.XYZ[0], &warp.XYZ[1], &warp.XYZ[2])
	if err != nil {
		return warp.XYZ, err
	}
	return warp.XYZ, nil
}

func (r *WarpSqlite) SelectAll() ([]models.Warp, error) {
	var warps []models.Warp

	row, err := r.db.Query("SELECT * FROM warp")
	if err != nil {
		return warps, err
	}

	for row.Next() {
		var warp models.Warp
		err = row.Scan(&warp.Name, &warp.XYZ[0], &warp.XYZ[1], &warp.XYZ[2])
		if err != nil {
			return warps, err
		}
		warps = append(warps, warp)
	}
	return warps, nil
}
