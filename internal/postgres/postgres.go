package postgres

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ctxKey string

const dbCtxKey ctxKey = "db"

type Postgres struct {
	TestApi TestApi
}

func New() *Postgres {
	return &Postgres{
		TestApi: &testApi{},
	}
}

func Connect() (*gorm.DB, error) {
	host := "db" // 	docker-composeの場合、 各サービスは独自のネットワークで実行されるため、他のサービスを探すときはサービス名で探せる。
	user := "postgres"
	password := "postgres"
	dbname := "develop"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbname, port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Close(db *gorm.DB) {
	d, errInDefer := db.DB()
	if errInDefer != nil {
		panic(errInDefer)
	}
	_ = d.Close()
}

func DBFromCtx(ctx context.Context) *gorm.DB {
	return ctx.Value("db").(*gorm.DB)
}
