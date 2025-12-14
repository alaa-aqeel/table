package table

type LoaderMany[T any] func(items []T) ([]T, error)

type LoaderOne[T any] func(item T) (T, error)

func OneToOne[T any, R any, K comparable](
	items []T,
	fk func(T) K,
	load func([]K) ([]R, error),
	ref func(R) K,
	set func(*T, R),
) ([]T, error) {

	keys := make([]K, 0, len(items))
	keySet := make(map[K]struct{})

	for _, item := range items {
		k := fk(item)
		if _, ok := keySet[k]; !ok {
			keySet[k] = struct{}{}
			keys = append(keys, k)
		}
	}

	relations, err := load(keys)
	if err != nil {
		return nil, err
	}

	relMap := make(map[K]R, len(relations))
	for _, r := range relations {
		relMap[ref(r)] = r
	}

	for i := range items {
		if r, ok := relMap[fk(items[i])]; ok {
			set(&items[i], r)
		}
	}

	return items, nil
}

func OneToMany[T any, R any, K comparable](
	items []T,
	fk func(T) K, // parent key (e.g. service.ID)
	load func([]K) ([]R, error), // load children by parent keys
	ref func(R) K, // child foreign key (e.g. booking.ServiceID)
	set func(*T, []R), // attach children to parent
) ([]T, error) {

	keys := make([]K, 0, len(items))
	keySet := make(map[K]struct{})

	for _, item := range items {
		k := fk(item)
		if _, ok := keySet[k]; !ok {
			keySet[k] = struct{}{}
			keys = append(keys, k)
		}
	}

	relations, err := load(keys)
	if err != nil {
		return nil, err
	}

	relMap := make(map[K][]R)
	for _, r := range relations {
		k := ref(r)
		relMap[k] = append(relMap[k], r)
	}

	for i := range items {
		if children, ok := relMap[fk(items[i])]; ok {
			set(&items[i], children)
		}
	}

	return items, nil
}
