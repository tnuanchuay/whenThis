package whenThis

type whenThis struct {
	condition     bool
	defaultString string
}

func IsEmptyString(s string) whenThis {
	return whenThis{
		condition:     s == "",
		defaultString: s,
	}
}

func (wt whenThis) UseThisString(d string) string {
	if wt.condition {
		return d
	} else {
		return wt.defaultString
	}
}
