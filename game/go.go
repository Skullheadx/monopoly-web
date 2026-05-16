package game

func ProcessGo() {
	for _, playerID := range GoVisitors {
		Users[playerID].Money += 200
	}
}
