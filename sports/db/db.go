package db

import (
	"time"
	"syreclabs.com/go/faker"
)

func (s *sportsRepo) seed() error {
	statement, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, type TEXT, name TEXT, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

    sports_array := []string{"Football","Australian Rules Football","Rugby League","Cricket","American Football","Basketball","Baseball","Tennis","Rugby Union","Motor Sports"}
	for i := 1; i <= 100; i++ {
		statement, err = s.db.Prepare(`INSERT OR IGNORE INTO sports(id, type, name, advertised_start_time) VALUES (?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.RandomChoice(sports_array),
				faker.Company().CatchPhrase(),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
