package serve

import (
	"fmt"
	"os"
	"strings"

	"github.com/radekg/hibp/api/server"
	"github.com/radekg/hibp/api/server/restapi"
	"github.com/radekg/hibp/api/server/restapi/range_restapi"
	"github.com/radekg/hibp/model"
	"github.com/spf13/cobra"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/jmoiron/sqlx"

	// import postgres
	_ "github.com/lib/pq"
)

// Command is the cobra command.
var Command = &cobra.Command{
	Use:   "serve",
	Short: "Serves the HTTP hash check endpoint",
	RunE:  run,
}

type commandConfig struct {
	dsn      string
	bindHost string
	bindPort int
	schemes  []string
}

var config = new(commandConfig)

func initFlags() {
	Command.Flags().StringVar(&config.dsn, "dsn", "", "Database connection string")
	Command.Flags().StringVar(&config.bindHost, "host", "127.0.0.1", "Host to bind the API on")
	Command.Flags().IntVar(&config.bindPort, "port", 15000, "Port to bind the API on")
	Command.Flags().StringSliceVar(&config.schemes, "scheme", []string{"http"}, "Enabled schemes")
}

func init() {
	initFlags()
}

func run(cmd *cobra.Command, _ []string) error {

	db, err := sqlx.Connect("postgres", config.dsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error establishing database connection", err)
		os.Exit(1)
	}
	defer db.Close()

	doc, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading Swagger file", err)
		os.Exit(1)
	}

	api := restapi.NewSelfHostedHIBPPasswordHashCheckerAPI(doc)
	api.RangeRestapiRangeSearchHandler = range_restapi.RangeSearchHandlerFunc(func(rsp range_restapi.RangeSearchParams) middleware.Responder {

		// make sure input is correct:
		if len(rsp.HashPrefix) != 5 {
			return range_restapi.NewRangeSearchBadRequest()
		}

		// select items from the database:
		rows, err := db.NamedQuery("select \"hash\", \"count\" from hibp where \"prefix\"=:prefix",
			map[string]interface{}{
				"prefix": strings.ToUpper(rsp.HashPrefix),
			})
		if err != nil {
			fmt.Fprintln(os.Stderr, "error while executing SQL query", err)
			return range_restapi.
				NewRangeSearchInternalServerError().
				WithPayload("error while executing SQL query")
		}

		var sb strings.Builder

		row := model.Row{}
		foundRows := 0

		// read data out of the SQL result:
		for rows.Next() {
			foundRows = foundRows + 1
			err := rows.StructScan(&row)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error while processing response record", err)
				return range_restapi.
					NewRangeSearchInternalServerError().
					WithPayload("error while processing response record")
			}
			sb.WriteString(fmt.Sprintf("%s:%d\n", row.Hash, row.Count))
		}

		// if nothing found, return 404
		// according to the data documentation from HiBP, this should never be the case when
		// full data set is imported, but we should handle this anyway
		if foundRows == 0 {
			return range_restapi.NewRangeSearchNotFound()
		}

		// everything went okay, return records
		return range_restapi.NewRangeSearchOK().
			WithPayload(sb.String())
	})

	s := server.NewServer(api)
	s.Host = config.bindHost
	s.Port = config.bindPort
	s.EnabledListeners = config.schemes
	if err := s.Serve(); err != nil {
		return err
	}

	return nil
}
