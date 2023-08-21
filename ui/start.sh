# Start the backend API

#!/bin/bash

script_path=$(dirname $0)
orig_path=$(pwd)

cd ${script_path}

npm start

cd ${orig_path}
