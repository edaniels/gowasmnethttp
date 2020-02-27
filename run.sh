set -e

trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT
GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
python -m SimpleHTTPServer 8000 &
open http://localhost:8000
read
