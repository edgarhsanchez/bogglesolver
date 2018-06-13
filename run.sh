#! /bin/bash
echo "building go application..."
CGO_ENABLED=0 GOOS=linux go build -tags=jsoniter -a -installsuffix cgo -o main .

echo "setting permission on file..."
chmod +x ./main

echo "building client..."
cd ./client
ng build
cd ..

echo "building docker image..."
docker build --no-cache=false -t=bogglesolver -f Dockerfile .
((docker stop bogglesolver || echo 'No such container to stop' ) && docker rm bogglesolver || echo 'No such container to remove') && docker run -p 8000:8000 --name bogglesolver -d bogglesolver