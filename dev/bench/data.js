window.BENCHMARK_DATA = {
  "lastUpdate": 1662496514393,
  "repoUrl": "https://github.com/finacore/commons-errors",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "name": "finacore",
            "username": "finacore"
          },
          "committer": {
            "name": "finacore",
            "username": "finacore"
          },
          "id": "17115f29a311bde5c813c6e6260fa69a30f07e31",
          "message": "organize benchmark",
          "timestamp": "2022-09-06T05:09:51Z",
          "url": "https://github.com/finacore/commons-errors/pull/10/commits/17115f29a311bde5c813c6e6260fa69a30f07e31"
        },
        "date": 1662496513624,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_CreateDefaultError",
            "value": 1e-7,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_MakeDefaultError",
            "value": 0.000002,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Status",
            "value": 0.000002,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_CreateValidationError",
            "value": 1e-7,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}