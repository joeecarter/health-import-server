VERSION=$1
if [[ $# -ne 1 ]]; then
    echo "USAGE: ./build-docker.sh <VERSION>" >&2
    exit 1
fi

echo "version=$VERSION"
echo ""

docker buildx create --name healthimport
docker buildx use healthimport
docker buildx inspect --bootstrap

docker buildx build \
	--push \
	--build-arg VERSION=$VERSION \
	--platform linux/amd64,linux/arm64,linux/arm/v7 \
	-t joecarter/health-import-server:$VERSION \
	-t joecarter/health-import-server:latest \
	--push .

docker buildx rm healthimport
