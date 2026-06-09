package game

type Statuses map[string]struct{}

func (s Statuses) Add(status string) {
	s[status] = struct{}{}
}

func (s Statuses) Remove(status string) {
	delete(s, status)
}

func (s Statuses) Contains(status string) bool {
	_, exists := s[status]
	return exists
}

type Player struct {
	n int
	alive bool
	name string
	team string
	alignment string
	role string
	atk int
	def int
	votes int
	statuses Statuses
}

func InitPlayers() int {
	Players = make([]Player, 15)

	for i := 0; i < 15; i++ {
		Players[i] = Player{
			n int
			alive:    true,
			name: "",
			team:     "",
			role:     "",
			alignment: "",
			atk:      0,
			def:      0,
			votes:    0,
			statuses: Statuses{},
		}
	}

	return 0
}
