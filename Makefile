setup:
	docker-compose up -d --build

stop:
	docker-compose down -v

start:
	docker-compose up -d

remove:
	docker-compose down --rmi all --volumes --remove-orphans

restart:
	docker-compose down --rmi all --volumes --remove-orphans
	docker-compose up -d --build

logs:
	docker-compose logs -f