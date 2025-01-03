# DB と app を起動
up:
	docker compose up -d db app

# migrate コンテナでマイグレーション
migrate:
	docker compose run --rm migrate

# シード投入
seed:
	docker compose run --rm seed-runner

# 全部落とす
down:
	docker compose down