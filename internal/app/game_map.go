package app

import (
	"math/rand"
)

type TileType string

const (
	TileTypeForest   TileType = "forest"
	TileTypeMountain TileType = "mountain"
	TileTypeCliff    TileType = "cliff"
	TileTypeWater    TileType = "water"
)

type Tile struct {
	Type TileType
}

type GameMap struct {
	character *Character
	tiles     map[int]map[int]*Tile
}

// tileTypes returns tile types that could be generated
func tileTypes() []TileType {
	return []TileType{
		TileTypeMountain,
		TileTypeWater,
		TileTypeForest,
	}
}

func NewGameMap(xMin int, xMax int, yMin int, yMax int, character *Character) *GameMap {
	return &GameMap{
		tiles:     createTiles(xMin, xMax, yMin, yMax),
		character: character,
	}
}

func createTiles(xMin int, xMax int, yMin int, yMax int) map[int]map[int]*Tile {
	var tileType TileType
	tiles := make(map[int]map[int]*Tile)
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			if tiles[x] == nil {
				tiles[x] = make(map[int]*Tile)
			}

			tileType = tileTypes()[rand.Intn(len(tileTypes()))]

			tiles[x][y] = &Tile{Type: tileType}
		}
	}

	return tiles
}

func (gm *GameMap) FindTileByXY(x int, y int) *Tile {
	return gm.tiles[x][y]
}
