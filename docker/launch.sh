#/bin/sh
cd /
./entrypoint.sh arangod --server.authentication=false

cd /stg/back/main
./main

cd /stg/front
npm start


