package bulma

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/tags"
	"github.com/markbates/inflect"
)

func buildOptions(opts tags.Options, err bool) {
	if opts["class"] == nil {
		opts["class"] = ""
	}

	if opts["tag_only"] != true {
		opts["class"] = strings.Join([]string{fmt.Sprint(opts["class"]), "input"}, " ")
	}

	if err {
		opts["class"] = strings.Join([]string{fmt.Sprint(opts["class"]), "is-danger"}, " ")
	}

	delete(opts, "hide_label")
}

func divWrapper(opts tags.Options, fn func(opts tags.Options) tags.Body) *tags.Tag {
	divClass := "field"
	hasErrors := false
	errors := []string{}

	if opts["errors"] != nil && len(opts["errors"].([]string)) > 0 {
		divClass = "field is-danger"
		hasErrors = true
		errors = opts["errors"].([]string)
		delete(opts, "errors")
	}

	div := tags.New("div", tags.Options{
		"class": divClass,
	})

	if opts["label"] == nil && opts["tags-field"] != nil {
		if tf, ok := opts["tags-field"].(string); ok {
			tf = strings.Join(strings.Split(tf, "."), " ")
			opts["label"] = inflect.Titleize(tf)
		}
	}

	delete(opts, "tags-field")

	useLabel := opts["hide_label"] == nil
	if useLabel && opts["label"] != nil {
		div.Append(tags.New("label", tags.Options{
			"body": opts["label"],
			"class": "label",
		}))

		delete(opts, "label")
	}

	buildOptions(opts, hasErrors)

	if opts["tag_only"] == true {
		return fn(opts).(*tags.Tag)
	}
	controldiv := tags.New("div", tags.Options{
		"class": "control",
	})
	controldiv.Append(fn(opts))
	div.Append(controldiv)

	if hasErrors {
		for _, err := range errors {
			div.Append(tags.New("p", tags.Options{
				"class": "help is-danger",
				"body":  err,
			}))
		}
	}
	return div
}
