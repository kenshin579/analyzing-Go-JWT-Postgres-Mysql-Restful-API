
# TESTS
# docker-compose start
.PHONY: docker-compose-test-start
docker-compose-test-start:
	docker-compose -f docker-compose-mysql-test.yml up --build --abort-on-container-exit

# docker-compose stop
.PHONY: docker-compose-test-stop
docker-compose-test-stop:
	docker-compose -f docker-compose-mysql-test.yml down --remove-orphans --volumes
