# PostgreSQL
.PHONY: up log exec rm
up:
	docker compose up -d postgres
log:
	docker compose logs -f postgres
exec:
	docker compose exec postgres psql -U user -d sampledb
rm:
	docker compose rm -fsv postgres
