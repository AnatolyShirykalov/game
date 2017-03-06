package game

type Room struct {
	Game *Game
	Name string
	Msg  map[string]string
}

func (r *Room) Label() string {
	return r.Name
}

func (r *Room) Type() string {
	return "комнаты"
}

func (r *Room) Links() []Link {
	ret := make([]Link, 0, 4)
	for _, link := range r.Game.Links {
		if link.Rfrom.Name == r.Name {
			ret = append(ret, link)
		}
	}
	return ret
}

func (r *Room) LinkedWith(rto *Room) bool {
	for _, link := range r.Links() {
		if link.Rfrom.Name == r.Name && link.Rto.Name == rto.Name {
			return true
		}
	}
	return false
}
