set -o errexit
set -o nounset
set -o pipefail

find . -name "*.go" | grep -v "\/vendor\/" | xargs gofmt -s -w
