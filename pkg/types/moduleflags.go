package types

type ModuleFlag uint8
type ModuleFlags []ModuleFlag

const (
	DISABLE_GLOBAL_SEARCH ModuleFlag = iota
	HIDDEL
)

func (flags ModuleFlags) Has(flag ModuleFlag) bool {
	for _, f := range flags {
		if f == flag {
			return true
		}
	}

	return false
}
