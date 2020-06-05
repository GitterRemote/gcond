package configuration

// Group is the a group of configuration
type Group interface {
	Evaluate(ctx Context) (conf Configuration)
}

// group is a group of Configurations
type group []Configuration

// Evaluate the configuration
func (g *group) Evaluate(ctx Context) (conf Configuration) {
	for _, conf := range *g {
		ok := conf.Condition.Evaluate(ctx)
		if ok {
			return conf
		}
	}
	return
}

// NewGroup a group
func NewGroup(confs ...Configuration) Group {
	g := group(confs)
	return &g
}
