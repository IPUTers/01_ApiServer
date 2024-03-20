.PHONY: run down stop migrate

# Docker Composeの起動（ログをターミナルに流すために[-d]は使用しない）
run:
	docker-compose up

stop:
	docker-compose stop

# Docker Composeの停止とボリュームの削除
down:
	docker-compose down -v

migrate:
	docker exec ApiServer-app go run cmd/migrate.go