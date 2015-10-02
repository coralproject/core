# Database layer


#### Tasks we need to implement related with the db layer

- map query results to structs/classes (serialization)
- simple CRUD-style object persistence
- provide programmatic query construction
- create INSERT schemas or provide schema migration
- hooks into database usage (ie. logging all query times, pre/post save, etc.)
- manage database connections
    - done with dabase/sql library
- manage caching 


#### DBMS being considered

- PostgreSQL
    - Support for JSONB (that is faster to process and supports indexing)
- MongoDB
    - Washington Post is using it
    - It can help with part of the application where the structure may be changing frequently
- Neo4J
    - Popular graph database

#### Main SQL packages being used in the go ecosystem:

- gorp
    - maintained and many people contributing to it
    - bind struct fields to table columns
    - support for embedded structs
    - support for transactions
    - forward engineer db schema from structs
    - pre/post instert/update/delete hooks
    - automatically generate sql statements for a struct
    - delete/select by primary key/s
    - optional trace sql logging
    - bind arbitrary sql queries to a struct
    - bind slice to select query results without type assertions
    - supports go version 1.3 and 1.4
- gorm
    - maintained and many people contributing to it
    - Chainable API
    - Auto Migrations
    - Relations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
    - Callbacks (Before/After Create/Save/Update/Delete/Find)
    - Preloading (eager loading)
    - Transactions
    - Embed Anonymous Struct
    - Soft Deletes
    - Customizable Logger
    - Iteration Support via Rows
    - It supports postgersql/mysql and there is an extension for mongodb
- sqlx
    - it’s an extension for standard database/sql library
    - marshall rows into structs, maps and slices
    - named parameter support including prepared statement
    - get and select to go quickly from query to struct/slice
- goose
    - mantained
    - migration tool
- squirrel
    - mantained
    - build SQL from composable parts
    - can execute quires directly
- gocraft/dbr
    - maintained
    - it's an extension for standard database/sql library
    - fast performance (it does not create prepare statement and )
    - join multiple Tables
    - load struct data into the Tables
    - create subqueries
    - Identity, Union, Condition, Alias supported
    - Automatically map results to structs
    - Supports plain sql
    - It has a query builder
    - Transactions, Insert/Update records
    - It supports mysql and postgresql
- upper.io/db
    - supports mysql, postgresql
    - provides partial support for mongodb


## Coral options

### Go Packages

- GORM
  - it is the better maintained of the quasi-ORMs packages
- Squirel
- Goose
  - migration package that gorm is lacking of
- database/sql standard library



References

- Old article on some features expected for a database layer in Go http://jmoiron.net/blog/golang-orms/
- Json data types in Postgresql http://www.postgresql.org/docs/current/static/datatype-json.html
