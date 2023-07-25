
generate-api:
	oapi-codegen -package handlers -generate chi-server ./api/life-sheet.yaml > ./pkg/handlers/api.gen.go