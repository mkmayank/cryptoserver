package util

type ArrayFlags []string

func (a *ArrayFlags) String() string {
	return ""
}

func (a *ArrayFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}
