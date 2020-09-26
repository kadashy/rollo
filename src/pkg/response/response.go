package response

import (
	"time"
)

type DirectionResponse struct {
	PointID  string  `json:"pointID"`
	Distance int     `json:"distanceMeters"`
	Duration float64 `json:"durationSecond"`
}

func (f *DirectionResponse) SetDistance(distance int) {
	f.Distance = distance
}
func (f *DirectionResponse) SetDuration(duration time.Duration) {
	f.Duration = duration.Seconds()
}

func (f *DirectionResponse) SetPoint(point string) {
	f.PointID = point
}
