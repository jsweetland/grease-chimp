# Start the backend API

#!/bin/bash

script_path=$(dirname $0)
orig_path=$(pwd)

cd ${script_path}

go run main.go

cd ${orig_path}
