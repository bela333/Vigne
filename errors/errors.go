package errors

import "errors"

var NoConfig = errors.New("couldn't find configuration in the database")
var CreatedConfig = errors.New("couldn't find configuration in the database. Created default one")