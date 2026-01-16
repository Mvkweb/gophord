package main

import (
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
	Name    string       `json:"name"`
	Doc     string       `json:"doc"`
	Structs []DocStruct  `json:"structs"`
	Types   []DocType    `json:"types"`
	Funcs   []DocFunc    `json:"funcs"`
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
	Recv      string `json:"recv,omitempty"` // Receiver if method
}

type Docs struct {
	Packages []DocPackage `json:"packages"`
}

func main() {
	// Root of the repo (assuming running from repo root or cmd/docgen)
	// We want to walk "pkg"
	root := "./pkg"
	if _, err := os.Stat(root); os.IsNotExist(err) {
		// Try going up if we are in cmd/docgen
		root = "../../pkg"
	}

	docs := Docs{Packages: []DocPackage{}}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		// Skip root pkg dir itself
		if path == root {
			return nil
		}

		// Parse the directory
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil // Skip unparsable
		}

		for _, pkg := range pkgs {
			// We only want the package that matches the directory name usually, or the main one
			// Use go/doc to get nice docs
			d := doc.New(pkg, path, 0)

			p := DocPackage{
				Name:    d.Name,
				Doc:     strings.TrimSpace(d.Doc),
				Structs: []DocStruct{},
				Types:   []DocType{},
				Funcs:   []DocFunc{},
			}

			// Process Types (Structs and others)
			for _, t := range d.Types {
				// Check if it's a struct
				isStruct := false
				var fields []DocField

				// Find the AST declaration to check implementation details not in doc.Type
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
											// Embedded field
											fieldName = fmt.Sprintf("%s", field.Type)
											if idx := strings.LastIndex(fieldName, "."); idx != -1 {
												fieldName = fieldName[idx+1:]
											}
										}
										
										// Get type string (simplified)
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
                                // It's a type alias or other type
                                p.Types = append(p.Types, DocType{
                                    Name: t.Name,
                                    Doc: strings.TrimSpace(t.Doc),
                                    Type: "type alias", // Simplified
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
                
                // Methods for this type
                for _, m := range t.Methods {
                     p.Funcs = append(p.Funcs, DocFunc{
                         Name: m.Name,
                         Doc: strings.TrimSpace(m.Doc),
                         Recv: t.Name,
                         Signature: fmt.Sprintf("func (%s) %s", t.Name, m.Name), // simplified signature
                     })
                }
			}

			// Process Standalone Funcs
			for _, f := range d.Funcs {
				p.Funcs = append(p.Funcs, DocFunc{
					Name:      f.Name,
					Doc:       strings.TrimSpace(f.Doc),
					Signature: fmt.Sprintf("func %s", f.Name), // simplified
				})
			}

			docs.Packages = append(docs.Packages, p)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	// Output to JSON
	// Check where to save
	outputPath := "../../docs.json" // Default relative to cmd/docgen
    if _, err := os.Stat("package.json"); err == nil {
        // If we happen to be in frontend root (unlikely for go run, but possible), just docs.json
        outputPath = "docs.json"
    } else if _, err := os.Stat("gophord.go"); err == nil {
        // In backend root
        outputPath = "../docs.json"
    }

	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating file relative path %s: %v\n", outputPath, err)
        // Fallback to absolute path or just print
        return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(docs); err != nil {
		panic(err)
	}
    
    fmt.Printf("Generated docs.json at %s with %d packages\n", outputPath, len(docs.Packages))
}
