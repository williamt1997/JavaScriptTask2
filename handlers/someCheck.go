package handlers

func SomeCheck() map[string]interface{} {
	check := 1+1 == 2
	if check {
		return map[string]interface{}{
			"result": false,
		}
	}
	return map[string]interface{}{
		"result": true,
	}
}
