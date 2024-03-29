// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssociatedEventColumns holds the columns for the "AssociatedEvent" table.
	AssociatedEventColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "earthquake_id", Type: field.TypeInt, Nullable: true},
	}
	// AssociatedEventTable holds the schema information for the "AssociatedEvent" table.
	AssociatedEventTable = &schema.Table{
		Name:       "AssociatedEvent",
		Columns:    AssociatedEventColumns,
		PrimaryKey: []*schema.Column{AssociatedEventColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "AssociatedEvent_Earthquakes_associated_events",
				Columns:    []*schema.Column{AssociatedEventColumns[3]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EarthquakesColumns holds the columns for the "Earthquakes" table.
	EarthquakesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "tsunami", Type: field.TypeInt32, Nullable: true},
		{Name: "net", Type: field.TypeString, Nullable: true},
		{Name: "code", Type: field.TypeString, Nullable: true},
		{Name: "sources", Type: field.TypeString, Nullable: true},
		{Name: "nst", Type: field.TypeInt, Nullable: true},
		{Name: "dmin", Type: field.TypeFloat64, Nullable: true},
		{Name: "rms", Type: field.TypeFloat64, Nullable: true},
		{Name: "gap", Type: field.TypeFloat64, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "location_id", Type: field.TypeInt, Nullable: true},
		{Name: "magitude_id", Type: field.TypeInt, Nullable: true},
		{Name: "time_id", Type: field.TypeInt, Nullable: true},
	}
	// EarthquakesTable holds the schema information for the "Earthquakes" table.
	EarthquakesTable = &schema.Table{
		Name:       "Earthquakes",
		Columns:    EarthquakesColumns,
		PrimaryKey: []*schema.Column{EarthquakesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Earthquakes_Location_earthquakes",
				Columns:    []*schema.Column{EarthquakesColumns[13]},
				RefColumns: []*schema.Column{LocationColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Earthquakes_Magnitudes_earthquakes",
				Columns:    []*schema.Column{EarthquakesColumns[14]},
				RefColumns: []*schema.Column{MagnitudesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Earthquakes_Times_earthquakes",
				Columns:    []*schema.Column{EarthquakesColumns[15]},
				RefColumns: []*schema.Column{TimesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LocationColumns holds the columns for the "Location" table.
	LocationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "dept", Type: field.TypeFloat64},
		{Name: "place", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// LocationTable holds the schema information for the "Location" table.
	LocationTable = &schema.Table{
		Name:       "Location",
		Columns:    LocationColumns,
		PrimaryKey: []*schema.Column{LocationColumns[0]},
	}
	// MagnitudesColumns holds the columns for the "Magnitudes" table.
	MagnitudesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "magnitude_value", Type: field.TypeFloat64, Nullable: true},
		{Name: "magnitude_type", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// MagnitudesTable holds the schema information for the "Magnitudes" table.
	MagnitudesTable = &schema.Table{
		Name:       "Magnitudes",
		Columns:    MagnitudesColumns,
		PrimaryKey: []*schema.Column{MagnitudesColumns[0]},
	}
	// RequestsColumns holds the columns for the "Requests" table.
	RequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "request_url", Type: field.TypeString, Nullable: true},
		{Name: "request_method", Type: field.TypeString, Nullable: true},
		{Name: "request_headers", Type: field.TypeJSON, Nullable: true},
		{Name: "request_body", Type: field.TypeJSON, Nullable: true},
		{Name: "response_status_code", Type: field.TypeInt32, Nullable: true},
		{Name: "response_body", Type: field.TypeJSON, Nullable: true},
		{Name: "request_timestamp", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// RequestsTable holds the schema information for the "Requests" table.
	RequestsTable = &schema.Table{
		Name:       "Requests",
		Columns:    RequestsColumns,
		PrimaryKey: []*schema.Column{RequestsColumns[0]},
	}
	// RolesColumns holds the columns for the "Roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_name", Type: field.TypeString, Nullable: true},
	}
	// RolesTable holds the schema information for the "Roles" table.
	RolesTable = &schema.Table{
		Name:       "Roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SchemaMigrationsColumns holds the columns for the "schema_migrations" table.
	SchemaMigrationsColumns = []*schema.Column{
		{Name: "version", Type: field.TypeInt, Increment: true},
		{Name: "dirty", Type: field.TypeBool},
	}
	// SchemaMigrationsTable holds the schema information for the "schema_migrations" table.
	SchemaMigrationsTable = &schema.Table{
		Name:       "schema_migrations",
		Columns:    SchemaMigrationsColumns,
		PrimaryKey: []*schema.Column{SchemaMigrationsColumns[0]},
	}
	// TimesColumns holds the columns for the "Times" table.
	TimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_time", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// TimesTable holds the schema information for the "Times" table.
	TimesTable = &schema.Table{
		Name:       "Times",
		Columns:    TimesColumns,
		PrimaryKey: []*schema.Column{TimesColumns[0]},
	}
	// TypesColumns holds the columns for the "Types" table.
	TypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type_eathquake", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "earthquake_id", Type: field.TypeInt, Nullable: true},
	}
	// TypesTable holds the schema information for the "Types" table.
	TypesTable = &schema.Table{
		Name:       "Types",
		Columns:    TypesColumns,
		PrimaryKey: []*schema.Column{TypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Types_Earthquakes_types",
				Columns:    []*schema.Column{TypesColumns[4]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "Users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "role_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "Users" table.
	UsersTable = &schema.Table{
		Name:       "Users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Users_Roles_users",
				Columns:    []*schema.Column{UsersColumns[4]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssociatedEventTable,
		EarthquakesTable,
		LocationTable,
		MagnitudesTable,
		RequestsTable,
		RolesTable,
		SchemaMigrationsTable,
		TimesTable,
		TypesTable,
		UsersTable,
	}
)

func init() {
	AssociatedEventTable.ForeignKeys[0].RefTable = EarthquakesTable
	AssociatedEventTable.Annotation = &entsql.Annotation{
		Table: "AssociatedEvent",
	}
	EarthquakesTable.ForeignKeys[0].RefTable = LocationTable
	EarthquakesTable.ForeignKeys[1].RefTable = MagnitudesTable
	EarthquakesTable.ForeignKeys[2].RefTable = TimesTable
	EarthquakesTable.Annotation = &entsql.Annotation{
		Table: "Earthquakes",
	}
	LocationTable.Annotation = &entsql.Annotation{
		Table: "Location",
	}
	MagnitudesTable.Annotation = &entsql.Annotation{
		Table: "Magnitudes",
	}
	RequestsTable.Annotation = &entsql.Annotation{
		Table: "Requests",
	}
	RolesTable.Annotation = &entsql.Annotation{
		Table: "Roles",
	}
	TimesTable.Annotation = &entsql.Annotation{
		Table: "Times",
	}
	TypesTable.ForeignKeys[0].RefTable = EarthquakesTable
	TypesTable.Annotation = &entsql.Annotation{
		Table: "Types",
	}
	UsersTable.ForeignKeys[0].RefTable = RolesTable
	UsersTable.Annotation = &entsql.Annotation{
		Table: "Users",
	}
}
