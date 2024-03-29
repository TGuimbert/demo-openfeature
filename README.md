# demo-openfeature

```shell
docker run -p 8013:8013 -v $(pwd)/:/etc/flagd/ -it --rm ghcr.io/open-feature/flagd:latest start --uri file:/etc/flagd/flags.flagd.json
go run main.go
curl http://localhost:8080/hello
curl http://localhost:8080/feature -H "email: tguimbert@domain.com"
curl http://localhost:8080/color -H "email: $(random)@domain.com"
```
