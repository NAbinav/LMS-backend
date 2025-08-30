package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Initiate_DB() error {
	// Get DB URL from env
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgresql://postgres.tlxspcbwdtgzudlqhrcr:PVY2Pj8aVBPYRUGd@aws-1-ap-south-1.pooler.supabase.com:6543/postgres"
	}

	// Parse config
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("error parsing db config: %w", err)
	}

	// ✅ Disable prepared statements for PgBouncer (use simple protocol)
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	// Limit pool size (adjust to Supabase plan)
	config.MaxConns = 10

	// Create pool
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("error creating db pool: %w", err)
	}

	// Test connection
	if err := DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("error pinging db: %w", err)
	}

	fmt.Println("✅ Database connected successfully")
	return nil
}
