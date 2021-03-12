package db

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"
	"github.com/rolzy/entain-racing-test/sports/proto/sports"
)

// RacesRepo provides repository access to races.
type SportsRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List() ([]*sports.Sport, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the race repository dummy data.
func (s *sportsRepo) Init() error {
	var err error

	s.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = s.seed()
	})

	return err
}

func (s *sportsRepo) List() ([]*sports.Sport, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportQueries()[sportsList]

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.scanSports(rows)
}

func (m *sportsRepo) scanSports(
	rows *sql.Rows,
) ([]*sports.Sport, error) {
	var events []*sports.Sport

	for rows.Next() {
		var sport sports.Sport
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.Type, &sport.Name, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		events = append(events, &sport)
	}

	return events, nil
}
