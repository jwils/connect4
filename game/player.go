package game

type Player interface {
	Move(ConnectFourBoard) uint8
}
