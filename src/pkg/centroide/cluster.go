package centroide

type Cluster struct {
	IdCluster  string             `json:"clusterId"`
	Directions []ClusterDirection `json:"directions"`
}
