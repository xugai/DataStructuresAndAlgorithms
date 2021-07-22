package regexp

var matched = false
var rege = "abc"
func match(text string) bool {
	rematch(0, 0, text, len(text))
	return matched
}

func rematch(ri, ti int, text string, tl int){
	if matched {
		return
	}
	if ri == len(rege) {
		if ti == tl {
			matched = true
		}
		return
	}
	if string(rege[ti]) == "*" {
		for k := 0; k < tl - ti; k++ {
			rematch(ri + 1, ti + k, text, tl)
		}
	} else if string(rege[ri]) == "?" {
		rematch(ri + 1, ti, text, tl)
		rematch(ri + 1, ti + 1, text, tl)
	} else if ti < tl && rege[ri] == text[ti] {
		rematch(ri + 1, ti + 1, text, tl)
	}
}
