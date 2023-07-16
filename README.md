# influxdb-sandbox

```sh
docker run -d --name influxdb -p 8086:8086 influxdb
```

Connect to `localhost:8086` and proceed with the setup.

Get the token:

```sh
export INFLUXDB_TOKEN="<token>"
```

Run the code with data feed and queries:

```sh
go run .
```
