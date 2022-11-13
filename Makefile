.PHONY: bitbucketserver
bitbucketserver:
	docker run --name "bitbucket" -v bitbucketVolume:/var/atlassian/application-data/bitbucket -p 80:7990 -d atlassian/bitbucket