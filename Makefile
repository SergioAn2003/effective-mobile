.PHONY: app-start app-stop

pg_dsn := "user=postgres dbname=postgres password=songs-password host=localhost port=5431 sslmode=disable"

# app
app-start:
	@docker compose up -d --build --remove-orphans --force-recreate
app-stop:
	@docker compose down

# migrations
migrate-new:
	@goose -dir ./migrations create $(name) sql && goose -dir ./migrations fix
migrate-up:
	@goose -dir postgres $(pg_dsn) ./migrations up
migrate-down:
	@goose -dir postgres $(pg_dsn) ./migrations down 1
migrate-drop:
	@goose -dir postgres $(pg_dsn) ./migrations reset

lint:
	@golangci-lint run -v -c=./.golangci.yml
		

