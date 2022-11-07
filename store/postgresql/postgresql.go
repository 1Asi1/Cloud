package postgresql

import (
	"context"
	"log"
	"math"
	"solution/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresql struct {
	config *model.Config
	pool   *pgxpool.Pool
}

func New() (*postgresql, error) {
	pool, err := Connect()
	if err != nil {
		return nil, err
	}

	return &postgresql{config: model.NewConfig(), pool: pool}, nil
}

func Connect() (*pgxpool.Pool, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctx, "postgres://postgres:cloud@localhost:5432/cloud")
	if err != nil {
		log.Fatalf("it is not possible to create a pool:%s", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("it is not possible to ping db:%s", err)
		return nil, err
	}

	return pool, nil
}

func (db *postgresql) Close() {
	db.pool.Close()
}

func (db *postgresql) Create(m *model.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := make(map[string]string)
	for _, dic := range m.Data {
		for key, value := range dic {
			data[key] = value
		}
	}

	db.pool.QueryRow(ctx, "INSERT INTO config(service,data) VALUES($1,$2)", m.Service, data)
}

func (db *postgresql) Read(m *model.Config, s string, v *float64) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := make(map[string]string)
	row := db.pool.QueryRow(ctx, "SELECT version,service,data FROM config WHERE service=$1 ORDER BY version DESC", s)

	err := row.Scan(v, &m.Service, &data)
	if err != nil {
		return err
	}

	m.Data[0] = data

	return nil
}

func (db *postgresql) Update(m *model.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := db.pool.Query(ctx, "SELECT version,service FROM config")
	if err != nil {
		return err
	}

	var version float64
	conf := model.NewConfig()
	for row.Next() {
		row.Scan(&version, &conf.Service)
	}

	data := make(map[string]string)
	version = roundFloat(version, 1)
	version += 0.1
	conf.Data = m.Data
	for _, dic := range conf.Data {
		for key, value := range dic {
			data[key] = value
		}
	}

	db.pool.QueryRow(ctx, "INSERT INTO config(version,service,data) VALUES($1,$2,$3)", roundFloat(version, 2), conf.Service, data)

	return nil
}

func (db *postgresql) Delete(cVersion float64, s string, version float64) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if cVersion != version {
		db.pool.QueryRow(ctx, "DELETE FROM config WHERE version=$1 AND service=$2", version, s)
	}

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
