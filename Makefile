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

deploy:
	cp deployments/job.yaml job.yaml.tmp ;\
	sed -i "s/TAG/$(GIT_REV_HEAD)/" job.yaml.tmp ;\
	gcloud run jobs replace --region $(LOCATION) job.yaml.tmp;\
	rm -f job.yaml.tmp
