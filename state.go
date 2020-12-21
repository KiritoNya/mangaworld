package mangaworld

//State is a type that defines the state of manga.
type State string

const (
	Deleted   State = "Cancellato"
	Dropped   State = "Droppato"
	Finish    State = "Finito"
	Paused    State = "In pausa"
	Releasing State = "In corso"
)

//Function for mangaworld search pre-custom function.
func searchState(stat State) string {
	switch stat {
	case Deleted:
		return "canceled"
	case Dropped:
		return "dropped"
	case Finish:
		return "completed"
	case Paused:
		return "paused"
	case Releasing:
		return "ongoing"
	}
	return ""
}
