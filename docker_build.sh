# We are compiling our code specifically for a linux environment and
# then creating a docker image and copying in our binary.  This way
# we do NOT need a build environment in our Docker image, thus keeping
# it quite small.

# Build program for a linux target
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git rev-parse HEAD` -w -s" .

# Build Docker image
docker build -t dstroot/collection_machine .

# Push the docker image to Docker Hub
docker push dstroot/collection_machine

# Remove binaries
rm -rf collection_machine && rm -rf collection_machine.exe
