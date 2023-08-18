# Start the backend API

#!/bin/bash

original_path=$(pwd)
api_path=./api

cd ${api_path}
go run main.go
cd ${original_path}
