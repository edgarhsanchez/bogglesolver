#! /bin/bash
echo "building go application..."
CGO_ENABLED=0 GOOS=linux go build -tags=jsoniter -a -installsuffix cgo -o main .
echo "setting permission on file..."
chmod +x ./main
go test bogglesolver