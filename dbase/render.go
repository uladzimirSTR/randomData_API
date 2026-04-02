package dbase

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func RenderTemplateFromFile(templatePath string, data any) (string, error) {
	b, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("read template %q: %w", templatePath, err)
	}

	tmpl, err := template.New(templatePath).Parse(string(b))
	if err != nil {
		return "", fmt.Errorf("parse template %q: %w", templatePath, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute template %q: %w", templatePath, err)
	}

	return buf.String(), nil
}
