package a

import xcontext "context"

type namedng struct {
	ctx xcontext.Context // want "found a struct that contains a context.Context field"
}
