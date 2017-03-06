package game


type Player struct {
        InRoom * Room
}

func (p *Player)MoveTo(r *Room)string {
        if !p.InRoom.LinkedWith(r) {
                msg, ok := p.InRoom.Msg["notlinked"]; if !ok {
                        panic(HaveNotMsg(r, "notlinked"))
                } else {
                        return msg
                }
        }
        if msg, ok := r.Msg["enter"]; !ok {
                panic(HaveNotMsg(r, "enter") )
        } else {
                p.InRoom = r
                return msg
        }
}
