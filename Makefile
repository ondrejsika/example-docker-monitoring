up:
	docker-compose up -d

down:
	docker-compose down

reload:
	curl -X POST http://localhost:9090/-/reload
	curl -X POST http://localhost:9093/-/reload

loadgen:
	ab -n 100000 http://web1-127-0-0-1.nip.io/
