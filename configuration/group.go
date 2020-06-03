package configuration

// Group is a group of Configurations
type Group []Configuration

// Evaluate the configuration
func (g *Group) Evaluate(ctx Context) (conf Configuration) {
	for _, conf := range *g {
		ok := conf.Condition.Evaluate(ctx)
		if ok {
			return conf
		}
	}
	return
}

// NewGroup a group
func NewGroup(confs ...Configuration) *Group {
	g := Group(confs)
	return &g
}
