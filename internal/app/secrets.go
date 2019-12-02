package app

import "fmt"

func (a *Application) GetSecret(name string, defaultValue string) string {
	value, exists := a.appConfig.Secrets[name]

	if !exists {
		fmt.Printf("Secret \"%s\" has not been set: default value will be used.\n", name)
		return defaultValue
	}

	if value == defaultValue {
		fmt.Printf("Secret \"%s\" is the same as the default value.\n", name)
	}

	return value
}