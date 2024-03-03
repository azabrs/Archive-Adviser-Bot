package events


type Fetcher interface{
	Fetch(limit int) ([]Event, error)
}


type Proccesor interface{
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
}