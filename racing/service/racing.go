package service

import (
	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"golang.org/x/net/context"
    "fmt"
)

type Racing interface {
	// ListRaces will return a collection of races.
	ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
    fmt.Printf("%#v\n", in.View)

    // Not too sure how to proceed with the resource view from here.
    // Tried FieldMasks, JSON omitempty and map[string]interface{} but 
    // couldn't quite get it to work. I could create a new Race resource with
    // only the name and number field, but I would have to repeat so much code.
	races, err := s.racesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}
    if in.View == 1 {
        for _, element := range races {
            races2 := map[string]interface{} {
                "name": element.Name,
                "number": element.Number,
            }
            fmt.Printf("%#v\n", races2)
        }
    }

	return &racing.ListRacesResponse{Races: races}, nil
}
