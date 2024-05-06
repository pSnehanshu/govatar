include .env

generate_migration:
	@read -p "Enter the migration name: " name; \
	atlas migrate diff $$name --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://postgres/16/test?search_path=public"

migration_hash:
	atlas migrate hash --dir "file://ent/migrate/migrations"

migration_apply:
	atlas migrate apply --dir "file://ent/migrate/migrations" --url $(DATABASE_URL)

migrate:
	generate_migration

run:
	air

ent_generate:
	go generate ./ent