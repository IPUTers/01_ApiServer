# API_Server
IPUTers APIリポジトリ

## コマンド群
```sh
# コンテナ立ち上げ
make run

# ボリューム削除 (イメージは残ります)
make down

# 例) go mod tidy をコンテナ内で行いたい場合
docker exec ApiServer-app go mod tidy

# 例) postgresに接続する場合
docker exec -it ApiServer-db psql -U postgres -d develop
```