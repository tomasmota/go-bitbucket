.PHONY: bitbucketserver
bitbucketserver:
	docker run --name "bitbucket" -v bitbucketVolume:/var/atlassian/application-data/bitbucket -p 80:7990 -d atlassian/bitbucket

.PHONY: integration-tests
integration-tests:
	go run ./examples/project/main.go
	go run ./examples/repo/main.go