package gen

import (
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/template"
	"github.com/tal-tech/go-zero/tools/goctl/templatex"
)

func genTag(in string) (string, error) {
	if in == "" {
		return in, nil
	}
	output, err := templatex.With("tag").
		Parse(template.Tag).
		Execute(map[string]interface{}{
			"field": in,
		})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
