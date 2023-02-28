package db

type DriverMode string

const (
	MasterMode  DriverMode = "master"
	ReplicaMode DriverMode = "replica"
)
