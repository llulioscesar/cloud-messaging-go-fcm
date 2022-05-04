#!make
define setup_env
	$(eval ENV_FILE := .env)
	@echo "Setup env $(ENV_FILE)"
	$(eval include .env)
	$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' $(ENV_FILE)))
endef



.PHONY: run
run:
	$(call setup_env)
	go run main.go