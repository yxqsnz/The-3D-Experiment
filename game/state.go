package game

type State struct {
	IsShowingDebugInfo bool
	Text               TextManager
}

func InitializeState() State {
	return State{Text: TextManager{}}
}
