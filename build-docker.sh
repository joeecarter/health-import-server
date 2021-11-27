VERSION=$1
if [[ $# -ne 1 ]]; then
    echo "USAGE: ./build-docker.sh <VERSION>" >&2
    exit 1
fi

echo "version=$VERSION"
echo ""

CGO_ENABLED=0 go build \
	-ldflags="-X 'main.Version=$VERSION'" \
	-o ./server cmd/server/main.go

docker build -t joecarter/health-import-server:$VERSION .
docker tag joecarter/health-import-server:$VERSION joecarter/health-import-server:latest

rm server

echo ""
echo "To publish the images:"
echo "docker push joecarter/health-import-server:$VERSION"
echo "docker push joecarter/health-import-server:latest"
