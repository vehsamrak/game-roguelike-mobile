package app

const (
	characterActionMove = "move"
)

type CharacterAction interface {
	Act(parameters []string)
}

type CharacterActionProvider struct {
	character *Character
	gameMap   *GameMap
}

func NewCharacterActionProvider(character *Character, gameMap *GameMap) *CharacterActionProvider {
	return &CharacterActionProvider{
		character: character,
		gameMap:   gameMap,
	}
}

func (p *CharacterActionProvider) ProvideAction(actionName string) CharacterAction {
	var characterAction CharacterAction

	switch actionName {
	case characterActionMove:
		characterAction = &CharacterActionMove{
			character: p.character,
			gameMap:   p.gameMap,
		}
	}

	return characterAction
}
