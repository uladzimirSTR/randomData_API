package dbase

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func RenderCreateTableSQL(templatePath string, data TableTemplateData) (string, error) {
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("read template file %q: %w", templatePath, err)
	}

	tmpl, err := template.New("create_table").Parse(string(templateBytes))
	if err != nil {
		return "", fmt.Errorf("parse template %q: %w", templatePath, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("execute template %q: %w", templatePath, err)
	}

	return buf.String(), nil
}
