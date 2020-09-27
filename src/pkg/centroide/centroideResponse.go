package centroide

type CentroideResponse struct {
	CentroideCluster string  `json:"centroideId"`
	CentroideLat     float64 `json:"centroideLatitude"`
	CentroideLon     float64 `json:"centroideLongitude"`
}

func (f *CentroideResponse) SetLat(lat float64) {
	f.CentroideLat = lat
}

func (f *CentroideResponse) SetLon(lon float64) {
	f.CentroideLon = lon
}

func (f *CentroideResponse) SetCLuster(cluster string) {
	f.CentroideCluster = cluster
}
