package shared

import (
	. "goa.design/goa/v3/dsl"
)

var Row = ResultType("application/vnd.goa.row", "Row", func() {
	Description("Shared/Common Resource MediaType")

	Attributes(func() {
		Field(0, "id", String, "Unique identifier a row")
		Field(1, "first", String, "First name of a row")
		Field(2, "last", String, "Last name of a row")
		Field(3, "email", String, "Email address of a row")
		Field(4, "ssn", String, "Social Security Number of a row")
	})

	View("benign", func() {
		Description("Decidedly safe view")
		Field(0, "id")
		Field(1, "first")
		Field(1, "last")
	})

	Meta("struct:pkg:path", "types")
})
