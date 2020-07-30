package handler

func HandleError(err string) int {
	switch err {
	case "OAuthException":
		return 401
	default:
		return 0
	} 
}

