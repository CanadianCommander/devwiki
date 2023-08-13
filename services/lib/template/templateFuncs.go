package template

// ===============================================
// Public
// ===============================================

func Args(args ...interface{}) map[string]interface{} {
	argMap := make(map[string]interface{})

	if len(args)%2 != 0 {
		panic("Args must be specified as key, value pairs. Uneven number of arguments provided")
	}

	for i := 0; i < len(args); i += 2 {
		argMap[args[i].(string)] = args[i+1]
	}

	return argMap
}
