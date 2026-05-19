package game

func (ctx *Context) ProcessTax() {
	for _, tV := range ctx.Visitors.Tax {
		playerID := tV.visitorID
		taxID := tV.taxID
		ctx.AdjustPlayerMoney(playerID, -TaxSpaces[taxID].Amount)
	}
}
