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
)

type config struct {
	serverPort int
	envMode    string
	db         *dbConfig
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

var cfg *config
var dbConf *dbConfig

func main() {
	if err := initEnv(); err != nil {
		panic("env dont loading correctly")
	}
	cfg = getConfig()
	fmt.Sprintf("%+v", cfg)
	os.Exit(1)
	app, err := initApp(cfg)
	if err != nil {
		slog.Error(err.Error())
	}
	db, err := initDB(app.cfg)
	if err != nil {
		slog.Error(err.Error())

	}
	app.db = db
	os.Exit(1)
	// // TODO: move this to a another service
	// url := "https://challenges.coode.sh/food/data/json/index.txt"
	// httpClient := HttpService.New()
	// response, err := httpClient.Get(url)
	// if err != nil {
	// 	slog.Error("Somethings happen: ", err)
	// }
	// fmt.Println(string(response))
}
func initApp(conf *config) (*application, error) {
	app := &application{
		cfg: conf,
	}
	return app, nil
}

func getConfig() *config {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConf.port = port
	dbConf.dbName = os.Getenv("DB_NAME")
	dbConf.host = os.Getenv("DB_HOST")
	dbConf.user = os.Getenv("DB_USER")
	dbConf.password = os.Getenv("DB_PASSWORD")
	dbConf.sslMode = os.Getenv("DB_SSL_MODE")
	dbConf.maxIndleTime = "15m"
	dbConf.maxIdleConnect = 25
	dbConf.maxOpenConnect = 25
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
