package expr

var operators = map[string]interface{}{}

type operator struct{}

func (o *operator) And(values ...bool) bool {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}

// newInExp create a string contains bool expression
func (o *operator) In(item string, items []string) bool {
	//refer: https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
	// sort.Strings(items)
	// i := sort.SearchStrings(items, item)
	// return i < len(items) && items[i] == item
	panic("")
}
