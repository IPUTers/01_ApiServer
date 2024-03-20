# API_Server
IPUTers APIリポジトリ

## コマンド群
```shell
# コンテナ立ち上げ（ログを見るためバックグラウンド動作ではありません）
make run

# ボリューム削除 (イメージは残ります)
make down

# マイグレーション（コンテナを立ち上げた状態で使用して下さい）
make migrate

# 例) go mod tidy をコンテナ内で行いたい場合
docker exec ApiServer-app go mod tidy

# 例) postgresに接続する場合
docker exec -it ApiServer-db psql -U postgres -d develop
```