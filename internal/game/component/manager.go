package component

type Manager struct {
	Positions  map[Entity]*Position
	Velocities map[Entity]*Velocity
	Renders    map[Entity]*Render
	Tags       map[Entity]*Tag
}

func NewManager() *Manager {
	return &Manager{
		Positions:  make(map[Entity]*Position),
		Velocities: make(map[Entity]*Velocity),
		Renders:    make(map[Entity]*Render),
		Tags:       make(map[Entity]*Tag),
	}
}
