package weather

import "time"

type MockWeatherRepo struct {
}

func (mwr *MockWeatherRepo) GetSummary(time time.Time) (*DailySummary, error) {
	return &DailySummary{
		Date:    time,
		Low:     1.0,
		High:    10.0,
		Average: 5.0,
	}, nil
}

func NewMockWeatherRepo() *MockWeatherRepo {
	return &MockWeatherRepo{}
}
