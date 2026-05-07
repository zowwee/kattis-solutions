# Kattis Solutions

My solutions to problems from [Kattis](https://open.kattis.com/), organized by language.

## Structure
```
kattis-solutions/
├── go/
│   └── <problem-name>/
│       └── main.go
├── python/
│   └── <problem-name>.py
└── README.md
```
Each Go solution lives in its own folder (required by Go's package structure). Other languages are organized as flat files within their language folder.

## Running a Solution

**Go:**
```bash
cd go/<problem-name>
go run main.go < sample.in
```

**Python:**
```bash
python3 python/<problem-name>.py < sample.in
```
