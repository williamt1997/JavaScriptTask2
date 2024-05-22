package main

import (
	checks "github.com/Enterprise-Automation/lms-check-ms-helper"
	handlers "github.com/Enterprise-Automation/lms-check-template-ms/handlers"
)

func main() {

	checks.NewCheck("some_check", handlers.SomeCheck)

	checks.RegisterChecks()
}
