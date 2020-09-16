 1. build app `docker build -t hellogo .` sesuaikan nama app nya 
 2. cek `docker images`
 3. run local `docker run -it -p 8080:8081 hellogo` 
 4. run in background `docker run -d -p 8080:8081 hellogo `
 5. cek running `docker ps`
 6. stop `docker kill 68748506924f` sesuaikan container idnya
