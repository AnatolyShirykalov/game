package game

type Game struct {
        Rooms map[string]Room
        Links []Link
        Players map[string]Player
}

func InitGame() Game {
        g := Game{Rooms: make(map[string]Room), Links: make([]Link, 0, 20), Players: make(map[string]Player)}

        return g
}

func Run(){
        InitGame()
}
