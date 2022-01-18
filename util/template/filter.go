package template

import (
	"github.com/flosch/pongo2/v4"

	"meta-go/util/timeutil"
)

func init() {
	pongo2.RegisterFilter("timesince", FilterTimeSince)
}

func FilterTimeSince(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(timeutil.TimeSince(in.Time())), nil
}
