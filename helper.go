package table

func ScanRows[T any](rows IRows, callback func(row IRow) (T, error)) ([]T, error) {
	var objs []T
	for rows.Next() {
		obj, err := callback(rows)
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}

	return objs, nil
}
