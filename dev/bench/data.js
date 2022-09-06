window.BENCHMARK_DATA = {
  "lastUpdate": 1662495671852,
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
          "id": "9c5b78e34041c2f942207e182253de2fea7ee4ba",
          "message": "Worflow",
          "timestamp": "2022-09-06T05:09:51Z",
          "url": "https://github.com/finacore/commons-errors/pull/9/commits/9c5b78e34041c2f942207e182253de2fea7ee4ba"
        },
        "date": 1662495671018,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkCreateDefaultError",
            "value": 1e-7,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkMakeDefaultError",
            "value": 0.0000026,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkStatus",
            "value": 0.0000027,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateValidationError/with-data",
            "value": 0.0000017,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkCreateValidationError/without-data",
            "value": 0.0000019,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}