# Schemas
Migrations should be stored in /schema

Goose is used for this, run in order starting at 001


## Updating the database schemas

### Modify /schema files to update table definitions.
- Use goose up / down to migrate
- Name migrations xxx_migration_name.sql
- No migrations should be in the /sql directory

### Modify /queries files to update the queries.

- sqlc generate should then be run to generate the new internal/databases code
- Queries should be in files pertaining to their respective tables
- No queries should be in the /sql directory


