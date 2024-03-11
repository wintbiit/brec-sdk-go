package webhook

import "time"

const (
	EventTypeRecordStart = "SessionStarted"
	EventTypeFileOpen    = "FileOpening"
	EventTypeFileClose   = "FileClosed"
	EventTypeRecordStop  = "SessionEnded"
	EventTypeLiveStart   = "StreamStarted"
	EventTypeLiveStop    = "StreamEnded"
)

type EventMeta struct {
	EventType      string    `json:"EventType"`
	EventTimestamp time.Time `json:"EventTimestamp"`
	EventId        string    `json:"EventId"`
}

type EventRecordStart struct {
	EventMeta
	EventData struct {
		SessionId        string `json:"SessionId"`
		RoomId           int    `json:"RoomId"`
		ShortId          int    `json:"ShortId"`
		Name             string `json:"Name"`
		Title            string `json:"Title"`
		AreaNameParent   string `json:"AreaNameParent"`
		AreaNameChild    string `json:"AreaNameChild"`
		Recording        bool   `json:"Recording"`
		Streaming        bool   `json:"Streaming"`
		DanmakuConnected bool   `json:"DanmakuConnected"`
	} `json:"EventData"`
}

type EventRecordStartHandler func(evt *EventRecordStart)

type EventFileOpen struct {
	EventMeta
	EventData struct {
		RelativePath     string    `json:"RelativePath"`
		FileOpenTime     time.Time `json:"FileOpenTime"`
		SessionId        string    `json:"SessionId"`
		RoomId           int       `json:"RoomId"`
		ShortId          int       `json:"ShortId"`
		Name             string    `json:"Name"`
		Title            string    `json:"Title"`
		AreaNameParent   string    `json:"AreaNameParent"`
		AreaNameChild    string    `json:"AreaNameChild"`
		Recording        bool      `json:"Recording"`
		Streaming        bool      `json:"Streaming"`
		DanmakuConnected bool      `json:"DanmakuConnected"`
	} `json:"EventData"`
}

type EventFileOpenHandler func(evt *EventFileOpen)

type EventFileClose struct {
	EventMeta
	EventData struct {
		RelativePath     string    `json:"RelativePath"`
		FileSize         int       `json:"FileSize"`
		Duration         float64   `json:"Duration"`
		FileOpenTime     time.Time `json:"FileOpenTime"`
		FileCloseTime    time.Time `json:"FileCloseTime"`
		SessionId        string    `json:"SessionId"`
		RoomId           int       `json:"RoomId"`
		ShortId          int       `json:"ShortId"`
		Name             string    `json:"Name"`
		Title            string    `json:"Title"`
		AreaNameParent   string    `json:"AreaNameParent"`
		AreaNameChild    string    `json:"AreaNameChild"`
		Recording        bool      `json:"Recording"`
		Streaming        bool      `json:"Streaming"`
		DanmakuConnected bool      `json:"DanmakuConnected"`
	} `json:"EventData"`
}

type EventFileCloseHandler func(evt *EventFileClose)

type EventRecordStop struct {
	EventMeta
	EventData struct {
		SessionId        string `json:"SessionId"`
		RoomId           int    `json:"RoomId"`
		ShortId          int    `json:"ShortId"`
		Name             string `json:"Name"`
		Title            string `json:"Title"`
		AreaNameParent   string `json:"AreaNameParent"`
		AreaNameChild    string `json:"AreaNameChild"`
		Recording        bool   `json:"Recording"`
		Streaming        bool   `json:"Streaming"`
		DanmakuConnected bool   `json:"DanmakuConnected"`
	} `json:"EventData"`
}

type EventRecordStopHandler func(evt *EventRecordStop)

type EventLiveStart struct {
	EventMeta
	EventData struct {
		RoomId           int    `json:"RoomId"`
		ShortId          int    `json:"ShortId"`
		Name             string `json:"Name"`
		Title            string `json:"Title"`
		AreaNameParent   string `json:"AreaNameParent"`
		AreaNameChild    string `json:"AreaNameChild"`
		Recording        bool   `json:"Recording"`
		Streaming        bool   `json:"Streaming"`
		DanmakuConnected bool   `json:"DanmakuConnected"`
	} `json:"EventData"`
}

type EventLiveStartHandler func(evt *EventLiveStart)

type EventLiveStop struct {
	EventMeta
	EventData struct {
		RoomId           int    `json:"RoomId"`
		ShortId          int    `json:"ShortId"`
		Name             string `json:"Name"`
		Title            string `json:"Title"`
		AreaNameParent   string `json:"AreaNameParent"`
		AreaNameChild    string `json:"AreaNameChild"`
		Recording        bool   `json:"Recording"`
		Streaming        bool   `json:"Streaming"`
		DanmakuConnected bool   `json:"DanmakuConnected"`
	} `json:"EventData"`
}

type EventLiveStopHandler func(evt *EventLiveStop)
