# Let's Go Bowling

An implementation of bowling scorer in CLI using go

## Usage

Inputs must be provided as following format:

| Input      | Translation|
| ----------- | ----------- |
| 'X'      | Strike       |
| '/'   | Spare     |
| '1'   | 1 point     |
| '2'   | 2 point     |
| '3'   | 3 point     |
| '4'   | 4 point     |
| '5'   | 5 point     |
| '6'   | 6 point     |
| '7'   | 7 point     |
| '8'   | 8 point     |
| '9'   | 9 point     |
| '-'   | 0 point     |

Scores are seperated by ",". 



## Build and Run

### Docker

```bash
docker build -t bowling .
docker run -it bowling
```

### local

```bash
make build
./bin/bowling
```
