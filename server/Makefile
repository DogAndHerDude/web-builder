migrate_dev:
	npx prisma migrate dev --schema=./third_party/prisma/schema.prisma --skip-generate
dev:
	go run cmd/web_builder/web_builder.go
prod:
	go build -o build/app cmd/web_builder/web_builder.go
