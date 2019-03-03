gokoku
=======

[![Build Status](https://travis-ci.org/Songmu/gokoku.png?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/gokoku/badge.png?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/gokoku?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/gokoku
[coveralls]: https://coveralls.io/r/Songmu/gokoku?branch=master
[license]: https://github.com/Songmu/gokoku/blob/master/LICENSE
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
