
guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi

check: guard-project
	npx nx test $(project)

check-all:
	npx nx run-many --target=test --all

tidy: guard-project
	npx nx tidy $(project)

dep-graph:
	npx nx graph

run:
	docker-compose up -d

stop:
	docker-compose down