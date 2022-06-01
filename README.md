# QSSE Test
This is a sample project using [QSSE](https://www.github.com/snapp-incubator/qsse) library. QSSE is a library that implements SSE pattern over QUIC protocol.

## Commands
- `serve`: run QSSE server. used to publish and subscribe for exchanging events.
  - `ADDRESS/event/publush`: endpoint to publish events.
  - `ADDRESS:PORT`: endpoint to subscribe for events.
- `generate`: generates events with specified rate and size and publishes them to server.
- `subscribe`: subscribes from server on the given topics.


## Deployment
See `./deployment` directory for helm charts. deploy using `Makefile` or helm commands.
```shell
# install charts
make server-i
make generator-i
make subscriber-i

# uninstall charts
make server-u
make generator-u
make subscriber-u
```
### Notes
- deploy `server`,`generator`, `subscriber` in order to get proper functionality.   
- configure project in `./deployment/[chart-name]/values.yaml` before deployment.
