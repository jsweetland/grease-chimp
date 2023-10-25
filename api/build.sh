# Build the backend API

#!/bin/bash

script_path=$(dirname $0)
orig_path=$(pwd)

cd ${script_path}

go build -o gc-api

cd ${orig_path}
