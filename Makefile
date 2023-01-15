.PHONY: bitbucketserver
bitbucketserver:
	podman run --name "bitbucket" -v bitbucketVolume:/var/atlassian/application-data/bitbucket -p 7990:7990 -d atlassian/bitbucket


.PHONY: integration-tests
integration-tests:
	go test ./...
