package color

func ColorSelector(str string) string {

	switch str {
	case "white":
		return "\x1b[37m"
	case "black":
		return "\x1b[30m"
	}
	return "\x1b[37m"
}
