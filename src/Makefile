test:
	go test ./...

coverage:
	go test -cover ./...

mock:
	minimock -i github.com/docker/docker/client.CommonAPIClient -o gateway/dockergateway/docker_client_mock_test.go
	minimock -i ./gateway/containergateway.Container -o api/handler

# Requires modd: https://github.com/cortesi/modd
serve-api:
	cd api && modd