package disque

// Job represents job/message returned from a Disque server.
type Job struct {
	ID                   string
	Data                 string
	Queue                string
	Nacks                int64
	AdditionalDeliveries int64
}

type Queue struct {
	Name string
	Len  int64
}
