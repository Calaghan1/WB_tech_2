package store

import (
	"calandary/helpers"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)


func CreateNewStore() *Store {
	ev := make(map[int][]Event)
	return &Store{
		Events: ev,
		Cap: 0,
	}
}

type Store struct {
	Events map[int][]Event `json:"data"`
	Cap int
}


type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

type Response struct {
	Result []Event `json:"result"`
}

func (s *Store)CreateJson(events []Event) ([]byte, error) {
	fmt.Println(events)
	res := Response{events}
	data, err := json.MarshalIndent(res, " ", " ")
	if !helpers.CheckError(err) {
		return nil, err
	}
	
	return data, nil
}
func (s *Store) CreateNew(userid string, title string, date time.Time) ([]byte, error) {
	user_id, err := strconv.Atoi(userid)
	if !helpers.CheckError(err) {
		return  nil, fmt.Errorf("Uncorrect user_id")
	}
	_, ok := s.Events[user_id]
	if !ok {
		s.Events[user_id] = make([]Event, 2)
	}
	s.Cap += 1
	New_event := Event{
		ID: s.Cap,
		UserID: user_id,
		Date: date,
		Title: title,
	}
	s.Events[user_id] = append(s.Events[user_id], New_event)
	return s.CreateJson([]Event{New_event})
}

func (s *Store)UpdateEvent(id_str, userid, title string, date time.Time) ([]byte, error) {
	user_id, err := strconv.Atoi(userid)
	if !helpers.CheckError(err) {
		return  nil, fmt.Errorf("Uncorrect user_id")
	}
	id, err := strconv.Atoi(id_str)
	if !helpers.CheckError(err) {
		return  nil, fmt.Errorf("Uncorrect user_id")
	}
	_, ok := s.Events[user_id]
	if !ok {
		return nil, fmt.Errorf("No such user_id")
	}
	New_event := Event{
		ID: id,
		UserID: user_id,
		Date: date,
		Title: title,
	}
	// err := json.Unmarshal(data, &New_event)
	// if !helpers.CheckError(err) {
	// 	return false
	// }
	for i := 0; i < len(s.Events[user_id]); i++ {
		if s.Events[user_id][i].ID == id {
			s.Events[user_id][i] = New_event
		}
	}
	return s.CreateJson([]Event{New_event})
}

func (s *Store)DeleteEvent(id_str, userid string) ([]byte, error) {
	user_id, err := strconv.Atoi(userid)
	if !helpers.CheckError(err) {
		return  nil, fmt.Errorf("Uncorrect user_id")
	}
	id, err := strconv.Atoi(id_str)
	if !helpers.CheckError(err) {
		return  nil, fmt.Errorf("Uncorrect id")
	}
	_, ok := s.Events[user_id]
	if !ok {
		return nil, fmt.Errorf("No such user_id")
	}
	buff := 0
	for i := 0; i < len(s.Events[user_id]); i++ {
		if s.Events[user_id][i].ID == id {
			buff = i
		}
	}
	s.Events[user_id][buff] = s.Events[user_id][len(s.Events[user_id]) - 1]
	s.Events[user_id] = s.Events[user_id][:len(s.Events[user_id]) - 1]
	return s.CreateJson([]Event{s.Events[user_id][buff]})
}



func (s *Store)SaveData() bool {
	data, err := json.MarshalIndent(s, " ", " ")
	if !helpers.CheckError(err) {
		return false
	}
	file, err := os.Create("data.json")
	if !helpers.CheckError(err) {
		return false
	}
	_, err = file.Write(data)
	if !helpers.CheckError(err) {
		return false
	}
	file.Close()
	return true
}

func (s *Store)LoadData() bool {
	file, err := os.Open("data.json")
	if !helpers.CheckError(err) {
		return false
	}
	var data []byte
	_, err = file.Read(data)
	if !helpers.CheckError(err) {
		return false
	}
	err = json.Unmarshal(data, s)
	if !helpers.CheckError(err) {
		return false
	}	
	return true
}

func (s *Store) GetEventsByDay(userid string, t time.Time) ([]byte, error) {
	user_id, err := strconv.Atoi(userid)
	if !helpers.CheckError(err) {
		return nil, fmt.Errorf("Uncorrect user_id")
	}
	d, m, y := ParsData(t)
	val, ok := s.Events[user_id] 
	if !ok {
		return nil, fmt.Errorf("Cant find user")
	}
	Response := make([]Event, 0, 1)
	for _, events := range val {
		buff_d, buff_m, buff_y := ParsData(events.Date)
		if buff_d == d && buff_m == m && buff_y == y {
			Response = append(Response, events)
		}
	}
	return s.CreateJson(Response)
}

func (s *Store) GetEventsByMonth(userid string, t time.Time) ([]byte, error) {
	user_id, err := strconv.Atoi(userid)
	if helpers.CheckError(err) {
		return nil, fmt.Errorf("Uncorrect user_id")
	}
	val, ok := s.Events[user_id] 
	if !ok {
		return nil, fmt.Errorf("Cant find user")
	}
	response := make([]Event, 0)
	for _, event := range val {
		
		if event.Date.Equal(t) || (event.Date.After(t) && event.Date.Before(t.Add(30*24*time.Hour))) {
			response = append(response, event)
		}
	}
	if len(response) == 0 {
		return s.CreateJson([]Event{})
	} else {
		return s.CreateJson(response)
	}
}

func (s *Store) GetEventsByWeek(userid string, t time.Time) ([]Event, error) {
	user_id, err := strconv.Atoi(userid)
	if helpers.CheckError(err) {
		return nil, fmt.Errorf("Uncorrect user_id")
	}
	val, ok := s.Events[user_id] 
	if !ok {
		return nil, fmt.Errorf("Cant find user")
	}
	response := make([]Event, 0)
	for _, event := range val {
		if event.Date.Equal(t) || (event.Date.After(t) && event.Date.Before(t.Add(7*24*time.Hour))) {
			response = append(response, event)
		}
	}
	if len(response) == 0 {
		return nil, nil
	} else {
		return response, nil
	}
}

func ParsData(t time.Time) (int, string, int) {
	return t.Year(), t.Month().String(), t.Day()
}
