RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/ok/s//$$(printf "$(GREEN)ok$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

lint:
	@ golangci-lint run --path-prefix src --config ../.github/.golangci.yaml ./...

test:
	@ go test -race ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

# goldenテストデータ更新
test-update:
	@ go test ./... -update

# goldenテストデータを全て削除して新規作成
test-clean:
	@ go test ./... -update -clean

gen:
	@ go generate ./...
