package game

func (ctx *Context) AdjustPlayerMoney(playerID PlayerID, amount int32) {
	ctx.Players.Alive[playerID.Index()].Money += amount

	if ctx.Players.Alive[playerID.Index()].Money < 0 {
		ctx.Turn.InDebt = true
	} else { // Money >= 0
		ctx.Turn.InDebt = false
	}

}
