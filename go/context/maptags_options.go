// Code generated by "go-option -type=mapTags"; DO NOT EDIT.

package context

var _default_mapTags_value = func() (val mapTags) { return }()

// A MapTagsOption sets options.
type MapTagsOption interface {
	apply(*mapTags)
}

// EmptyMapTagsOption does not alter the configuration. It can be embedded
// in another structure to build custom options.
//
// This API is EXPERIMENTAL.
type EmptyMapTagsOption struct{}

func (EmptyMapTagsOption) apply(*mapTags) {}

// MapTagsOptionFunc wraps a function that modifies mapTags into an
// implementation of the MapTagsOption interface.
type MapTagsOptionFunc func(*mapTags)

func (f MapTagsOptionFunc) apply(do *mapTags) {
	f(do)
}

// sample code for option, default for nothing to change
func _MapTagsOptionWithDefault() MapTagsOption {
	return MapTagsOptionFunc(func(*mapTags) {
		// TODO nothing to change
	})
}

func (o *mapTags) ApplyOptions(options ...MapTagsOption) *mapTags {
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt.apply(o)
	}
	return o
}
