package helpers

// OK is used as a generic json return object for ok status
type OK struct {
	OK bool
}

// NewOK returns an instanciated OK object
func NewOK() OK {
	return OK{OK: true}
}

// CheckErr is a generic error method used everywhere
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
