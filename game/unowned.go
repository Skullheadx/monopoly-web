package game

func (ctx *Context) ProcessUnowned() {
	for _, uV := range ctx.Visitors.Unowned {
		visitorID := uV.visitorID
		propID := uV.propertyID

		prop := ctx.Properties.Owners[propID.Index()]

		// TODO: trigger buy or auction

		// spaceID := prop.SpaceID
		// space := BoardSpaces[spaceID.Index()]
		//
		// var price int32 = 0
		//
		// switch space.PropertyType {
		// case TypeColor:
		// 	price = ColorProperties[space.SubIndexID].Price
		// case TypeUtility:
		// 	price = UtilityPrice
		// case TypeRailroad:
		// 	price = RailroadPrice
		// }
		//
		// ctx.AdjustPlayerMoney(visitorID, -price)

	}
}
