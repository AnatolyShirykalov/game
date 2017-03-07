package game

import (
	"fmt"
)

type Player struct {
	InRoom *Room
}

func (p *Player) MoveTo(r *Room) string {
	if !p.InRoom.LinkedWith(r) {
		msg, ok := p.InRoom.Msg["notlinked"]
		if !ok {
			panic(HaveNotMsg(p.InRoom, "notlinked"))
		} else {
			return msg
		}
	}
        if !p.InRoom.UnlockedLinkTo(r) {
                if msg, ok := r.Msg["locked"]; !ok {
                        panic(HaveNotMsg(r, "locked"))
                } else {
                        return msg
                }
        }
	if msg, ok := r.Msg["enter"]; !ok {
		panic(HaveNotMsg(r, "enter"))
	} else {
		p.InRoom = r
		for i, link := range r.Links() {
			if i == 0 {
				msg = fmt.Sprintf("%s можно пройти - ", msg)
			} else {
				msg = fmt.Sprintf("%s, ", msg)
			}
			name := link.Name
			if len(name) == 0 {
				name = link.Rto.Name
			}
			msg = fmt.Sprintf("%s%s", msg, name)
		}
		return msg
	}
}
