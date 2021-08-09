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
	tiles map[int]map[int]*Tile
}

// tileTypes returns tile types that could be generated
func tileTypes() []TileType {
	return []TileType{
		TileTypeMountain,
		TileTypeWater,
		TileTypeForest,
	}
}

func NewGameMap(xMax int, yMax int) *GameMap {
	return &GameMap{
		tiles: createTiles(xMax, yMax),
	}
}

func createTiles(xMax int, yMax int) map[int]map[int]*Tile {
	borderX := xMax + 1
	borderY := yMax + 1

	var tileType TileType
	tiles := make(map[int]map[int]*Tile)
	for x := 0; x <= borderX; x++ {
		for y := 0; y <= borderY; y++ {
			if x == borderX || y == borderY {
				tileType = TileTypeCliff
			}

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
