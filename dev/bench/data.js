window.BENCHMARK_DATA = {
  "lastUpdate": 1640118649964,
  "repoUrl": "https://github.com/structsdev/gen",
  "entries": {
    "Benchmark Results": [
      {
        "commit": {
          "author": {
            "email": "benji@devnw.com",
            "name": "Benji Vesterby",
            "username": "benjivesterby"
          },
          "committer": {
            "email": "benji@devnw.com",
            "name": "Benji Vesterby",
            "username": "benjivesterby"
          },
          "distinct": true,
          "id": "02562ada5308d3c4faee0a57db82fd0844299750",
          "message": "Package gen is a generic general use Go functions library\nwhich can be used throughout any project.\n\nAs a rule the methods in this package are non-destructive\nand return a new value rather than modifying the original.\n\nTo import this package it is recommended that you use a `.`\nimport like so:\n\nimport . \"go.structs.dev/gen\"\n\nThis allows for methods in `gen` to be used directly within\nthe same namespace without needing to prefix them with `gen.`",
          "timestamp": "2021-12-21T14:58:53-05:00",
          "tree_id": "669301bb7711a90e19d213bb0dde6a5da9115195",
          "url": "https://github.com/structsdev/gen/commit/02562ada5308d3c4faee0a57db82fd0844299750"
        },
        "date": 1640118649162,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 846.6,
            "unit": "ns/op",
            "extra": "1406085 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 743.6,
            "unit": "ns/op",
            "extra": "1636719 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.44,
            "unit": "ns/op",
            "extra": "220793038 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.096,
            "unit": "ns/op",
            "extra": "196073817 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 700.3,
            "unit": "ns/op",
            "extra": "1722424 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 737.2,
            "unit": "ns/op",
            "extra": "1624768 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 109.1,
            "unit": "ns/op",
            "extra": "10883948 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 9.774,
            "unit": "ns/op",
            "extra": "122799405 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.717,
            "unit": "ns/op",
            "extra": "323402560 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.729,
            "unit": "ns/op",
            "extra": "209928667 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.772,
            "unit": "ns/op",
            "extra": "207804331 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 5.662,
            "unit": "ns/op",
            "extra": "193629882 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 5.653,
            "unit": "ns/op",
            "extra": "216803198 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 4.466,
            "unit": "ns/op",
            "extra": "264280047 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 6.29,
            "unit": "ns/op",
            "extra": "191238183 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.055,
            "unit": "ns/op",
            "extra": "132664876 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 223.5,
            "unit": "ns/op",
            "extra": "5426756 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 177.7,
            "unit": "ns/op",
            "extra": "6778473 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 109.3,
            "unit": "ns/op",
            "extra": "10982353 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 456.2,
            "unit": "ns/op",
            "extra": "2496158 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 645.9,
            "unit": "ns/op",
            "extra": "1875405 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 567.6,
            "unit": "ns/op",
            "extra": "2103117 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 236,
            "unit": "ns/op",
            "extra": "4914632 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 290.3,
            "unit": "ns/op",
            "extra": "3986506 times\n2 procs"
          }
        ]
      }
    ]
  }
}