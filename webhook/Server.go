package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Server struct {
	Name               string
	recordStartHandler EventRecordStartHandler
	fileOpenHandler    EventFileOpenHandler
	fileCloseHandler   EventFileCloseHandler
	recordStopHandler  EventRecordStopHandler
	liveStartHandler   EventLiveStartHandler
	liveStopHandler    EventLiveStopHandler
}

var (
	server  = http.NewServeMux()
	servers = make(map[string]*Server)
)

func NewWebhookServer(name string) *Server {
	s := &Server{
		Name: name,
	}

	server.HandleFunc("/"+name, s.handleConnection)
	servers[name] = s

	return s
}

func (s *Server) handleConnection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Body.Close()
	var evt EventMeta
	err = json.Unmarshal(body, &evt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch evt.EventType {
	case EventTypeRecordStart:
		var recordStartEvt EventRecordStart
		err = json.Unmarshal(body, &recordStartEvt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if s.recordStartHandler != nil {
			s.recordStartHandler(&recordStartEvt)
		}

	case EventTypeFileOpen:
		var fileOpenEvt EventFileOpen
		err = json.Unmarshal(body, &fileOpenEvt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if s.fileOpenHandler != nil {
			s.fileOpenHandler(&fileOpenEvt)
		}

	case EventTypeFileClose:
		var fileCloseEvt EventFileClose
		err = json.Unmarshal(body, &fileCloseEvt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if s.fileCloseHandler != nil {
			s.fileCloseHandler(&fileCloseEvt)
		}

	case EventTypeRecordStop:
		var recordStopEvt EventRecordStop
		err = json.Unmarshal(body, &recordStopEvt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if s.recordStopHandler != nil {
			s.recordStopHandler(&recordStopEvt)
		}

	case EventTypeLiveStart:
		var liveStartEvt EventLiveStart
		err = json.Unmarshal(body, &liveStartEvt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if s.liveStartHandler != nil {
			s.liveStartHandler(&liveStartEvt)
		}

	case EventTypeLiveStop:
		var liveStopEvt EventLiveStop
		err = json.Unmarshal(body, &liveStopEvt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if s.liveStopHandler != nil {
			s.liveStopHandler(&liveStopEvt)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Close() {
	delete(servers, s.Name)
}

func (s *Server) OnRecordStart(handler EventRecordStartHandler) {
	s.recordStartHandler = handler
}

func (s *Server) OnFileOpen(handler EventFileOpenHandler) {
	s.fileOpenHandler = handler
}

func (s *Server) OnFileClose(handler EventFileCloseHandler) {
	s.fileCloseHandler = handler
}

func (s *Server) OnRecordStop(handler EventRecordStopHandler) {
	s.recordStopHandler = handler
}

func (s *Server) OnLiveStart(handler EventLiveStartHandler) {
	s.liveStartHandler = handler
}

func (s *Server) OnLiveStop(handler EventLiveStopHandler) {
	s.liveStopHandler = handler
}

func StartServers(addr string) {
	server := &http.Server{
		Addr:    addr,
		Handler: server,
	}

	go func() {
		defer func() {
			for _, s := range servers {
				s.Close()
			}
		}()

		server.ListenAndServe()
	}()

	fmt.Println("Webhook server started at", addr)
	for _, s := range servers {
		fmt.Printf("Webhook server %s started: %s\n", s.Name, addr+"/"+s.Name)
	}

	select {}
}
