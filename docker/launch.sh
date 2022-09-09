#/bin/sh
cd /
./entrypoint.sh arangod --server.authentication=false --config /tmp/arangod.conf
