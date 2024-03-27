test:
	rm -f cover.out cover.html
	go test -cover -coverprofile cover.out -coverpkg .,./model,./internal/multimodel . ./model ./internal/multimodel && \
	go tool cover -html=cover.out -o cover.html

functest:
	go test -test.v internal/functest
