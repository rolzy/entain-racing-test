package service

import (
	"github.com/rolzy/entain-racing-test/sports/db"
	"github.com/rolzy/entain-racing-test/sports/proto/sports"
	"golang.org/x/net/context"
)

type Sports interface {
	// ListRaces will return a collection of races.
	ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error)
}

// sportsService implements the Sports interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportsService.
func NewSportsService(sportsRepo db.SportsRepo) Sports {
	return &sportsService{sportsRepo}
}

func (s *sportsService) ListSports(ctx context.Context, in *sports.ListSportsRequest) (*sports.ListSportsResponse, error) {
	events, err := s.sportsRepo.List()
	if err != nil {
		return nil, err
	}

	return &sports.ListSportsResponse{Sports: events}, nil
}
