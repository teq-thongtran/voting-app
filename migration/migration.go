package migration

//nolint:revive
import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"

	"myapp/config"
)

func Up(db *gorm.DB) {
	getDB, err := db.DB()
	if err != nil {
		fmt.Println("MIGRATION ERROR: ", err)
	}

	driver, err := mysql.WithInstance(getDB, &mysql.Config{MigrationsTable: "migration"})
	if err != nil {
		fmt.Println("MIGRATION WITHINSTANCE: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migration", config.GetConfig().MySQL.DBName, driver)
	if err != nil {
		fmt.Println("MIGRATION WithDatabaseInstance: ", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("MIGRATION OTHER ERROR: ", err)
	}

	fmt.Println("Up done!")
}
