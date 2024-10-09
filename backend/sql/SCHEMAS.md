# Schemas
Migrations should be stored in /schema

Goose is used for this, run in order starting at 001


## Updating the database schemas
Use goose up / down to run the up/down sections defined in the schema files

Update the queries as needed for the changes

sqlc generate should then be run to generate the new internal/databases code

Update the handlers/whatever other code needs to be updated to complete the updates

New schemas should be versioned numerically
