package repo

import (
	"context"
	"fmt"

	"github.com/akash-scalent/gotodo/configs"
	"github.com/akash-scalent/gotodo/models"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type TodoRepo struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

func NewTodoRepo(logger *zerolog.Logger) TodoRepository {
	dbLogger := logger.With().Str("component", "database").Logger()
	// newLogger := gormlogger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	gormlogger.Config{
	// 		SlowThreshold:              time.Nanosecond,   // Slow SQL threshold
	// 		LogLevel:                   gormlogger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,          // Disable color
	// 	},
	// )
	dbLogger.Info().Msg("Connecting to database")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", configs.Config.Host, configs.Config.User, configs.Config.Password, configs.Config.Dbname, configs.Config.DPort, configs.Config.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormlogger.Default.LogMode(gormlogger.Info)})
	if err != nil {
		dbLogger.Fatal().Err(err).Msg("Error connecting to database")
	}
	dbLogger.Info().Msg("Database connection done")
	// Migration
	db.AutoMigrate(&models.Todo{})
	return &TodoRepo{db: db, logger: &dbLogger}
}

type TodoRepository interface {
	AddTodo(ctx context.Context, todo *models.Todo) error
	GetAllTodos(ctx context.Context) []*models.Todo
	GetTodoByID(ctx context.Context, todoID int) *models.Todo
}
