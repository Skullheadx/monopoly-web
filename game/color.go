package game

func processOwnedColors() {
	for _, oCV := range OwnedColorVisitors {
		visitorID := oCV.visitorID
		ownerID := oCV.ownerID
		colorID := oCV.colorID

		prop := ColorProperties[colorID]
		prices := ColorPropertyRents[colorID]

		var rent int32 = prices[prop.Houses]

		if prop.Houses == 0 && HasMonopoly(ownerID, prop.GroupID) {
			rent *= 2
		}

		AdjustPlayerMoney(visitorID, -rent)
		AdjustPlayerMoney(ownerID, rent)
	}
}
