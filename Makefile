test:
	cd src && go test ./...

coverage:
	cd src && go test -cover ./...

# Requires modd: https://github.com/cortesi/modd
serve-api:
	cd src/api && modd