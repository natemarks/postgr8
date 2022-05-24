.DEFAULT_GOAL := help

# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
DEFAULT_BRANCH := main
VERSION := 0.0.0
COMMIT := $(shell git rev-parse HEAD)

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DEFAULT_BRANCH := main
VENV := deployments/.venv

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'


clean-venv: ## re-create virtual env
	[[ -e $(VENV) ]] && rm -rf $(VENV); \
	python3 -m venv $(VENV); \
	( \
       . $(VENV)/bin/activate; \
       pip install --upgrade pip; \
       pip install -r deployments/requirements.txt; \
    )

pylint: ## run pylint on python files
	( \
       . $(VENV)/bin/activate; \
       find . -type f -name "*.py"  -not -path "./$(VENV)/*" | xargs pylint --max-line-length=90; \
    )

black: ## use black to format python files
	( \
       . deployments/.venv/bin/activate; \
       find . -type f -name "*.py" -not -path "./deployments/.venv/*" | xargs black; \
    )

clean-cache: ## clean python adn pytest cache data
	@find . -type f -name "*.py[co]" -delete -not -path "./deployments/.venv/*"
	@find . -type d -name __pycache__ -not -path "./deployments/.venv/*" -exec rm -rf {} \;
	@rm -rf .pytest_cache


unittest: ## run test that don't require deployed resources
	go test -v ./... -tags unit

deploymenttest: ##  run all tests
	go test -v ./...

static: ## run black and pylint
	( \
			 gofmt -w  -s .; \
			 test -z $(go vet ./...); \
			 goimports -w .; \
			 test -z $(gocyclo -over 25 .); \
    )

lint:  ##  run golint
	( \
			 go install golang.org/x/lint/golint@latest; \
			 golint ./...; \
    )

git-status: ## require status is clean so we can use undo_edits to put things back
	@status=$$(git status --porcelain); \
	if [ ! -z "$${status}" ]; \
	then \
		echo "Error - working directory is dirty. Commit those changes!"; \
		exit 1; \
	fi

db-synth: ## Deploy RDS test DB
	( \
       . deployments/.venv/bin/activate; \
       cd deployments; \
       cdk diff; \
       cdk synth; \
    )

db-create: ## Deploy RDS test DB
	( \
       . deployments/.venv/bin/activate; \
       cd deployments; \
       cdk diff; \
       cdk deploy; \
    )

db-destroy: ## Deploy RDS test DB
	( \
       . deployments/.venv/bin/activate; \
       cd deployments; \
       cdk destroy; \
    )

.PHONY: build static test artifact	