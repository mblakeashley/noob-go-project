build:
	go build -o bin/main main.go

run:
	go run main.go

init:
	terraform init infra

plan:
	terraform plan infra

apply:
	terraform apply --auto-approve infra

destroy:
	terraform destroy --auto-approve infra
