package spandoc

import (
	"fmt"

	"cloud.google.com/go/spanner/spansql"
)

type Document struct {
	Tables []*Table
}

type Table struct {
	Name     string
	Comments []string
	Columns  []*Column
}

type Column struct {
	Name                 string
	Type                 string
	NotNull              bool
	PrimaryKey           bool
	Comments             []string
	AllowCommitTimestamp bool
}

// Build builds Document from schema of Cloud Spanner.
func Build(filename string, src []byte) (*Document, error) {
	ddl, err := spansql.ParseDDL(filename, string(src))
	if err != nil {
		return nil, fmt.Errorf("failed to parse DDL: %w", err)
	}
	doc := &Document{}
	for _, stmt := range ddl.List {
		switch n := stmt.(type) {
		case *spansql.CreateTable:
			t := &Table{
				Name:     string(n.Name),
				Comments: comments(ddl, n),
			}
			pkMap := make(map[spansql.ID]bool, len(n.PrimaryKey))
			for _, k := range n.PrimaryKey {
				pkMap[k.Column] = true
			}
			for _, c := range n.Columns {
				t.Columns = append(t.Columns, &Column{
					Name:       string(c.Name),
					Type:       c.Type.SQL(),
					NotNull:    c.NotNull,
					PrimaryKey: pkMap[c.Name],
					Comments:   comments(ddl, c),
					AllowCommitTimestamp: func() bool {
						b := c.Options.AllowCommitTimestamp
						return b != nil && *b
					}(),
				})
			}
			doc.Tables = append(doc.Tables, t)
		case *spansql.CreateIndex:
			// TODO
		default:
			return nil, fmt.Errorf("BUG: unknown statement: %s", stmt.SQL())
		}
	}
	return doc, nil
}

func comments(ddl *spansql.DDL, n spansql.Node) []string {
	var comments []string
	if comment := ddl.LeadingComment(n); comment != nil {
		comments = append(comments, comment.Text...)
	}
	if comment := ddl.InlineComment(n); comment != nil {
		comments = append(comments, comment.Text...)
	}
	return comments
}
