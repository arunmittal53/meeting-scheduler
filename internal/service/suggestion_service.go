package service

import (
	"meeting-scheduler/internal/model"
	"time"
)

func (s *SchedulerService) SuggestSlots(eventID string) ([]model.SlotSuggestion, error) {
	event, err := s.ensureEventExists(eventID)
	if err != nil {
		return nil, err
	}

	availMap := s.availabilityRepo.GetByEvent(eventID)
	if len(availMap) == 0 {
		return nil, nil
	}

	required := time.Duration(event.DurationMin) * time.Minute
	step := 15 * time.Minute

	var best []model.SlotSuggestion
	bestCount := 0

	for _, slot := range event.Slots {
		for start := slot.Start; start.Add(required).Before(slot.End) || start.Add(required).Equal(slot.End); start = start.Add(step) {
			end := start.Add(required)
			window := model.Slot{Start: start, End: end}

			var available []string
			for userID, av := range availMap {
				if isUserAvailableForExactWindow(window, av.Slots) {
					available = append(available, userID)
				}
			}

			if len(available) > bestCount {
				bestCount = len(available)
				best = []model.SlotSuggestion{{
					Slot:             window,
					UnavailableUsers: getMissingUsers2(event.Participants, available),
				}}
			} else if len(available) == bestCount && bestCount > 0 {
				best = append(best, model.SlotSuggestion{
					Slot:             window,
					UnavailableUsers: getMissingUsers2(event.Participants, available),
				})
			}
		}
	}

	return best, nil
}
