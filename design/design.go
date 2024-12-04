package design

import (
	_ "github.com/bkeane/gists/services/imports_shared/design"

	. "goa.design/goa/v3/dsl"
)

var _ = API("RowFactory", func() {
	Title("Advanced System For Mapping Row Data")
	Version("1.0")
})
