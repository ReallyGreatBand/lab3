package plants

import (
	"database/sql"
	"fmt"
)

type Plant struct {
	Id                int     `json:"id"`
	SoilMoistureLevel float64 `json:"soilMoistureLevel"`
	SoilDataTimestamp string  `json:"SoilDataTimestamp"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListPlants() ([]*Plant, error) {
	rows, err := s.Db.Query("select * from plants p where p.soil_moisture_level < 0.2;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Plant
	for rows.Next() {
		var p Plant
		if err := rows.Scan(&p.Id, &p.SoilMoistureLevel, &p.SoilDataTimestamp); err != nil {
			return nil, err
		}
		res = append(res, &p)
	}
	if res == nil {
		res = make([]*Plant, 0)
	}
	return res, nil
}

func (s *Store) UpdatePlant(id int, soilMoistureLevel float64) error {
	rows, err := s.Db.Query("select * from plants p where id = $1", id)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return fmt.Errorf("Record with given id does not exist")
	}
	_, err = s.Db.Exec("update plants set soil_moisture_level = $1 where id = $2",
		+soilMoistureLevel, id)
	return err
}
