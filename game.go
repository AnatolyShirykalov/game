package game

import (
  "fmt"
)

type Game struct {
        Rooms map[string]Room
        Links []Link
        Players []Player
}

func (g *Game)GetRoom(name string)*Room{
  r, _ := g.Rooms[name]
  return &r
}

func InitGame() Game {
        g := Game{Rooms: make(map[string]Room), Links: make([]Link, 0, 20), Players: make([]Player,0,1)}
        startRoom := Room{Game: &g, Name: "кухня", Msg: map[string]string{
                "notlinked": "notlinked",
                "enter": "enter",
        }}
        g.Rooms["кухня"] = startRoom
        g.Rooms["коридор"] = Room{Game: &g, Name: "коридор", Msg: map[string]string{
                "notlinked": "Корnotlinked",
                "enter": "Корenter",
        }}
        g.Players = []Player{
          Player{InRoom: &startRoom},
        }

        g.Links = []Link{
                Link{Rfrom: g.GetRoom("кухня"), Rto: g.GetRoom("коридор")},
                Link{Rfrom: g.GetRoom("коридор"), Rto: g.GetRoom("кухня")},
        }

        return g
}

func Run(){
        g := InitGame()
        p := g.Players[0]
        fmt.Println(p.InRoom.LinkedWith(g.GetRoom("коридор")))
        fmt.Println(p.MoveTo(g.GetRoom("коридор")))
}
