package whenThis

type whenThis struct {
	condition        bool
	defaultString    string
	defaultInterface interface{}
}

func IsEmptyString(s string) whenThis {
	return whenThis{
		condition:     s == "",
		defaultString: s,
	}
}

func IsNull(v interface{}) whenThis {
	return whenThis{
		condition:        v == nil,
		defaultInterface: v,
	}
}

func (wt whenThis) DoThis(f func(), ) {
	if wt.condition {
		f()
	}
}

func (wt whenThis) UseThisString(d string) string {
	if wt.condition {
		return d
	} else {
		return wt.defaultString
	}
}
