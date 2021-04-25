package migrate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jmoiron/sqlx"

	// import postgres
	_ "github.com/lib/pq"
)

// Command is the cobra command.
var Command = &cobra.Command{
	Use:   "migrate",
	Short: "Create the SQL table required to run this program",
	RunE:  run,
}

type commandConfig struct {
	dsn string
}

var config = new(commandConfig)

func initFlags() {
	Command.Flags().StringVar(&config.dsn, "dsn", "", "Database connection string")
}

func init() {
	initFlags()
}

func run(cmd *cobra.Command, _ []string) error {

	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", config.dsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error establishing database connection", err)
		os.Exit(1)
	}
	defer db.Close()

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	_, schemaErr := db.Exec(schema)
	if schemaErr != nil {
		fmt.Fprintln(os.Stderr, "error crating schema", schemaErr)
		os.Exit(1)
	}

	return nil
}

const schema = `
CREATE TABLE public.hibp (
	row_id serial NOT NULL,
	prefix varchar(5) NOT NULL,
	hash varchar(40) NOT NULL,
	count integer NOT NULL,
    CONSTRAINT hibp_pkey PRIMARY KEY (row_id)
);
CREATE UNIQUE INDEX hibp_prefix_idx ON public.hibp (prefix,hash,count);
`
