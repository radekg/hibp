package dataimport

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/jmoiron/sqlx"

	// import postgres
	_ "github.com/lib/pq"
)

// Command is the cobra command.
var Command = &cobra.Command{
	Use:   "data-import",
	Short: "Import the password file to the table",
	RunE:  run,
}

type commandConfig struct {
	dsn        string
	pwdFile    string
	noTruncate bool
}

var config = new(commandConfig)

func initFlags() {
	Command.Flags().StringVar(&config.dsn, "dsn", "", "Database connection string")
	Command.Flags().StringVar(&config.pwdFile, "password-file", "", "Password file path")
	Command.Flags().BoolVar(&config.noTruncate, "no-truncate", false, "If set, do not truncate the table before import")
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

	if !config.noTruncate {
		_, sqlErr := db.Exec("truncate table hibp restart identity")
		if sqlErr != nil {
			fmt.Fprintln(os.Stderr, "error truncating SQL table", sqlErr)
			os.Exit(1)
		}
	}

	file, err := os.Open(config.pwdFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening password file", err)
		os.Exit(1)
	}
	defer file.Close()

	currentLine := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine = currentLine + 1
		parts := strings.Split(strings.TrimSpace(scanner.Text()), ":")
		if len(parts) != 2 {
			fmt.Fprintln(os.Stderr, "line", currentLine, "skipped, split by ':' did not result in 2 items")
			continue
		}
		hash := parts[0]
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "line", currentLine, "skipped, error converting assumed count", parts[1], " as integer", err)
			continue
		}

		_, sqlErr := db.Exec("insert into hibp (`prefix`,`hash`,`count`) values ($1, $2, $3)", hash[0:5], hash, count)
		if sqlErr != nil {
			fmt.Fprintln(os.Stderr, "line", currentLine, "no inserted because of an SQL error", sqlErr)
		}

		if currentLine%1000 == 0 {
			fmt.Println("imported ", currentLine, "no inserted because of an SQL error", sqlErr)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}