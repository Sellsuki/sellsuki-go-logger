unit-test:
	go test -v ./...

coverage-test:
	go test -coverprofile cover.out ./...

coverage-test-html: coverage-test
	go tool cover -html=cover.out

benchmark-test:
	go test -bench=. -benchtime=10s -count 3 ./...
