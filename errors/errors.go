package errors

import "errors"

var NoConfig = errors.New("couldn't find configuration in the database")
var NoRoles = errors.New("couldn't find role command configuration in the database")
var CreatedConfig = errors.New("couldn't find configuration in the database. Created default one")
var NoModule = errors.New("couldn't find registered module")
