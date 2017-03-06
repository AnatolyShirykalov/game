package game

import (
	"fmt"
)

type Game struct {
	Rooms   map[string]Room
	Links   []Link
	Players []Player
	Aliases map[string]string
}

func (g *Game) GetRoom(name string) *Room {
	r, ok := g.Rooms[name]
	if ok {
		return &r
	} else {
		alias, ok1 := g.Aliases[name]
		if ok1 {
			return g.GetRoom(alias)
		} else {
			panic(fmt.Sprintf("Не могу найти комнату по ключу %s", name))
		}
	}
}

func InitGame() Game {
	g := Game{
		Rooms:   make(map[string]Room),
		Links:   make([]Link, 0, 20),
		Players: make([]Player, 0, 1),
		Aliases: make(map[string]string),
	}
	g.Rooms = map[string]Room{
		"кухня": Room{Game: &g, Name: "кухня", Msg: map[string]string{
			"notlinked":  "нет пути кухня",
			"enter":      "кухня, ничего интересного.",
			"lookaround": "ты находишься на кухне, на столе чай,",
		}},
		"коридор": Room{Game: &g, Name: "коридор", Msg: map[string]string{
			"notlinked":  "нет пути коридор",
			"enter":      "ничего интересного.",
			"lookaround": "ничего интересного",
		}},
		"комната": Room{Game: &g, Name: "комната", Msg: map[string]string{
			"notlinked":  "нет пути комната",
			"enter":      "ты в своей комнате",
			"lookaround": "ты находишься на кухне, на столе чай,",
		}},
		"улица": Room{Game: &g, Name: "улица", Msg: map[string]string{
			"notlinked":  "нет пути улица",
			"enter":      "на улице весна",
			"lookaround": "ты находишься на кухне, на столе чай,",
			"locked":     "дверь закрыта",
		}},
	}
	g.Aliases = map[string]string{
		"домой": "коридор",
	}
	g.Players = []Player{
		Player{InRoom: g.GetRoom("кухня")},
	}

	g.Links = []Link{
		Link{Rfrom: g.GetRoom("кухня"), Rto: g.GetRoom("коридор")},
		Link{Rfrom: g.GetRoom("коридор"), Rto: g.GetRoom("кухня")},
		Link{Rfrom: g.GetRoom("коридор"), Rto: g.GetRoom("комната")},
		Link{Rfrom: g.GetRoom("коридор"), Rto: g.GetRoom("улица"), Lock: true},
		Link{Rfrom: g.GetRoom("комната"), Rto: g.GetRoom("коридор")},
		Link{Rfrom: g.GetRoom("улица"), Rto: g.GetRoom("коридор"), Name: "домой"},
	}

	return g
}

func Run() {
	g := InitGame()
	p := g.Players[0]
	fmt.Println(p.InRoom.LinkedWith(g.GetRoom("коридор")))
	fmt.Println(p.MoveTo(g.GetRoom("коридор")))
	fmt.Println(p.MoveTo(g.GetRoom("улица")))
	fmt.Println(p.MoveTo(g.GetRoom("домой")))
}
