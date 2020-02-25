package webui

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// MatchServer Options passed during creation
type MatchServer struct {
	Port     int
	Hostname string
	// CertFile string
}

func (ms MatchServer) Setup() {
	go Start()
	http.HandleFunc("/ws", WsPage)
	http.Handle("/", http.FileServer(http.Dir("/web/build")))
	log.Warnf("Started UI client in http://%s:%d/\n", ms.Hostname, ms.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", ms.Hostname, ms.Port), nil))
	// TODO error handling
}

type matchData struct {
	Time int64 `json:"time"`
	Host
}

type Host struct {
	shotFile string
	url      string
}

// PushMatch Broadcasts given Match to all websocket clients
func (ms MatchServer) PushMatch(host Host) (err error) {
	mg := Message{
		Event:  MATCH,
		Sender: ms.Hostname,
		Content: matchData{
			time.Now().Unix(),
			host,
		},
	}
	val, err := json.Marshal(mg)
	BroadcastData(val)
	return
}
