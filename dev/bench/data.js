window.BENCHMARK_DATA = {
  "lastUpdate": 1647392485034,
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
      },
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
          "id": "14f01848f8e38b623b4ab9e8af6d677d94b2e27f",
          "message": "Adding pre-release support",
          "timestamp": "2021-12-21T15:40:57-05:00",
          "tree_id": "1f83ee479a93fbc4d637ec5015074756699caacd",
          "url": "https://github.com/structsdev/gen/commit/14f01848f8e38b623b4ab9e8af6d677d94b2e27f"
        },
        "date": 1640119350835,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 887.8,
            "unit": "ns/op",
            "extra": "1340035 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 738.9,
            "unit": "ns/op",
            "extra": "1586739 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.502,
            "unit": "ns/op",
            "extra": "218426530 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.054,
            "unit": "ns/op",
            "extra": "198191524 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 699.8,
            "unit": "ns/op",
            "extra": "1671849 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 739.2,
            "unit": "ns/op",
            "extra": "1627898 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 109,
            "unit": "ns/op",
            "extra": "10569428 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 9.812,
            "unit": "ns/op",
            "extra": "122811325 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.736,
            "unit": "ns/op",
            "extra": "322788740 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.723,
            "unit": "ns/op",
            "extra": "210126021 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.752,
            "unit": "ns/op",
            "extra": "210045289 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 6.381,
            "unit": "ns/op",
            "extra": "185508796 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.319,
            "unit": "ns/op",
            "extra": "191180750 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.078,
            "unit": "ns/op",
            "extra": "235865794 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 7.104,
            "unit": "ns/op",
            "extra": "168975421 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.048,
            "unit": "ns/op",
            "extra": "132532044 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 221.9,
            "unit": "ns/op",
            "extra": "5361208 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 175.5,
            "unit": "ns/op",
            "extra": "6769393 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 109.2,
            "unit": "ns/op",
            "extra": "10995709 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 473.5,
            "unit": "ns/op",
            "extra": "2548551 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 689.9,
            "unit": "ns/op",
            "extra": "1741477 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 592.8,
            "unit": "ns/op",
            "extra": "2018983 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 241.3,
            "unit": "ns/op",
            "extra": "4984998 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 295.4,
            "unit": "ns/op",
            "extra": "4039898 times\n2 procs"
          }
        ]
      },
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
          "id": "25e6b8b4393e57a09474a6f6ad1480cecc4df37f",
          "message": "removing pre-release extraneous file",
          "timestamp": "2021-12-21T15:41:37-05:00",
          "tree_id": "e8113e676961ef537af248d632b70a2a40be39e8",
          "url": "https://github.com/structsdev/gen/commit/25e6b8b4393e57a09474a6f6ad1480cecc4df37f"
        },
        "date": 1640119373740,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1086,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 905.3,
            "unit": "ns/op",
            "extra": "1343647 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 6.307,
            "unit": "ns/op",
            "extra": "189950068 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 7.347,
            "unit": "ns/op",
            "extra": "167674318 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 849.8,
            "unit": "ns/op",
            "extra": "1412298 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 922.8,
            "unit": "ns/op",
            "extra": "1347706 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 127.1,
            "unit": "ns/op",
            "extra": "8508195 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 11.04,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 4.384,
            "unit": "ns/op",
            "extra": "277426980 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 6.715,
            "unit": "ns/op",
            "extra": "175670286 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 6.99,
            "unit": "ns/op",
            "extra": "182355734 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 7.795,
            "unit": "ns/op",
            "extra": "152337396 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 7.359,
            "unit": "ns/op",
            "extra": "160280418 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 6.082,
            "unit": "ns/op",
            "extra": "196431330 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 8.443,
            "unit": "ns/op",
            "extra": "139511617 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 10.56,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 269.2,
            "unit": "ns/op",
            "extra": "4422501 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 211.2,
            "unit": "ns/op",
            "extra": "5906156 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 122.9,
            "unit": "ns/op",
            "extra": "8707904 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 544.3,
            "unit": "ns/op",
            "extra": "2158142 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 818.3,
            "unit": "ns/op",
            "extra": "1498717 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 693,
            "unit": "ns/op",
            "extra": "1740441 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 282.6,
            "unit": "ns/op",
            "extra": "4408306 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 340.4,
            "unit": "ns/op",
            "extra": "3557041 times\n2 procs"
          }
        ]
      },
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
          "id": "c34d692341fac2517c284c0b47363091b0da6b73",
          "message": "Adding temp docs to readme until package site supports 1.18",
          "timestamp": "2021-12-21T16:37:57-05:00",
          "tree_id": "35b882b7e905cdd1047f77c03d4d927d8e1df6de",
          "url": "https://github.com/structsdev/gen/commit/c34d692341fac2517c284c0b47363091b0da6b73"
        },
        "date": 1640122750751,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 898.8,
            "unit": "ns/op",
            "extra": "1334900 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 749.2,
            "unit": "ns/op",
            "extra": "1601238 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.411,
            "unit": "ns/op",
            "extra": "221017009 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.081,
            "unit": "ns/op",
            "extra": "197609134 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 703.7,
            "unit": "ns/op",
            "extra": "1672268 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 738.6,
            "unit": "ns/op",
            "extra": "1501159 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 110.4,
            "unit": "ns/op",
            "extra": "10354953 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 9.769,
            "unit": "ns/op",
            "extra": "122222208 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.708,
            "unit": "ns/op",
            "extra": "322272927 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.731,
            "unit": "ns/op",
            "extra": "209705900 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.661,
            "unit": "ns/op",
            "extra": "217462939 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 6.384,
            "unit": "ns/op",
            "extra": "187751887 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.304,
            "unit": "ns/op",
            "extra": "188558094 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.068,
            "unit": "ns/op",
            "extra": "235238506 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 7.089,
            "unit": "ns/op",
            "extra": "169217475 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.042,
            "unit": "ns/op",
            "extra": "132097852 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 224.7,
            "unit": "ns/op",
            "extra": "4882898 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 178.7,
            "unit": "ns/op",
            "extra": "6631556 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 109.3,
            "unit": "ns/op",
            "extra": "10992548 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 481.1,
            "unit": "ns/op",
            "extra": "2415184 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 783.9,
            "unit": "ns/op",
            "extra": "1720015 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 600.3,
            "unit": "ns/op",
            "extra": "2012856 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 243.4,
            "unit": "ns/op",
            "extra": "4879921 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 299,
            "unit": "ns/op",
            "extra": "3998184 times\n2 procs"
          }
        ]
      },
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
          "id": "729b8ec47621cdde9d45ca44bd009c58175fc191",
          "message": "Update docs link for now",
          "timestamp": "2021-12-21T16:38:51-05:00",
          "tree_id": "e854f7fdd59f31f696894e7bd3e95a978acad082",
          "url": "https://github.com/structsdev/gen/commit/729b8ec47621cdde9d45ca44bd009c58175fc191"
        },
        "date": 1640122804771,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 883.6,
            "unit": "ns/op",
            "extra": "1359088 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 743.6,
            "unit": "ns/op",
            "extra": "1615004 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.406,
            "unit": "ns/op",
            "extra": "220407991 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.054,
            "unit": "ns/op",
            "extra": "198142250 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 733.5,
            "unit": "ns/op",
            "extra": "1701938 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 735,
            "unit": "ns/op",
            "extra": "1631238 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 109.2,
            "unit": "ns/op",
            "extra": "10771124 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 9.793,
            "unit": "ns/op",
            "extra": "121887679 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.766,
            "unit": "ns/op",
            "extra": "321877840 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.736,
            "unit": "ns/op",
            "extra": "209476618 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.826,
            "unit": "ns/op",
            "extra": "204432547 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 6.379,
            "unit": "ns/op",
            "extra": "186625851 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.269,
            "unit": "ns/op",
            "extra": "190293669 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.061,
            "unit": "ns/op",
            "extra": "233312378 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 7.058,
            "unit": "ns/op",
            "extra": "169185300 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.058,
            "unit": "ns/op",
            "extra": "131867041 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 225.3,
            "unit": "ns/op",
            "extra": "5401989 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 176.5,
            "unit": "ns/op",
            "extra": "6843694 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 109.3,
            "unit": "ns/op",
            "extra": "10989315 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 476.6,
            "unit": "ns/op",
            "extra": "2388470 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 694.5,
            "unit": "ns/op",
            "extra": "1732095 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 601.8,
            "unit": "ns/op",
            "extra": "2007049 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 241.9,
            "unit": "ns/op",
            "extra": "4861455 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 298.5,
            "unit": "ns/op",
            "extra": "3998232 times\n2 procs"
          }
        ]
      },
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
          "id": "5d5cfd4a006aeae0adca27371eec682ee51d76ec",
          "message": "Documentation updates",
          "timestamp": "2021-12-21T17:11:04-05:00",
          "tree_id": "ee2eeb3961808a85a1385ee80e16a5758eeff86f",
          "url": "https://github.com/structsdev/gen/commit/5d5cfd4a006aeae0adca27371eec682ee51d76ec"
        },
        "date": 1640124741490,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1146,
            "unit": "ns/op",
            "extra": "877141 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 939.9,
            "unit": "ns/op",
            "extra": "1274752 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 7.205,
            "unit": "ns/op",
            "extra": "159174075 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 7.238,
            "unit": "ns/op",
            "extra": "152286277 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 797.2,
            "unit": "ns/op",
            "extra": "1442079 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 823.3,
            "unit": "ns/op",
            "extra": "1522392 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 119.3,
            "unit": "ns/op",
            "extra": "9401952 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 11.96,
            "unit": "ns/op",
            "extra": "97944127 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.165,
            "unit": "ns/op",
            "extra": "349390258 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 6.528,
            "unit": "ns/op",
            "extra": "195507026 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 9.043,
            "unit": "ns/op",
            "extra": "125818892 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 9.14,
            "unit": "ns/op",
            "extra": "131567844 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.787,
            "unit": "ns/op",
            "extra": "168006884 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 6.363,
            "unit": "ns/op",
            "extra": "178660689 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 11.62,
            "unit": "ns/op",
            "extra": "88741708 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.19,
            "unit": "ns/op",
            "extra": "132538761 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 262.6,
            "unit": "ns/op",
            "extra": "5017768 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 172.3,
            "unit": "ns/op",
            "extra": "6998053 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 116.4,
            "unit": "ns/op",
            "extra": "11148278 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 629.9,
            "unit": "ns/op",
            "extra": "1943820 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 802.9,
            "unit": "ns/op",
            "extra": "1543612 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 664.6,
            "unit": "ns/op",
            "extra": "1664566 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 266.3,
            "unit": "ns/op",
            "extra": "4224739 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 334.8,
            "unit": "ns/op",
            "extra": "3773780 times\n2 procs"
          }
        ]
      },
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
          "id": "cf472da15fd11d1618fd748b5f5cba6d5a8ce19a",
          "message": "Documentation updates",
          "timestamp": "2021-12-21T17:11:45-05:00",
          "tree_id": "a7d20d25ac180429217071840e59adeb7034c8b8",
          "url": "https://github.com/structsdev/gen/commit/cf472da15fd11d1618fd748b5f5cba6d5a8ce19a"
        },
        "date": 1640124781845,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1081,
            "unit": "ns/op",
            "extra": "1091250 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 903.9,
            "unit": "ns/op",
            "extra": "1301865 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 6.558,
            "unit": "ns/op",
            "extra": "180719984 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 7.304,
            "unit": "ns/op",
            "extra": "154933354 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 864.8,
            "unit": "ns/op",
            "extra": "1368447 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 905.2,
            "unit": "ns/op",
            "extra": "1330126 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 132.6,
            "unit": "ns/op",
            "extra": "8817411 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 11.98,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 4.456,
            "unit": "ns/op",
            "extra": "268308478 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 6.924,
            "unit": "ns/op",
            "extra": "173957931 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 6.974,
            "unit": "ns/op",
            "extra": "174351189 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 7.665,
            "unit": "ns/op",
            "extra": "154513873 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 7.526,
            "unit": "ns/op",
            "extra": "155289939 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 6.13,
            "unit": "ns/op",
            "extra": "196631502 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 8.545,
            "unit": "ns/op",
            "extra": "139950787 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 10.85,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 268.9,
            "unit": "ns/op",
            "extra": "4463294 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 211.2,
            "unit": "ns/op",
            "extra": "5476144 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 132.6,
            "unit": "ns/op",
            "extra": "9154161 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 566.5,
            "unit": "ns/op",
            "extra": "2102337 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 831,
            "unit": "ns/op",
            "extra": "1450114 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 724.5,
            "unit": "ns/op",
            "extra": "1654984 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 294.3,
            "unit": "ns/op",
            "extra": "4087375 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 366.2,
            "unit": "ns/op",
            "extra": "3379506 times\n2 procs"
          }
        ]
      },
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
          "id": "ee08c021867b0248189f24071943e88bf7d92a2f",
          "message": "Adding As method for casting disparate types that implement the same interface",
          "timestamp": "2021-12-21T19:08:57-05:00",
          "tree_id": "256f68fd7d40355246892eb158d3a2ffedb41535",
          "url": "https://github.com/structsdev/gen/commit/ee08c021867b0248189f24071943e88bf7d92a2f"
        },
        "date": 1640131831722,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1033,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 797.8,
            "unit": "ns/op",
            "extra": "1342971 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 7.107,
            "unit": "ns/op",
            "extra": "186081985 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.712,
            "unit": "ns/op",
            "extra": "185227152 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 858.2,
            "unit": "ns/op",
            "extra": "1426656 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 897.4,
            "unit": "ns/op",
            "extra": "1431937 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 130.6,
            "unit": "ns/op",
            "extra": "9659966 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 13.15,
            "unit": "ns/op",
            "extra": "78754998 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.533,
            "unit": "ns/op",
            "extra": "338564676 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 7.218,
            "unit": "ns/op",
            "extra": "165846574 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 10.33,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 10.91,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 7.085,
            "unit": "ns/op",
            "extra": "170338245 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 7.035,
            "unit": "ns/op",
            "extra": "180674365 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 10.82,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 8.864,
            "unit": "ns/op",
            "extra": "133765206 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 258,
            "unit": "ns/op",
            "extra": "4235760 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 188.8,
            "unit": "ns/op",
            "extra": "6441552 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 119.9,
            "unit": "ns/op",
            "extra": "9788416 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 601.2,
            "unit": "ns/op",
            "extra": "2076860 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 842,
            "unit": "ns/op",
            "extra": "1398800 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 728.5,
            "unit": "ns/op",
            "extra": "1647612 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 290.4,
            "unit": "ns/op",
            "extra": "4439005 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 357.6,
            "unit": "ns/op",
            "extra": "3066010 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 97.57,
            "unit": "ns/op",
            "extra": "12787668 times\n2 procs"
          }
        ]
      },
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
          "id": "cfab69334e2a40f71193664dfbbd89af26afdf6e",
          "message": "Simplified As",
          "timestamp": "2021-12-22T05:12:46-05:00",
          "tree_id": "63023c71af2f40fcc347a1f487901b5e1af8c95d",
          "url": "https://github.com/structsdev/gen/commit/cfab69334e2a40f71193664dfbbd89af26afdf6e"
        },
        "date": 1640168045331,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 888.7,
            "unit": "ns/op",
            "extra": "1333962 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 742.5,
            "unit": "ns/op",
            "extra": "1586782 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.725,
            "unit": "ns/op",
            "extra": "208690321 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 5.752,
            "unit": "ns/op",
            "extra": "208739691 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 704.6,
            "unit": "ns/op",
            "extra": "1575880 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 756.4,
            "unit": "ns/op",
            "extra": "1599879 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 105.5,
            "unit": "ns/op",
            "extra": "10986272 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 12.39,
            "unit": "ns/op",
            "extra": "95883612 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.941,
            "unit": "ns/op",
            "extra": "311285142 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.776,
            "unit": "ns/op",
            "extra": "209300823 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 8.726,
            "unit": "ns/op",
            "extra": "137387656 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 9.039,
            "unit": "ns/op",
            "extra": "132784075 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 5.755,
            "unit": "ns/op",
            "extra": "208035318 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.081,
            "unit": "ns/op",
            "extra": "233225418 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 9.377,
            "unit": "ns/op",
            "extra": "127939176 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 5.58,
            "unit": "ns/op",
            "extra": "213812862 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 235.5,
            "unit": "ns/op",
            "extra": "4788390 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 173.3,
            "unit": "ns/op",
            "extra": "6926654 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 92.68,
            "unit": "ns/op",
            "extra": "12714752 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 475.2,
            "unit": "ns/op",
            "extra": "2543935 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 676.1,
            "unit": "ns/op",
            "extra": "1750147 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 588.6,
            "unit": "ns/op",
            "extra": "2031399 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 237.5,
            "unit": "ns/op",
            "extra": "4997341 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 310.8,
            "unit": "ns/op",
            "extra": "3849628 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.6694,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      },
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
          "id": "9fde850148d185617b44faaa26e48b0e5c0a4cb3",
          "message": "Adding conversion functions for channels",
          "timestamp": "2021-12-22T07:04:30-05:00",
          "tree_id": "976a1444527f73b12c43104ba96059464f84b5cf",
          "url": "https://github.com/structsdev/gen/commit/9fde850148d185617b44faaa26e48b0e5c0a4cb3"
        },
        "date": 1640174749542,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1241,
            "unit": "ns/op",
            "extra": "989551 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 1021,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 7.921,
            "unit": "ns/op",
            "extra": "154945718 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 8.178,
            "unit": "ns/op",
            "extra": "146470798 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 978.8,
            "unit": "ns/op",
            "extra": "1262792 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 1016,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 140.9,
            "unit": "ns/op",
            "extra": "8343006 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 14.9,
            "unit": "ns/op",
            "extra": "75820629 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.719,
            "unit": "ns/op",
            "extra": "320812562 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 7.684,
            "unit": "ns/op",
            "extra": "158146654 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 11.39,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 11.69,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 8.04,
            "unit": "ns/op",
            "extra": "152108923 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 7.742,
            "unit": "ns/op",
            "extra": "159239766 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 11.96,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 10.43,
            "unit": "ns/op",
            "extra": "99210589 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 333.8,
            "unit": "ns/op",
            "extra": "3427723 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 250.7,
            "unit": "ns/op",
            "extra": "4423426 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 137.8,
            "unit": "ns/op",
            "extra": "8910760 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 642.4,
            "unit": "ns/op",
            "extra": "1786789 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 992.3,
            "unit": "ns/op",
            "extra": "1300861 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 823.7,
            "unit": "ns/op",
            "extra": "1510219 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 324.3,
            "unit": "ns/op",
            "extra": "3977287 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 379,
            "unit": "ns/op",
            "extra": "3099450 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.4271,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 100.4,
            "unit": "ns/op",
            "extra": "12184453 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 104.1,
            "unit": "ns/op",
            "extra": "10626223 times\n2 procs"
          }
        ]
      },
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
          "id": "0ab8b615977b9fa87b53beb47f8342f7ad25d008",
          "message": "Adding a variadic close for multiple channels",
          "timestamp": "2021-12-22T15:58:54-05:00",
          "tree_id": "50ec73c3e723224089d15baf23742bb5db2362cf",
          "url": "https://github.com/structsdev/gen/commit/0ab8b615977b9fa87b53beb47f8342f7ad25d008"
        },
        "date": 1640206812267,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 887,
            "unit": "ns/op",
            "extra": "1335906 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 756,
            "unit": "ns/op",
            "extra": "1591798 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.404,
            "unit": "ns/op",
            "extra": "222910836 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.081,
            "unit": "ns/op",
            "extra": "195803894 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 721.9,
            "unit": "ns/op",
            "extra": "1682653 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 744.7,
            "unit": "ns/op",
            "extra": "1513059 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 110.1,
            "unit": "ns/op",
            "extra": "10837917 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 10.07,
            "unit": "ns/op",
            "extra": "122836389 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.736,
            "unit": "ns/op",
            "extra": "324858129 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.766,
            "unit": "ns/op",
            "extra": "196688350 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.852,
            "unit": "ns/op",
            "extra": "205883998 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 6.398,
            "unit": "ns/op",
            "extra": "184130246 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.264,
            "unit": "ns/op",
            "extra": "188799507 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.04,
            "unit": "ns/op",
            "extra": "237159180 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 7.094,
            "unit": "ns/op",
            "extra": "169618317 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.042,
            "unit": "ns/op",
            "extra": "132389830 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 226.7,
            "unit": "ns/op",
            "extra": "5412842 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 176.1,
            "unit": "ns/op",
            "extra": "6774502 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 109.1,
            "unit": "ns/op",
            "extra": "10794358 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 499.1,
            "unit": "ns/op",
            "extra": "2461579 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 702.4,
            "unit": "ns/op",
            "extra": "1725789 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 589.4,
            "unit": "ns/op",
            "extra": "2037460 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 238.6,
            "unit": "ns/op",
            "extra": "5252054 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 306.2,
            "unit": "ns/op",
            "extra": "3878106 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.3507,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 77.84,
            "unit": "ns/op",
            "extra": "14930173 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 78.86,
            "unit": "ns/op",
            "extra": "15103680 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1150,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          }
        ]
      },
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
          "id": "0f80950836eacd8d5041e5a7c9054660ebe3edd9",
          "message": "Adding documentation for Close to the readme",
          "timestamp": "2021-12-22T16:06:11-05:00",
          "tree_id": "ad655767edfcbf5275f1c3e268bed1675ba9ea8b",
          "url": "https://github.com/structsdev/gen/commit/0f80950836eacd8d5041e5a7c9054660ebe3edd9"
        },
        "date": 1640207244694,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 827.6,
            "unit": "ns/op",
            "extra": "1439307 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 692.4,
            "unit": "ns/op",
            "extra": "1663402 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.391,
            "unit": "ns/op",
            "extra": "251750094 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 5.463,
            "unit": "ns/op",
            "extra": "202547324 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 661.6,
            "unit": "ns/op",
            "extra": "1799448 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 689.1,
            "unit": "ns/op",
            "extra": "1695543 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 104.5,
            "unit": "ns/op",
            "extra": "10752661 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 9.747,
            "unit": "ns/op",
            "extra": "121980912 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.724,
            "unit": "ns/op",
            "extra": "324299223 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.734,
            "unit": "ns/op",
            "extra": "233107801 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.911,
            "unit": "ns/op",
            "extra": "238717180 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 6.386,
            "unit": "ns/op",
            "extra": "186779668 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.187,
            "unit": "ns/op",
            "extra": "193848735 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 4.446,
            "unit": "ns/op",
            "extra": "234988339 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 6.288,
            "unit": "ns/op",
            "extra": "191175534 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 8.039,
            "unit": "ns/op",
            "extra": "132749239 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 215.4,
            "unit": "ns/op",
            "extra": "5667304 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 164.9,
            "unit": "ns/op",
            "extra": "7171855 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 96.49,
            "unit": "ns/op",
            "extra": "10984514 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 460,
            "unit": "ns/op",
            "extra": "2669721 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 644.1,
            "unit": "ns/op",
            "extra": "1868577 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 554.3,
            "unit": "ns/op",
            "extra": "2137156 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 229.7,
            "unit": "ns/op",
            "extra": "5163919 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 283.2,
            "unit": "ns/op",
            "extra": "4211598 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.3346,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 71.32,
            "unit": "ns/op",
            "extra": "16383968 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 71.15,
            "unit": "ns/op",
            "extra": "16346478 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 942.3,
            "unit": "ns/op",
            "extra": "1252852 times\n2 procs"
          }
        ]
      },
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
          "id": "1dbfb098b311731af643135effcaa3d9b62c00a8",
          "message": "Adding Chan method to Slice, moving over from stream because it makes more sense",
          "timestamp": "2022-01-15T18:39:15-05:00",
          "tree_id": "98e5f2d9891cb1f63c1922f04f74a56b32d26402",
          "url": "https://github.com/structsdev/gen/commit/1dbfb098b311731af643135effcaa3d9b62c00a8"
        },
        "date": 1642290038457,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1034,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 835.6,
            "unit": "ns/op",
            "extra": "1407100 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 6.172,
            "unit": "ns/op",
            "extra": "194344338 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.852,
            "unit": "ns/op",
            "extra": "166400283 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 826.9,
            "unit": "ns/op",
            "extra": "1434482 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 840.3,
            "unit": "ns/op",
            "extra": "1347928 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 123.1,
            "unit": "ns/op",
            "extra": "9906406 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 10.99,
            "unit": "ns/op",
            "extra": "96193864 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 4.308,
            "unit": "ns/op",
            "extra": "280146636 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 6.833,
            "unit": "ns/op",
            "extra": "179607819 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 6.453,
            "unit": "ns/op",
            "extra": "183324078 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 7.501,
            "unit": "ns/op",
            "extra": "157148307 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 7.484,
            "unit": "ns/op",
            "extra": "163154212 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.656,
            "unit": "ns/op",
            "extra": "208600741 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 8.323,
            "unit": "ns/op",
            "extra": "145867158 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 10.48,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 274.1,
            "unit": "ns/op",
            "extra": "4156609 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 205.4,
            "unit": "ns/op",
            "extra": "5876319 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 125,
            "unit": "ns/op",
            "extra": "9439831 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 551.4,
            "unit": "ns/op",
            "extra": "2077384 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 801.8,
            "unit": "ns/op",
            "extra": "1499229 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 680.3,
            "unit": "ns/op",
            "extra": "1744298 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 274.1,
            "unit": "ns/op",
            "extra": "4255124 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 347.6,
            "unit": "ns/op",
            "extra": "3293686 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.3829,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 87.97,
            "unit": "ns/op",
            "extra": "13655185 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 90.47,
            "unit": "ns/op",
            "extra": "12897148 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1222,
            "unit": "ns/op",
            "extra": "1029399 times\n2 procs"
          }
        ]
      },
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
          "id": "d72cdb5da8e180427b248bb5a103ca5de552b20a",
          "message": "Correcting implementation of Slice.Map so it is comparable key, and any value",
          "timestamp": "2022-01-15T18:54:24-05:00",
          "tree_id": "d8022e78443601dfbd1bf379bef29f0e2c73e178",
          "url": "https://github.com/structsdev/gen/commit/d72cdb5da8e180427b248bb5a103ca5de552b20a"
        },
        "date": 1642290939296,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 939.4,
            "unit": "ns/op",
            "extra": "1287039 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 768.9,
            "unit": "ns/op",
            "extra": "1538524 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.717,
            "unit": "ns/op",
            "extra": "207591494 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 5.849,
            "unit": "ns/op",
            "extra": "207500818 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 700.1,
            "unit": "ns/op",
            "extra": "1703196 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 742.4,
            "unit": "ns/op",
            "extra": "1572356 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 104.8,
            "unit": "ns/op",
            "extra": "9730431 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 12.39,
            "unit": "ns/op",
            "extra": "95612318 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.865,
            "unit": "ns/op",
            "extra": "304448324 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.719,
            "unit": "ns/op",
            "extra": "208671079 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 8.726,
            "unit": "ns/op",
            "extra": "137621166 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 9.037,
            "unit": "ns/op",
            "extra": "132569664 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 5.744,
            "unit": "ns/op",
            "extra": "208556290 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.079,
            "unit": "ns/op",
            "extra": "234923798 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 9.386,
            "unit": "ns/op",
            "extra": "127690608 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 5.588,
            "unit": "ns/op",
            "extra": "212264721 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 224,
            "unit": "ns/op",
            "extra": "5151973 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 171.1,
            "unit": "ns/op",
            "extra": "6968660 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 93.03,
            "unit": "ns/op",
            "extra": "12625928 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 503,
            "unit": "ns/op",
            "extra": "2429152 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 695.9,
            "unit": "ns/op",
            "extra": "1728673 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 605.8,
            "unit": "ns/op",
            "extra": "1978932 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 247.9,
            "unit": "ns/op",
            "extra": "4776844 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 297.1,
            "unit": "ns/op",
            "extra": "4031918 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.6711,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 77.12,
            "unit": "ns/op",
            "extra": "15091500 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 78.67,
            "unit": "ns/op",
            "extra": "15061004 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1009,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          }
        ]
      },
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
          "id": "c58442d64fa19e214263757cf4d200de2eb2911c",
          "message": "Removing unnecessary recover",
          "timestamp": "2022-01-15T19:30:18-05:00",
          "tree_id": "04de5f5e04fd9d73e17aa17e8156731e86100fcf",
          "url": "https://github.com/structsdev/gen/commit/c58442d64fa19e214263757cf4d200de2eb2911c"
        },
        "date": 1642293100147,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1263,
            "unit": "ns/op",
            "extra": "835131 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 980.3,
            "unit": "ns/op",
            "extra": "1205823 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 8.161,
            "unit": "ns/op",
            "extra": "145702875 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 8.017,
            "unit": "ns/op",
            "extra": "156534326 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 893.4,
            "unit": "ns/op",
            "extra": "1333311 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 904.7,
            "unit": "ns/op",
            "extra": "1249033 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 131.8,
            "unit": "ns/op",
            "extra": "9345307 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 14.66,
            "unit": "ns/op",
            "extra": "82193035 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.44,
            "unit": "ns/op",
            "extra": "352559684 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 8.443,
            "unit": "ns/op",
            "extra": "137293232 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 10.62,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 10.57,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 7.812,
            "unit": "ns/op",
            "extra": "159927734 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 7.133,
            "unit": "ns/op",
            "extra": "172543694 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 10.99,
            "unit": "ns/op",
            "extra": "112000012 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.623,
            "unit": "ns/op",
            "extra": "121856818 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 301,
            "unit": "ns/op",
            "extra": "4092442 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 208.9,
            "unit": "ns/op",
            "extra": "5578252 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 124.2,
            "unit": "ns/op",
            "extra": "10382499 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 595.6,
            "unit": "ns/op",
            "extra": "2032910 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 880,
            "unit": "ns/op",
            "extra": "1386189 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 770,
            "unit": "ns/op",
            "extra": "1538277 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 299.4,
            "unit": "ns/op",
            "extra": "3960127 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 395,
            "unit": "ns/op",
            "extra": "3111321 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.4073,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 101.4,
            "unit": "ns/op",
            "extra": "11642370 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 103.7,
            "unit": "ns/op",
            "extra": "11233197 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1225,
            "unit": "ns/op",
            "extra": "902487 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benji@devnw.com",
            "name": "Benji Vesterby",
            "username": "benjivesterby"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "a4daa2fde544c313944b2ba6f6df11cb8a40fd0f",
          "message": "Merge pull request #3 from structsdev/dependabot/github_actions/golangci/golangci-lint-action-3",
          "timestamp": "2022-02-26T06:55:47-05:00",
          "tree_id": "f5d639c9de7f7187b8a7782ab76bbc712b5b0756",
          "url": "https://github.com/structsdev/gen/commit/a4daa2fde544c313944b2ba6f6df11cb8a40fd0f"
        },
        "date": 1645876624800,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1281,
            "unit": "ns/op",
            "extra": "842226 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 978.4,
            "unit": "ns/op",
            "extra": "1216718 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 7.406,
            "unit": "ns/op",
            "extra": "165544532 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 7.572,
            "unit": "ns/op",
            "extra": "159102336 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 908.9,
            "unit": "ns/op",
            "extra": "1249016 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 962.9,
            "unit": "ns/op",
            "extra": "1343102 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 133,
            "unit": "ns/op",
            "extra": "8210414 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 13.24,
            "unit": "ns/op",
            "extra": "96100135 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.41,
            "unit": "ns/op",
            "extra": "344842028 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 6.902,
            "unit": "ns/op",
            "extra": "166054071 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 10.05,
            "unit": "ns/op",
            "extra": "120981760 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 10.13,
            "unit": "ns/op",
            "extra": "120865527 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.58,
            "unit": "ns/op",
            "extra": "162975388 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 6.46,
            "unit": "ns/op",
            "extra": "177628059 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 10.29,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.702,
            "unit": "ns/op",
            "extra": "124790676 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 270.4,
            "unit": "ns/op",
            "extra": "4147894 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 190.5,
            "unit": "ns/op",
            "extra": "6668775 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 119,
            "unit": "ns/op",
            "extra": "10886788 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 605.9,
            "unit": "ns/op",
            "extra": "1949906 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 809.3,
            "unit": "ns/op",
            "extra": "1546521 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 742.1,
            "unit": "ns/op",
            "extra": "1630810 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 279.1,
            "unit": "ns/op",
            "extra": "3682653 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 356.3,
            "unit": "ns/op",
            "extra": "3570502 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.3657,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 95.68,
            "unit": "ns/op",
            "extra": "12539374 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 99.8,
            "unit": "ns/op",
            "extra": "12442370 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1236,
            "unit": "ns/op",
            "extra": "847666 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benji@devnw.com",
            "name": "Benji Vesterby",
            "username": "benjivesterby"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d3a7a5d832aa65a66ab96cf92f74e40e8a6e9612",
          "message": "Merge pull request #4 from structsdev/dependabot/github_actions/actions/checkout-3",
          "timestamp": "2022-03-02T09:26:44-05:00",
          "tree_id": "ba3c8dbb3994f829cbe0bd5ace6834004cb13ead",
          "url": "https://github.com/structsdev/gen/commit/d3a7a5d832aa65a66ab96cf92f74e40e8a6e9612"
        },
        "date": 1646231285011,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1250,
            "unit": "ns/op",
            "extra": "855205 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 1072,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 8.498,
            "unit": "ns/op",
            "extra": "130661595 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 8.348,
            "unit": "ns/op",
            "extra": "143716182 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 949,
            "unit": "ns/op",
            "extra": "1167157 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 1006,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 142.8,
            "unit": "ns/op",
            "extra": "7947097 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 14.77,
            "unit": "ns/op",
            "extra": "79243210 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.567,
            "unit": "ns/op",
            "extra": "334661673 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 9.219,
            "unit": "ns/op",
            "extra": "129729033 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 10.81,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 10.83,
            "unit": "ns/op",
            "extra": "110085680 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 8.1,
            "unit": "ns/op",
            "extra": "146073657 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 7.331,
            "unit": "ns/op",
            "extra": "162323924 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 11.49,
            "unit": "ns/op",
            "extra": "100000000 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 10.3,
            "unit": "ns/op",
            "extra": "98140675 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 341.2,
            "unit": "ns/op",
            "extra": "3736590 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 222.4,
            "unit": "ns/op",
            "extra": "5081560 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 129.5,
            "unit": "ns/op",
            "extra": "9344340 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 664.5,
            "unit": "ns/op",
            "extra": "1830297 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 927,
            "unit": "ns/op",
            "extra": "1302804 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 814.9,
            "unit": "ns/op",
            "extra": "1481600 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 320.5,
            "unit": "ns/op",
            "extra": "3700154 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 413.1,
            "unit": "ns/op",
            "extra": "2811762 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.4506,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 110,
            "unit": "ns/op",
            "extra": "11558701 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 111.7,
            "unit": "ns/op",
            "extra": "10117119 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1353,
            "unit": "ns/op",
            "extra": "895964 times\n2 procs"
          }
        ]
      },
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
          "id": "eeeb81c83b551379230a3b2a2cf5e2bcd93c3792",
          "message": "Fixing broken tests with release version of 1.18 and updated actions to pull release 1.18",
          "timestamp": "2022-03-15T19:53:55-04:00",
          "tree_id": "4588effca3e1d8ad78d299ed6c933f9b7b32a840",
          "url": "https://github.com/structsdev/gen/commit/eeeb81c83b551379230a3b2a2cf5e2bcd93c3792"
        },
        "date": 1647388546504,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 1061,
            "unit": "ns/op",
            "extra": "1148352 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 864.8,
            "unit": "ns/op",
            "extra": "1374780 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 7.169,
            "unit": "ns/op",
            "extra": "159991783 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 7.099,
            "unit": "ns/op",
            "extra": "169995744 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 790.6,
            "unit": "ns/op",
            "extra": "1448682 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 835.5,
            "unit": "ns/op",
            "extra": "1435990 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 117.1,
            "unit": "ns/op",
            "extra": "9774583 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 12.81,
            "unit": "ns/op",
            "extra": "91822621 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.17,
            "unit": "ns/op",
            "extra": "380285792 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 8.231,
            "unit": "ns/op",
            "extra": "146098334 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 9.64,
            "unit": "ns/op",
            "extra": "127262322 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 9.881,
            "unit": "ns/op",
            "extra": "120479936 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 7.24,
            "unit": "ns/op",
            "extra": "167712871 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 6.446,
            "unit": "ns/op",
            "extra": "188091044 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 10.63,
            "unit": "ns/op",
            "extra": "120598245 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.699,
            "unit": "ns/op",
            "extra": "131912907 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 259.1,
            "unit": "ns/op",
            "extra": "4275699 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 192.9,
            "unit": "ns/op",
            "extra": "5899593 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 114,
            "unit": "ns/op",
            "extra": "10139338 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 531.2,
            "unit": "ns/op",
            "extra": "2224834 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 784.5,
            "unit": "ns/op",
            "extra": "1526136 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 678.4,
            "unit": "ns/op",
            "extra": "1713084 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 265.6,
            "unit": "ns/op",
            "extra": "4546164 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 345.3,
            "unit": "ns/op",
            "extra": "3557563 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.4036,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 91.77,
            "unit": "ns/op",
            "extra": "12880587 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 83.69,
            "unit": "ns/op",
            "extra": "13710945 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1089,
            "unit": "ns/op",
            "extra": "1000000 times\n2 procs"
          }
        ]
      },
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
          "id": "a1c057f5eab5d55135448d0dfeef21c63a96fca9",
          "message": "Removing the notice regarding 1.18beta since 1.18 is released",
          "timestamp": "2022-03-15T21:00:07-04:00",
          "tree_id": "798235d868dc98b7701871f494262a7c0a36c8ae",
          "url": "https://github.com/structsdev/gen/commit/a1c057f5eab5d55135448d0dfeef21c63a96fca9"
        },
        "date": 1647392484571,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Unique",
            "value": 949.6,
            "unit": "ns/op",
            "extra": "1253372 times\n2 procs"
          },
          {
            "name": "Benchmark_Unique_NoUnique",
            "value": 787.9,
            "unit": "ns/op",
            "extra": "1513936 times\n2 procs"
          },
          {
            "name": "Benchmark_Has",
            "value": 5.451,
            "unit": "ns/op",
            "extra": "221897654 times\n2 procs"
          },
          {
            "name": "Benchmark_Has_DoesntHave",
            "value": 6.058,
            "unit": "ns/op",
            "extra": "197264743 times\n2 procs"
          },
          {
            "name": "Benchmark_Match",
            "value": 721.8,
            "unit": "ns/op",
            "extra": "1593207 times\n2 procs"
          },
          {
            "name": "Benchmark_Match_Doesnt",
            "value": 748.8,
            "unit": "ns/op",
            "extra": "1553952 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices",
            "value": 111,
            "unit": "ns/op",
            "extra": "10749937 times\n2 procs"
          },
          {
            "name": "Benchmark_Indices_None",
            "value": 9.785,
            "unit": "ns/op",
            "extra": "122622632 times\n2 procs"
          },
          {
            "name": "Benchmark_Index",
            "value": 3.701,
            "unit": "ns/op",
            "extra": "322704602 times\n2 procs"
          },
          {
            "name": "Benchmark_Index_None",
            "value": 5.729,
            "unit": "ns/op",
            "extra": "209778205 times\n2 procs"
          },
          {
            "name": "Benchmark_comp",
            "value": 5.418,
            "unit": "ns/op",
            "extra": "219903192 times\n2 procs"
          },
          {
            "name": "Benchmark_comp_nomatch",
            "value": 6.389,
            "unit": "ns/op",
            "extra": "186477996 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal",
            "value": 6.195,
            "unit": "ns/op",
            "extra": "192464308 times\n2 procs"
          },
          {
            "name": "Benchmark_Equal_not",
            "value": 5.058,
            "unit": "ns/op",
            "extra": "236466163 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare",
            "value": 7.137,
            "unit": "ns/op",
            "extra": "167272118 times\n2 procs"
          },
          {
            "name": "Benchmark_Compare_err",
            "value": 9.051,
            "unit": "ns/op",
            "extra": "132582028 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_full",
            "value": 234.7,
            "unit": "ns/op",
            "extra": "5189721 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_partial",
            "value": 181.6,
            "unit": "ns/op",
            "extra": "6563817 times\n2 procs"
          },
          {
            "name": "Benchmark_Intersect_none",
            "value": 109.4,
            "unit": "ns/op",
            "extra": "10933699 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_all",
            "value": 520.6,
            "unit": "ns/op",
            "extra": "2253091 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_partial",
            "value": 745.6,
            "unit": "ns/op",
            "extra": "1607192 times\n2 procs"
          },
          {
            "name": "Benchmark_Diff_none",
            "value": 629.3,
            "unit": "ns/op",
            "extra": "1901414 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude",
            "value": 255,
            "unit": "ns/op",
            "extra": "4632735 times\n2 procs"
          },
          {
            "name": "Benchmark_Exclude_none",
            "value": 325.1,
            "unit": "ns/op",
            "extra": "3709148 times\n2 procs"
          },
          {
            "name": "Benchmark_As",
            "value": 0.3358,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "Benchmark_ReadOnly",
            "value": 81.38,
            "unit": "ns/op",
            "extra": "14473999 times\n2 procs"
          },
          {
            "name": "Benchmark_WriteOnly",
            "value": 77.97,
            "unit": "ns/op",
            "extra": "15119395 times\n2 procs"
          },
          {
            "name": "Benchmark_Close",
            "value": 1061,
            "unit": "ns/op",
            "extra": "1139481 times\n2 procs"
          }
        ]
      }
    ]
  }
}