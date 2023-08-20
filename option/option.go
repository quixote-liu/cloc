package option

type Options struct {
	OuterOption *outerOption
	OrderOption *orderOption
	SortOption  *sortOption
}

func New() *Options {
	return &Options{
		OuterOption: newOuterOption(),
		OrderOption: newOrderOption(),
		SortOption:  newSortOption(),
	}
}

func (o *Options) ExtractArguments(args []string) ([]string, error) {
	args, err := o.OuterOption.extractArgs(args)
	if err != nil {
		return nil, err
	}

	args, err = o.OrderOption.extractArgs(args)
	if err != nil {
		return nil, err
	}

	args, err = o.SortOption.extractArgs(args)
	if err != nil {
		return nil, err
	}

	return args, nil
}

func stringsContains(src []string, target string) bool {
	for _, s := range src {
		if target == s {
			return true
		}
	}
	return false
}
