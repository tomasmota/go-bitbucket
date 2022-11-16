.PHONY: bitbucketserver
bitbucketserver:
	docker run --name "bitbucket" -v bitbucketVolume:/var/atlassian/application-data/bitbucket -p 80:7990 -d atlassian/bitbucket

# TODO: make this run all tests without hardcoding
# Perhaps these should just be written as actual go tests
.PHONY: integration-tests
integration-tests:
	go run ./examples/project/main.go
	go run ./examples/repo/main.go