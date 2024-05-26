GIT_REV_HEAD := $$(git rev-parse HEAD)
LOCATION := asia-northeast1
CLOUD_REPO_URL = "$(LOCATION)-docker.pkg.dev/elite-caster-125113/household-expense/dev"

test:
	go test -v ./tests/...

build:
	docker build -t household-expense:$(GIT_REV_HEAD) .

push:
	docker tag  household-expense:$(GIT_REV_HEAD) $(CLOUD_REPO_URL):$(GIT_REV_HEAD)
	docker push $(CLOUD_REPO_URL):$(GIT_REV_HEAD)