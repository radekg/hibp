package serve

import (
	"fmt"
	"os"

	"github.com/radekg/hibp/api/server"
	"github.com/radekg/hibp/api/server/restapi"
	"github.com/radekg/hibp/api/server/restapi/range_restapi"
	"github.com/spf13/cobra"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

// Command is the cobra command.
var Command = &cobra.Command{
	Use:   "serve",
	Short: "Serves the HTTP hash check endpoint",
	RunE:  run,
}

type commandConfig struct {
	bindHost string
	bindPort int
	schemes  []string
}

var config = new(commandConfig)

func initFlags() {
	Command.Flags().StringVar(&config.bindHost, "host", "127.0.0.1", "Host to bind the API on")
	Command.Flags().IntVar(&config.bindPort, "port", 15000, "Port to bind the API on")
	Command.Flags().StringSliceVar(&config.schemes, "scheme", []string{"http"}, "Enabled schemes")
}

func init() {
	initFlags()
}

func run(cmd *cobra.Command, _ []string) error {

	doc, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading Swagger file", err)
		os.Exit(1)
	}

	api := restapi.NewSelfHostedHIBPPasswordHashCheckerAPI(doc)
	api.RangeRestapiRangeSearchHandler = range_restapi.RangeSearchHandlerFunc(func(rsp range_restapi.RangeSearchParams) middleware.Responder {
		if len(rsp.HashPrefix) != 5 {
			return range_restapi.NewRangeSearchBadRequest()
		}
		return nil
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
