package game

const GO_SALARY int32 = 200

func ProcessGo() {
	for _, playerID := range GoVisitors {
		AdjustPlayerMoney(playerID, GO_SALARY)
	}
}
