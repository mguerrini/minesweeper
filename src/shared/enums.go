package shared

type GameStatusType string
type CellType string
type CellMarkType string


const(
	GameStatus_Created = "GameStatus_Created"
	GameStatus_Playing = "GameStatus_Playing"
	GameStatus_Lost    = "GameStatus_Lost"
	GameStatus_Won     = "GameStatus_Won"

	CellType_Mine    = "CellType_Mine"
	CellType_Number  = "CellType_Number"
	CellType_Empty   = "CellType_Empty"
	CellType_Unknown = "CellType_Unknown"

	CellMarkType_None    	= "CellMarkType_None"
	CellMarkType_Flag    	= "CellMarkType_Flag"
	CellMarkType_Question   = "CellMarkType_Question"
)