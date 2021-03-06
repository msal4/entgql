// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssignmentsColumns holds the columns for the "assignments" table.
	AssignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "is_exam", Type: field.TypeBool, Default: false},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "duration", Type: field.TypeInt64, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// AssignmentsTable holds the schema information for the "assignments" table.
	AssignmentsTable = &schema.Table{
		Name:       "assignments",
		Columns:    AssignmentsColumns,
		PrimaryKey: []*schema.Column{AssignmentsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssignmentsTable,
	}
)

func init() {
}
