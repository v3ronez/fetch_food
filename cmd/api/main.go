package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/v3ronez/fetch_food/internal/data"
	"github.com/v3ronez/fetch_food/pkg"
)

type config struct {
	serverPort int
	envMode    string
	db         dbConfig
}

type dbConfig struct {
	host           string
	dbName         string
	port           int
	user           string
	password       string
	sslMode        string
	maxOpenConnect int
	maxIdleConnect int
	maxIndleTime   string
}

type application struct {
	cfg     *config
	moodels data.Models
	db      *sql.DB
}

func main() {
	if err := initEnv(); err != nil {
		panic("env dont loading correctly")
	}
	cfg := getConfig()
	app, err := initApp(cfg)
	if err != nil {
		slog.Error(err.Error())
	}
	_ = app
	httpClient := pkg.NewHttpClient()
	foodService := pkg.NewFoodService(httpClient)
	foodService.CheckForNewFiles()
}
func initApp(conf *config) (*application, error) {
	db, err := initDB(conf)
	if err != nil {
		return nil, err
	}
	app := &application{
		cfg: conf,
		db:  db,
	}
	return app, nil
}

func getConfig() *config {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConf := dbConfig{
		port:           port,
		dbName:         os.Getenv("DB_NAME"),
		host:           os.Getenv("DB_HOST"),
		user:           os.Getenv("DB_USER"),
		password:       os.Getenv("DB_PASSWORD"),
		sslMode:        os.Getenv("DB_SSL_MODE"),
		maxIndleTime:   "15m",
		maxIdleConnect: 25,
		maxOpenConnect: 25}
	return &config{
		serverPort: serverPort,
		db:         dbConf,
		envMode:    "development",
	}
}

func initDB(config *config) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		config.db.user, config.db.password,
		config.db.host, config.db.port, config.db.dbName, config.db.sslMode)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(config.db.maxOpenConnect)
	conn.SetMaxIdleConns(config.db.maxIdleConnect)
	d, err := time.ParseDuration(config.db.maxIndleTime)
	if err != nil {
		return nil, err
	}
	conn.SetConnMaxLifetime(d)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = conn.PingContext(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}

func initEnv() error {
	path, _ := os.Getwd()
	path = strings.Join([]string{path, "/.env"}, "")
	if err := godotenv.Load(path); err != nil {
		return err
	}
	return nil
}
