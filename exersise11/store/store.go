package store

import (
	"calandary/helpers"
	"encoding/json"
	"time"
)

type Store struct {
	Events map[time.Time][]Event `json:"data"`
}
type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

func (s *Store) CreateNew(data []byte) bool {
	var New_event Event
	err := json.Unmarshal(data, &New_event)
	if !helpers.CheckErorr(err) {
		return false
	}
	s.Events[New_event.Date] = append(s.Events[New_event.Date], New_event)
	return true
}

func (s *Store) GetEventsByDay(t time.Time) ([]Event, bool) {
	d, m, y := ParsData(t)
	for data := range s.Events {
		buff_d, buff_m, buff_y := ParsData(data)
		if buff_d == d && buff_m == m && buff_y == y {
			return s.Events[data], true
		}
	}
	return nil, false
}

func (s *Store) GetEventsByMonth(t time.Time) ([]Event, bool) {
	response := make([]Event, 0)
	for data := range s.Events {
		if data.Equal(t) || (data.After(t) && data.Before(t.Add(30*24*time.Hour))) {
			response = append(response, s.Events[data]...)
		}
	}
	if len(response) == 0 {
		return nil, false
	} else {
		return response, true
	}
}

func (s *Store) GetEventsByWeek(t time.Time) ([]Event, bool) {
	response := make([]Event, 0)
	for data := range s.Events {
		if data.Equal(t) || (data.After(t) && data.Before(t.Add(7*24*time.Hour))) {
			response = append(response, s.Events[data]...)
		}
	}
	if len(response) == 0 {
		return nil, false
	} else {
		return response, true
	}
}

func ParsData(t time.Time) (int, string, int) {
	return t.Year(), t.Month().String(), t.Day()
}
