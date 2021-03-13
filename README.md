gokoku
=======

[![Test Status](https://github.com/Songmu/gokoku/workflows/test/badge.svg?branch=main)][actions]
[![Coverage Status](https://coveralls.io/repos/Songmu/gokoku/badge.png?branch=main)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/gokoku?status.svg)][godoc]

[actions]: https://github.com/Songmu/gokoku/actions?workflow=test
[coveralls]: https://coveralls.io/r/Songmu/gokoku?branch=main
[license]: https://github.com/Songmu/gokoku/blob/main/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/gokoku

gokoku for scaffolding

## Synopsis

```go
data := interface{}{}
err := gokoku.Scaffold(http.Dir("templatedir", ".", "dstdir", data)
```

## Description

scaffolding from http.FileSystem

## Installation

```console
% go get github.com/Songmu/gokoku
```

## Author

[Songmu](https://github.com/Songmu)
