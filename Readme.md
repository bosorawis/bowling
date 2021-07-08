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

## Example Usage

```
>>> record your score for frame 1: X
scoring: 'X' for frame #1
finished frame #1
ScoreCard: [] | Current is 0
----------------------------
>>> record your score for frame 2: X
scoring: 'X' for frame #2
finished frame #2
ScoreCard: [] | Current is 0
----------------------------
>>> record your score for frame 3: X
scoring: 'X' for frame #3
finished frame #3
ScoreCard: [30] | Current is 30
----------------------------
>>> record your score for frame 4: X
scoring: 'X' for frame #4
finished frame #4
ScoreCard: [30 30] | Current is 60
----------------------------
>>> record your score for frame 5: X
scoring: 'X' for frame #5
finished frame #5
ScoreCard: [30 30 30] | Current is 90
----------------------------
>>> record your score for frame 6: X
scoring: 'X' for frame #6
finished frame #6
ScoreCard: [30 30 30 30] | Current is 120
----------------------------
>>> record your score for frame 7: X
scoring: 'X' for frame #7
finished frame #7
ScoreCard: [30 30 30 30 30] | Current is 150
----------------------------
>>> record your score for frame 8: X
scoring: 'X' for frame #8
finished frame #8
ScoreCard: [30 30 30 30 30 30] | Current is 180
----------------------------
>>> record your score for frame 9: X
scoring: 'X' for frame #9
finished frame #9
ScoreCard: [30 30 30 30 30 30 30] | Current is 210
----------------------------
>>> record your score for frame 10: X
scoring: 'X' for frame #10
failed to score: you get at least 2 rolls in the last frame
>>> record your score for frame 10: X,X,X
scoring: 'X,X,X' for frame #10
finished frame #10
ScoreCard: [30 30 30 30 30 30 30 30 30 30] | Current is 300
----------------------------
Total score is 300
good bye!
```


## How to run

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

### Tests

Here are some test cases with expected results in `game_test.go` under `TestRecordAndScore`

Some tests are also available in json form at `tests.json`
