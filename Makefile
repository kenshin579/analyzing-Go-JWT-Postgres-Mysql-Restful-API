

# TESTS
.PHONY: docker-compose-test
docker-compose-test:
	docker-compose -f docker-compose-mysql-test.yml up --build --abort-on-container-exit