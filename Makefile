.PHONY: run down stop

# Docker Composeの起動（ログをターミナルに流すために[-d]は使用しない）
run:
	docker-compose up

stop:
	docker-compose stop

# Docker Composeの停止とボリュームの削除
down:
	docker-compose down -v