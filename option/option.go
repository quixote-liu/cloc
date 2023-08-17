package option

type Options struct {
	FileOption  *outerOption
	OrderOption *orderOption
	SortOption  *sortOption
}

func New() *Options {
	return &Options{
		FileOption:  newOuterOption(),
		OrderOption: newOrderOption(),
		SortOption:  newSortOption(),
	}
}

func (o *Options) ExtractArguments(args []string) ([]string, error) {
	args, err := o.FileOption.extractArgs(args)
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
