package game


type Room struct {
        Game *Game
        Name string
        Msg map[string]string
}

func (r *Room) Label()string {
        return r.Name
}

func (r *Room) Type()string {
        return "комнаты"
}

func (r *Room) Links()[]Link {
        return r.Game.Links
}

func (r *Room) LinkedWith(rto *Room)bool {
        for _, link := range r.Links() {
                if link.Rfrom == r && link.Rto == rto {
                        return true
                }
        }
        return false
}
