package ascii

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
	case "orange":
		return "\x1b[38;5;208m"
	case "brown":
		return "\x1b[38;5;130m"
	case "azure":
		return "\x1b[38;5;39m"
	case "ivory":
		return "\x1b[38;5;231m"
	case "teal":
		return "\x1b[38;5;6m"
	case "silver":
		return "\x1b[38;5;7m"
	case "purple":
		return "\x1b[38;5;129m"
	case "navy blue":
		return "\x1b[38;5;17m"
	case "pea green":
		return "\x1b[38;5;148m"
	case "gray":
		return "\x1b[38;5;246m"
	case "maroon":
		return "\x1b[38;5;52m"
	case "charcoal":
		return "\x1b[38;5;238m"
	case "aquamarine":
		return "\x1b[38;5;122m"
	case "coral":
		return "\x1b[38;5;209m"
	case "fuchsia":
		return "\x1b[38;5;198m"
	case "wheat":
		return "\x1b[38;5;223m"
	case "lime":
		return "\x1b[38;5;154m"
	case "crimson":
		return "\x1b[38;5;160m"
	case "khaki":
		return "\x1b[38;5;185m"
	case "hot pink":
		return "\x1b[38;5;205m"
	case "olden":
		return "\x1b[38;5;202m"
	case "plum":
		return "\x1b[38;5;96m"
	case "olive":
		return "\x1b[38;5;58m"
	case "reset":
		return "\x1b[0m"
	}
	return "\x1b[37m" // Default to white if the color is not found
}