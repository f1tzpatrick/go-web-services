
.PHONY: frontend-dependencies backend-dependencies run-frontend run-backend

frontend-dependencies:
	@set -e; \
		cd frontend/inventory-mgmt; \
		npm install -g @angular/cli@9.0.2; \
		npm install

backend-dependencies:
	@set -e; \
		cd backend/webservice; \
		go get -u github.com/go-sql-driver/mysql

run-frontend:
	@set -e; \
		cd frontend/inventory-mgmt; \
		ng serve --open

run-backend:
	@set -e; \
		cd backend/webservice; \
		go run main.go
