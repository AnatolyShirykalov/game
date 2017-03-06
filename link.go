package game

type Link struct {
	Rfrom *Room
	Rto   *Room
	Name  string
	Lock  bool
}
