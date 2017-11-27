build_all: build_api build_listing build_order build_user

build_api:
	cd ../api && go build -ldflag "-s -w"

build_listing:
	cd ../srv/listing && go build -ldflag "-s -w"

build_order:
	cd ../srv/order && go build -ldflag "-s -w"

build_user:
	cd ../srv/user && go build -ldflag "-s -w"

build:
ifndef TAG
	$(error TAG is undefined)
endif
ifndef APP
	APP=(api listing order user)
	$(foreach d,$(APP),make build TAG=$(TAG) APP=$(d);)
	exit 0
endif
ifeq ($(APP),api)
	cp ../k8s/Dockerfile Dockerfile
	cp ../$(APP)/$(APP) .
	@sed -i -r "s#\{\{app\}\}#$(APP)#g" Dockerfile
	docker build . -t $(APP):$(TAG)
	docker tag $(APP):$(TAG) $(APP):latest
	@rm Dockerfile $(APP)
else
	dist="../srv/$(APP)"
	cp ../k8s/Dockerfile Dockerfile
	cp ../srv/$(APP)/$(APP) .
	@sed -i -r "s#\{\{app\}\}#$(APP)#g" Dockerfile
	docker build . -t $(APP):$(TAG)
	docker tag $(APP):$(TAG) $(APP):latest
	@rm Dockerfile $(APP)
endif
