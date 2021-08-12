package app

const (
	directionNorth = "north"
	directionSouth = "south"
	directionEast  = "east"
	directionWest  = "west"
)

type CharacterActionMove struct {
	character *Character
	gameMap   *GameMap
}

func (ca *CharacterActionMove) Act(directions []string) {
	characterX, characterY := ca.character.XY()
	for _, direction := range directions {
		// TODO[petr]: log invalid direction
		switch direction {
		case directionNorth:
			if ca.destinationAllowed(characterX, characterY-1) {
				ca.character.SetXY(characterX, characterY-1)
			}
		case directionSouth:
			if ca.destinationAllowed(characterX, characterY+1) {
				ca.character.SetXY(characterX, characterY+1)
			}
		case directionEast:
			if ca.destinationAllowed(characterX+1, characterY) {
				ca.character.SetXY(characterX+1, characterY)
			}
		case directionWest:
			if ca.destinationAllowed(characterX-1, characterY) {
				ca.character.SetXY(characterX-1, characterY)
			}
		}
	}
}

func (ca *CharacterActionMove) destinationAllowed(x int, y int) bool {
	destination := ca.gameMap.FindTileByXY(x, y)
	if destination == nil {
		return false
	}

	return destination.Type != TileTypeWater
}
