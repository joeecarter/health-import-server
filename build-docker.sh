VERSION=$1
if [[ $# -ne 1 ]]; then
    echo "USAGE: ./build-docker.sh <VERSION>" >&2
    exit 1
fi

echo "version=$VERSION"
echo ""

docker build --build-arg VERSION=$VERSION -t joecarter/health-import-server:$VERSION .
docker tag joecarter/health-import-server:$VERSION joecarter/health-import-server:latest

echo ""
echo "To publish the images:"
echo "docker push joecarter/health-import-server:$VERSION"
echo "docker push joecarter/health-import-server:latest"
