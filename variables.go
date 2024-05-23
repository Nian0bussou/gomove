package main

const (
	landscape_char string = "━"
	square_char    string = "■"
	portrait_char  string = "┃"
	video_char     string = "▶"
)

var (
	Choice          uint
	Count           uint
	Succeed         uint
	Failed          uint
	Landscape_count uint
	Portrait_count  uint
	Square___count  uint
	Video____count  uint
	reset_white     string = "\033[0m"

	colors = map[string]string{
		"red":     "\033[31m",
		"blue":    "\033[34m",
		"green":   "\033[32m",
		"cyan":    "\033[36m",
		"magenta": "\033[35m",
		"purple":  "\033[35;1m",
		"grey":    "\033[37m",
		"yellow":  "\033[33m",
	}
	special_char = map[string]string{
		"land":     landscape_char,
		"square":   square_char,
		"portrait": portrait_char,
		"video":    video_char,
	}

	countTypes = map[string]*uint{
		landscape_char: &Landscape_count,
		square_char:    &Square___count,
		portrait_char:  &Portrait_count,
		video_char:     &Video____count,
	}
)
