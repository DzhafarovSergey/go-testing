package reverse

func Reverse(s string) string {
	//if len(s) == 1 {
	//	return s, nil
	//}
	//if s == "" || s == " " {
	//	return "", errors.New("String is empty")
	//}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
