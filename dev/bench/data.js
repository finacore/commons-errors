window.BENCHMARK_DATA = {
  "lastUpdate": 1662490610501,
  "repoUrl": "https://github.com/finacore/commons-errors",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "committer": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "distinct": true,
          "id": "9c1fddc168c3a5ccd0cf57ee82d20f3a4db1066f",
          "message": "add bench to CI",
          "timestamp": "2022-09-06T14:53:50-04:00",
          "tree_id": "60b125cd6d24327522885e6f12473818eab39606",
          "url": "https://github.com/finacore/commons-errors/commit/9c1fddc168c3a5ccd0cf57ee82d20f3a4db1066f"
        },
        "date": 1662490608662,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCreateDefaultError/with-message",
            "value": 0.000002,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateDefaultError/without-message",
            "value": 0.0000015,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkMakeDefaultError/with-message",
            "value": 9e-7,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkMakeDefaultError/without-message",
            "value": 0.0000021,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateValidationError/with-data",
            "value": 0.0000018,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateValidationError/without-data",
            "value": 0.0000026,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}