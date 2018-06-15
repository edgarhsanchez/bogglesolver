#! /bin/bash
./build.sh

echo "building docker image..."
docker build --no-cache=false -t=bogglesolver -f Dockerfile .
((docker stop bogglesolver || echo 'No such container to stop' ) && docker rm bogglesolver || echo 'No such container to remove') && docker run -p 8000:80 --name bogglesolver -d bogglesolver