package color

func ColorSelector(str string) string {
	switch str {
	case "black":
		return "\x1b[30m"
	case "red":
		return "\x1b[31m"
	case "green":
		return "\x1b[32m"
	case "yellow":
		return "\x1b[33m"
	case "blue":
		return "\x1b[34m"
	case "magenta":
		return "\x1b[35m"
	case "cyan":
		return "\x1b[36m"
	case "white":
		return "\x1b[37m"
	case "reset":
		return "\x1b[0m"
	case "brightBlack":
		return "\x1b[30;1m"
	case "brightRed":
		return "\x1b[31;1m"
	case "brightGreen":
		return "\x1b[32;1m"
	case "brightYellow":
		return "\x1b[33;1m"
	case "brightBlue":
		return "\x1b[34;1m"
	case "brightMagenta":
		return "\x1b[35;1m"
	case "brightCyan":
		return "\x1b[36;1m"
	case "brightWhite":
		return "\x1b[37;1m"
	}
	return "\x1b[37m"
}
