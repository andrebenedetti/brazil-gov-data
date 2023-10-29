package data_loaders

import "errors"

var ErrorDataSourceChanged = errors.New("data source expectations changed since the data pipeline was implemented")
