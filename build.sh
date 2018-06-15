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