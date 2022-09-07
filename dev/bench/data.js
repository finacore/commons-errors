window.BENCHMARK_DATA = {
  "lastUpdate": 1662564259617,
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
          "id": "65ab3308b40a74728bba32fbfef8266f21056c2e",
          "message": "CI Organization",
          "timestamp": "2022-09-06T05:09:51Z",
          "url": "https://github.com/finacore/commons-errors/pull/11/commits/65ab3308b40a74728bba32fbfef8266f21056c2e"
        },
        "date": 1662564259035,
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
            "value": 0.0000024,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Set_DefaultError_Status",
            "value": 0.0000023,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_CreateValidationError",
            "value": 2e-7,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Set_ValidationError_Status",
            "value": 0.0000042,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}