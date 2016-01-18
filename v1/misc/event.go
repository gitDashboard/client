package misc

type EventLevel string

const (
	EventLevel_INFO  = EventLevel("INFO")
	EventLevel_WARN  = EventLevel("WARN")
	EventLevel_ERROR = EventLevel("ERROR")
)
