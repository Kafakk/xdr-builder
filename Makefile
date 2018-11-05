test:
	dep ensure
	go test -cover -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

test-report: test
	go tool cover -html=.coverage.txt
