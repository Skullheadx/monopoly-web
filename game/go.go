package game

func (ctx *Context) ProcessGo() {
	for _, playerID := range ctx.Visitors.Go {
		ctx.AdjustPlayerMoney(playerID, GoSalary)
	}
}
