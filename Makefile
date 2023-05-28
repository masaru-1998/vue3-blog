exec-web:
	docker compose exec web sh

exec-db:
	docker-compose exec db psql -U user -d database