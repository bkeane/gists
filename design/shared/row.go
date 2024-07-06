package shared

import (
	. "goa.design/goa/v3/dsl"
)

var Row = ResultType("application/vnd.goa.row", "Row", func() {
	Description("Shared/Common Resource MediaType")

	Attributes(func() {
		Attribute("id", String, "Unique identifier a row")
		Attribute("first", String, "First name of a row")
		Attribute("last", String, "Last name of a row")
		Attribute("email", String, "Email address of a row")
		Attribute("ssn", String, "Social Security Number of a row")
	})

	View("benign", func() {
		Description("Decidedly safe view")
		Attribute("id")
		Attribute("first")
		Attribute("last")
	})
})
