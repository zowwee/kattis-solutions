# Kattis Solutions

My solutions to problems from [Kattis](https://open.kattis.com/), organized by language.

## Background

I'm self-taught and learning to code from scratch — no prior programming background. Primarily working in **Go**, with some **Python** on the side.

This repo isn't just a solutions dump — it's meant to document my journey. As I improve, I plan to revisit old problems and commit updated versions of my solutions, so the history reflects how my thinking and skills evolve over time.

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

## Tracking Progress

When I revisit a problem with a better solution, I commit the updated version over the old one. Git history preserves earlier attempts, so the evolution of my thinking is visible via `git log` and `git diff`.

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

## Notes

- Solutions prioritize clarity over cleverness — I'm here to learn, not to golf.
- Sample inputs (when included) are named `sample.in`.
- Expect rough edges. Watching them get smoother is half the point.
