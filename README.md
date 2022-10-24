# reminder
remind important events, like family`s birthday

##cmd run
```shell
go run . -f conf.yaml
```





## build image
```shell
docker build -t reminder:latest .
```


## docker run
```shell
docker run -d -v $(pwd)/conf.yaml:/app/conf.yaml --name reminder reminder:latest /app/main -f /app/conf.yaml
```
