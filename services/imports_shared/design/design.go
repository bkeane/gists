package design

import (
	"github.com/bkeane/gists/design/shared"

	. "goa.design/goa/v3/dsl"
)

var _ = Service("AsJson", func() {
	Method("index", func() {
		Description("Returns all rows")
		Result(ArrayOf(shared.Row))
		HTTP(func() {
			GET("/")
		})
	})
})
