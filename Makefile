define \n


endef
VERSION := $(file < ./version.txt)
PACKAGE := github.com/MrRainbow0704/DnD
LDFLAGS := -ldflags="-X '$(PACKAGE)/internal/version.version=$(VERSION)-dev'"
LDFLAGS_R := -ldflags="-X '$(PACKAGE)/internal/version.version=$(VERSION)' -H windowsgui"
AIR_CONF := $(subst ${\n}${\n},${\n},$(subst @LDFLAGS@,$(subst ",\",$(LDFLAGS)),$(file < ./.air.toml)))
.PHONY: build release clear web live pre-go go go_r

build: clear web sqlc go

release: clear web sqlc go_r

web:
	cd ./web && npm install && npm run build

sqlc: 
	go tool sqlc generate

pre-go: 
	go mod download
	go mod tidy

go: pre-go
	go build $(LDFLAGS) -o ./bin/ ./cmd/...

go_r: pre-go
	go mod verify
	go build $(LDFLAGS_R) -o ./bin/ ./cmd/...

live: pre-go
	bash -c "mkdir -p tmp"
	echo. > ./tmp/.air.toml
	echo $(subst ${\n}, >> ./tmp/.air.toml ${\n}echo ,$(AIR_CONF)) >> ./tmp/.air.toml
	go tool air -c ./tmp/.air.toml

clear:
	bash -c "rm -rf ./bin"
	bash -c "rm -rf ./tmp"
	bash -c "rm -rf ./web/build"
