package server

type StatusMap map[string]string

type ServerStatus struct {
	ServerName string
	unixTimestamp int32
	statusMap StatusMap
}

