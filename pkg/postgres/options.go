package postgres

import (
	"fmt"
	"os"
)

func getConfigurationPostgres() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("pst_user"),
		os.Getenv("pst_password"), os.Getenv("pst_host"), os.Getenv("pst_port"),
		os.Getenv("pst_dbname"))
}
