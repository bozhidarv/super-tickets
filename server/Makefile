ENV ?= dev
-include .env
-include .env.$(ENV)
export

deploy: prod-only
	@echo deploying to prod

run:
	go run ./cmd/supertickets

prod-only:
ifneq ($( ENV), production)
	$(error ENV must be production)
endif
