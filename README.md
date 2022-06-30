
###

full stack web application (backend and frontend) using Rust. (based from https://github.com/zupzup/rust-fullstack-example)


Arango: https://aragog.rs/ 
# docker run -v /home/thebogie/work/arangodb:/data -p 8529:8529 -e ARANGO_ROOT_PASSWORD=letmein -d --name arangodb arangodb

## Backend


## Frontend


## Config


## TODO
FrontEnd:
- figure out what needs to be fixed to get it up and running
- 
Backend
- fetch venue by address
  - how to connect to arango and get venue by address
- push venue by address
  - if already exists, update?
  - push element into arango



##Done
- for JSON: arangoexport --server.database smacktalk --collection game --output-directory "games"
- for argango backup/restore: arangodump --server.database smacktalk --output-directory "today"
- docker-compose.yml -> docker exec -it backend_arangodb_db_container_1 sh




