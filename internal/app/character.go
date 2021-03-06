package app

type Character struct {
	x int
	y int
}

func NewCharacter(x int, y int) *Character {
	return &Character{
		x: x,
		y: y,
	}
}

func (c *Character) SetXY(x int, y int) {
	c.x = x
	c.y = y
}

func (c *Character) XY() (int, int) {
	return c.x, c.y
}
