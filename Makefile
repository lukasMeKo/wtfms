.PHONY: default

default:
	# Start meli search and melisearch-ui
	docker compose up -d
	# Open ui
	xdg-open http://localhost:24900 || true
	# run initial migration
	go run . migrate
	# register additional properties
	go run . register --type "property" ./test_data/property/*.json
	# register categories
	go run . register --type "category" ./test_data/category/*.json
	# add articles
	go run . add --category "glove" ./test_data/articles/glove/*.json
	go run . add --category "tube" ./test_data/articles/tube/*.json
