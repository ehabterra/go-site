# Go-Site

This is a sample website made using Go-Fiber, Gorm, and Postgres.
Furthermore, it has an example of end-to-end API test 
`pkg/web_api/server_test.go`.
The web application auto migrate the following schema on start:

```postgresql
CREATE TABLE go_site (
    site_id character(32) NOT NULL,
    name character varying(255) NOT NULL,
    active boolean DEFAULT true NOT NULL
);
CREATE TABLE go_site_attributes (
    site_id character(32) NOT NULL,
    key character varying(255) NOT NULL,
    value_str character varying(255),
    value_int integer
);
ALTER TABLE go_site ADD CONSTRAINT go_site_pkey PRIMARY KEY (site_id);
ALTER TABLE go_site_attributes ADD CONSTRAINT go_site_attributes_pkey PRIMARY KEY (site_id, key);
ALTER TABLE go_site_attributes ADD CONSTRAINT go_site_attributes_site_id_fkey FOREIGN KEY (site_id) REFERENCES go_site(site_id);
```

Please make sure before start the application that you 
have an open and running Postgres database, and the variables
in `.env` file is correct. If you haven't installed database
you can run bash file `run.sh` to run 2 containers one for 
Postgres and the other for pgAdmin4 to manage Postgres then 
create an empty database with the name that corresponds with
`.env` file `DB_NAME` variable.

## Project directory structure:

- `cmd/site/main.go` main function and start point
- `internal` has all application internal functionality.
- `mocks` I used mockery to generate mock for service that may
be used in future testing.
-  `pkg` contains api server and routes, also it contains
   an end-to-end test for the API `server_test.go` it creates
   a go-site and check if it's created correctly and then delete it.
- `Makefile` to run shortcut commands, you can use `make run`
to start the application.
- `run.sh` to run Postgres containers

