package shared

type GameStatusType string
type CellType string

const(
	GameStatus_Created = "GameStatus_Created"
	GameStatus_Playing = "GameStatus_Playing"
	GameStatus_Lost    = "GameStatus_Lost"
	GameStatus_Won     = "GameStatus_Won"
)

const(
	CellType_Bomb    = "CellType_Bomb"
	CellType_Number  = "CellType_Number"
	CellType_Empty   = "CellType_Empty"
	CellType_Unknown = "CellType_Unknown"
)