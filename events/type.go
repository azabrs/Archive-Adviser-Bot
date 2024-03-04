package events


type Fetcher interface{
	Fetch(limit int) ([]Event, error)
}


type Processor interface{
	Procces(e Event)
}

type Type int

const(
	Unknown = iota
	Message
)

type Event struct{
	Type Type
	Text string
	Meta interface{}
}