###########
## BUILD ##
###########
build-oauth-consumer:
	@docker build -t notes-app-oauth-consumer -f Dockerfile.prod oauth_consumer/

build-oauth-server:
	@docker build -t notes-app-oauth-server -f Dockerfile.prod oauth_server/

build-resource-provider:
	@docker build -t notes-app-resource-provider -f Dockerfile.prod resource_provider/

build: build-oauth-consumer build-oauth-server build-resource-provider

#########
## TAG ##
#########
tag-oauth-consumer:
	@docker tag notes-app-oauth-consumer registry.heroku.com/notes-app-oauth-consumer/web

tag-oauth-server:
	@docker tag notes-app-oauth-server registry.heroku.com/notes-app-oauth-server/web

tag-resource-provider:
	@docker tag notes-app-resource-provider registry.heroku.com/notes-app-resource-provider/web

tag: tag-oauth-consumer tag-oauth-server tag-resource-provider

##########
## PUSH ##
##########
push-oauth-consumer:
	@docker push registry.heroku.com/notes-app-oauth-consumer/web

push-oauth-server:
	@docker push registry.heroku.com/notes-app-oauth-server/web

push-resource-provider:
	@docker push registry.heroku.com/notes-app-resource-provider/web

push: push-oauth-consumer push-oauth-server push-resource-provider

#############
## RELEASE ##
#############
release-oauth-consumer:
	@heroku container:release web --app notes-app-oauth-consumer

release-oauth-server:
	@heroku container:release web --app notes-app-oauth-server

release-resource-provider:
	@heroku container:release web --app notes-app-resource-provider

release: release-oauth-consumer release-oauth-server release-resource-provider


############
## DEPLOY ##
############

deploy-oauth-consumer: build-oauth-consumer tag-oauth-consumer push-oauth-consumer release-oauth-consumer
deploy-oauth-server: build-oauth-server tag-oauth-server push-oauth-server release-oauth-server
deploy-resource-provider: build-resource-provider tag-resource-provider push-resource-provider release-resource-provider

deploy: build tag push release
