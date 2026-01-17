package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type DocPackage struct {
	Name    string      `json:"name"`
	Doc     string      `json:"doc"`
	Structs []DocStruct `json:"structs"`
	Types   []DocType   `json:"types"`
	Funcs   []DocFunc   `json:"funcs"`
}

type DocStruct struct {
	Name   string     `json:"name"`
	Doc    string     `json:"doc"`
	Fields []DocField `json:"fields"`
}

type DocField struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tag  string `json:"tag"`
	Doc  string `json:"doc"`
}

type DocType struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
	Type string `json:"type"`
}

type DocFunc struct {
	Name      string `json:"name"`
	Doc       string `json:"doc"`
	Signature string `json:"signature"`
	Recv      string `json:"recv,omitempty"`
}

type DocExample struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Notes       []string `json:"notes,omitempty"`
	Code        string   `json:"code"`
}

type Docs struct {
	Packages []DocPackage `json:"packages"`
	Examples []DocExample `json:"examples"`
}

func main() {
	// Root of the repo
	root := "./pkg"
	examplesRoot := "./examples"
	if _, err := os.Stat(root); os.IsNotExist(err) {
		root = "../../pkg"
		examplesRoot = "../../examples"
	}

	docs := Docs{Packages: []DocPackage{}, Examples: []DocExample{}}

	// Parse packages
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if path == root {
			return nil
		}

		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil
		}

		for _, pkg := range pkgs {
			d := doc.New(pkg, path, 0)

			p := DocPackage{
				Name:    d.Name,
				Doc:     strings.TrimSpace(d.Doc),
				Structs: []DocStruct{},
				Types:   []DocType{},
				Funcs:   []DocFunc{},
			}

			for _, t := range d.Types {
				isStruct := false
				var fields []DocField

				for _, spec := range t.Decl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if typeSpec.Name.Name == t.Name {
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								isStruct = true
								if structType.Fields != nil {
									for _, field := range structType.Fields.List {
										fieldName := ""
										if len(field.Names) > 0 {
											fieldName = field.Names[0].Name
										} else {
											fieldName = fmt.Sprintf("%s", field.Type)
											if idx := strings.LastIndex(fieldName, "."); idx != -1 {
												fieldName = fieldName[idx+1:]
											}
										}

										typeStr := "unknown"
										if ident, ok := field.Type.(*ast.Ident); ok {
											typeStr = ident.Name
										} else if star, ok := field.Type.(*ast.StarExpr); ok {
											if ident, ok := star.X.(*ast.Ident); ok {
												typeStr = "*" + ident.Name
											} else {
												typeStr = fmt.Sprintf("*%T", star.X)
											}
										} else if sel, ok := field.Type.(*ast.SelectorExpr); ok {
											if x, ok := sel.X.(*ast.Ident); ok {
												typeStr = x.Name + "." + sel.Sel.Name
											}
										} else {
											typeStr = fmt.Sprintf("%T", field.Type)
										}

										tag := ""
										if field.Tag != nil {
											tag = strings.Trim(field.Tag.Value, "`")
										}

										fields = append(fields, DocField{
											Name: fieldName,
											Type: typeStr,
											Tag:  tag,
											Doc:  strings.TrimSpace(field.Doc.Text()),
										})
									}
								}
							} else {
								p.Types = append(p.Types, DocType{
									Name: t.Name,
									Doc:  strings.TrimSpace(t.Doc),
									Type: "type alias",
								})
							}
						}
					}
				}

				if isStruct {
					p.Structs = append(p.Structs, DocStruct{
						Name:   t.Name,
						Doc:    strings.TrimSpace(t.Doc),
						Fields: fields,
					})
				}

				for _, m := range t.Methods {
					p.Funcs = append(p.Funcs, DocFunc{
						Name:      m.Name,
						Doc:       strings.TrimSpace(m.Doc),
						Recv:      t.Name,
						Signature: fmt.Sprintf("func (%s) %s", t.Name, m.Name),
					})
				}
			}

			for _, f := range d.Funcs {
				p.Funcs = append(p.Funcs, DocFunc{
					Name:      f.Name,
					Doc:       strings.TrimSpace(f.Doc),
					Signature: fmt.Sprintf("func %s", f.Name),
				})
			}

			docs.Packages = append(docs.Packages, p)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	// Extract examples
	docs.Examples = extractExamples(examplesRoot)

	// Output to JSON
	outputPath := "docs.json"
	if _, err := os.Stat("package.json"); err == nil {
		outputPath = "docs.json"
	} else if _, err := os.Stat("gophord.go"); err == nil {
		outputPath = "docs.json"
	} else if _, err := os.Stat("../../go.mod"); err == nil {
		outputPath = "../../docs.json"
	} else {
		// Default to current directory if unsure
		outputPath = "docs.json"
	}

	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating file relative path %s: %v\n", outputPath, err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(docs); err != nil {
		panic(err)
	}

	fmt.Printf("Generated docs.json at %s with %d packages, %d examples\n", outputPath, len(docs.Packages), len(docs.Examples))
}

// extractExamples reads example folders and extracts code between DOC:START and DOC:END
func extractExamples(examplesRoot string) []DocExample {
	var examples []DocExample

	// Read examples.json manifest
	manifestPath := filepath.Join(examplesRoot, "examples.json")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		fmt.Printf("No examples.json found at %s: %v\n", manifestPath, err)
		return examples
	}

	var manifest struct {
		Examples []struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Category    string `json:"category"`
		} `json:"examples"`
	}
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		fmt.Printf("Error parsing examples.json: %v\n", err)
		return examples
	}

	for _, ex := range manifest.Examples {
		exampleDir := filepath.Join(examplesRoot, ex.ID)

		// Read example.json for notes
		var notes []string
		exampleJsonPath := filepath.Join(exampleDir, "example.json")
		if data, err := os.ReadFile(exampleJsonPath); err == nil {
			var exMeta struct {
				Notes []string `json:"notes"`
			}
			if json.Unmarshal(data, &exMeta) == nil {
				notes = exMeta.Notes
			}
		}

		// Extract code from main.go between DOC:START and DOC:END
		code := extractDocCode(filepath.Join(exampleDir, "main.go"))

		examples = append(examples, DocExample{
			ID:          ex.ID,
			Title:       ex.Title,
			Description: ex.Description,
			Category:    ex.Category,
			Notes:       notes,
			Code:        code,
		})
	}

	return examples
}

// extractDocCode extracts code between // DOC:START and // DOC:END markers
func extractDocCode(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	var lines []string
	inDocSection := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "// DOC:START") {
			inDocSection = true
			continue
		}
		if strings.Contains(line, "// DOC:END") {
			inDocSection = false
			continue
		}
		if inDocSection {
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, "\n")
}
