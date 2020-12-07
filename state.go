package mangaworld

type State string

const (
	Deleted   State = "Cancellato"
	Dropped   State = "Droppato"
	Finish    State = "Finito"
	Paused    State = "In pausa"
	Releasing State = "In corso"
)
